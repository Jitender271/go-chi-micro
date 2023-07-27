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
    getRecordSetPost := ctx{store: store, h:getRecordSetPost}
    createBlogs := ctx{store: store, h:createBlogs}
    updateBlogs := ctx{store: store, h:updateBlogs}
    deleteBlogs := ctx{store: store, h:deleteBlogs}

    r.Get(httphandler.WrapHandlerFunc("/blog/{data}", "get blog", getRecordSetPost.handle()))
    r.Post(httphandler.WrapHandlerFunc("/blog", "create blog", createBlogs.handle()))
    r.Put(httphandler.WrapHandlerFunc("/blog/{data}", "update blog", updateBlogs.handle()))
    r.Delete(httphandler.WrapHandlerFunc("/blog/remove/{data}", "delete blog", deleteBlogs.handle()))

    return r
}


func createBlogs(store Service, w http.ResponseWriter, r *http.Request) {
    
    data := &Request{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
        return
    }
    recordSchema := data.Blogs
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
        log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
        httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
        return
    }

    render.Status(r, http.StatusOK)
    render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func updateBlogs(store Service, w http.ResponseWriter, r *http.Request) {
    data := chi.URLParam(r, "data")
    payload := &Request{}
    if err := render.Bind(r, payload); err != nil {
        render.Render(w, r, httphandler.ErrInvalidRequest(err, "Invalid Request"))
        return
    }
    recordSchema := payload.Blogs
    services, err := store.UpdateBlog(data, recordSchema)
    if err != nil {
        log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
        httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
        return
    }
    render.Status(r, http.StatusOK)
    render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}

func deleteBlogs(store Service, w http.ResponseWriter, r *http.Request) {
    data := chi.URLParam(r, "data")
    services, err := store.DeleteBlogs(data)
    if err != nil {
        log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
        httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
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
