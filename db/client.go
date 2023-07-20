package db

import (

    "github.com/go-chi-micro/model"
)

type SqlClient interface {

    CreateBlogRecord(bl *model.Blogs) (model.Blogs, error)
    GetBlogs(string) ( model.Blogs, error)

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


func (c *sqlClient) CreateBlogRecord(bl *model.Blogs) (model.Blogs, error) {
    return createBlog(bl)
}

func (c *sqlClient) GetBlogs(id string)  (model.Blogs, error) {
    return getAllBlog(id)
}