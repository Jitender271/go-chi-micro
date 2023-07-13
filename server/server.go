package server

import (
    "fmt"

    "github.com/go-chi-micro/config"
    "github.com/go-chi-micro/db"
    "github.com/go-chi-micro/handler"

    "net"
    "net/http"
    "os"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
)

type Server struct {
    httpServer *http.Server
    router     *chi.Mux
}

func New() *Server {
    r := chi.NewRouter()
    r.Use(render.SetContentType(render.ContentTypeJSON))
    db.InitMysql()

    mysqlFGClient := db.NewClient(
        &db.Config{
            DBConnection: "",
        })

    run := handler.NewService(mysqlFGClient)
    setupRoutesForUpdate(run, r)

    server := newServer(r)

    return server
}

func setupRoutesForUpdate(service handler.Service, r *chi.Mux) {
    // plug in sub-routers for resources: feature gate
    // this pattern also allows for easy integration testing. see api_test.go

    r.Route("/api", func(r chi.Router) {
        r.Mount("/v1", handler.Handler(service))
    })
}

func (s *Server) ListenAndServe() error {
    l, err := net.Listen("tcp", ":"+s.httpServer.Addr)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    return s.httpServer.Serve(l)
}

func newServer(r *chi.Mux) *Server {
    fmt.Println("****Server Started on", config.GetYamlValues().ServerConfig.Port, "****")
    return &Server{
        httpServer: &http.Server{Addr: config.GetYamlValues().ServerConfig.Port, Handler: r},
        router:     r,
    }
}