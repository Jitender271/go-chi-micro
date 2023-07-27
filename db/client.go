package db

import (
    "github.com/go-chi-micro/model"
)

type SqlClient interface {
    CreateBlogRecord(bl *model.Blogs) (model.BlogData, error)
    GetBlogs(string) ( model.Blogs, error)
    UpdateBlogs(id string, bl *model.Blogs) (model.BlogData, error)
    DeleteBlog(id string)(string, error)
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


func (c *sqlClient) CreateBlogRecord(bl *model.Blogs) (model.BlogData, error) {
    return createBlog(bl)
}

func (c *sqlClient) GetBlogs(id string)  (model.Blogs, error) {
    return getAllBlog(id)
}

func (c *sqlClient) UpdateBlogs(id string,bl *model.Blogs)(model.BlogData, error) {
    return updateBlog(id, bl)
}

func (c *sqlClient) DeleteBlog(id string)(string, error){
    return deleteBlog(id)
}