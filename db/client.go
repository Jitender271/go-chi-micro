package db

import (

    "github.com/go-chi-micro/model"
)

type SqlClient interface {

    CreateBlogRecord(bl *model.Blog) (model.Blog, error)
    GetBlogs(string) ( model.Blog, error)

}

func NewClient(config *Config) SqlClient {
    return &sqlClient{
        config: config,
    }
}

type Config struct {
    DBConnection string
}

type sqlClient struct {
    config *Config
}


func (c *sqlClient) CreateBlogRecord(bl *model.Blog) (model.Blog, error) {
    return createBlog(bl)
}

func (c *sqlClient) GetBlogs(id string)  (model.Blog, error) {
    return getAllBlog(id)
}