package handler

import (
    "net/http"

    "github.com/go-chi-micro/httphandler"
    "github.com/go-chi-micro/model"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    log "github.com/sirupsen/logrus"
)

type ctx struct {
    store Service
    h     func(Service, http.ResponseWriter, *http.Request)
}

func (g *ctx) handle() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        g.h(g.store, w, r)
    }
}

func Handler(store Service) http.Handler {
    r := chi.NewRouter()
    //createBlog := ctx{store: store, h: createBlogs}
    getRecordSetPost := ctx{store: store, h:getRecordSetPost}
    createBlogs := ctx{store: store, h:createBlogs}

    r.Get(httphandler.WrapHandlerFunc("/blog/{data}", "create blog", getRecordSetPost.handle()))
    r.Post(httphandler.WrapHandlerFunc("/blog", "create blog", createBlogs.handle()))

    return r
}


func createBlogs(store Service, w http.ResponseWriter, r *http.Request) {
    
    data := &Request{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
        return
    }
    recordSchema := data.Blogs

    //s := chi.URLParam(r,"db")
    services, err := store.CreateRecordCoreTeam(recordSchema)
    if err != nil {
        log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
        httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
        return
    }
    render.Status(r, http.StatusOK)
    render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func getRecordSetPost(store Service, w http.ResponseWriter, r *http.Request) {
    data := chi.URLParam(r, "data")
    services, err := store.GetRecordSetPost(data)
    if err != nil {
        log.Errorf("Unable To Fetch stats ", services, err)
        return
    }
    render.Status(r, http.StatusOK)
    render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

type Request struct {
    *model.Blogs
}

func (a *Request) Bind(r *http.Request) error {
    //TODO: to be expanded
    return nil
}

type Response struct {
    Meta interface{}
    Data interface{}
}
