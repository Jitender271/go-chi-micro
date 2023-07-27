package handler

import (

	"github.com/go-chi-micro/db"
	"github.com/go-chi-micro/model"
	log "github.com/sirupsen/logrus"
)

type Service interface {
    CreateRecordCoreTeam(bl *model.Blogs) (model.BlogData, error)
    GetRecordSetPost(id string) (model.Blogs, error)
    UpdateBlog(id string, bl *model.Blogs) (model.BlogData, error)
    DeleteBlogs(id string)(string, error)
    
}

func NewService(sqlDB db.SqlClient) Service {
    return &service{
        sqlDB: sqlDB,
    }
}

type service struct {
    sqlDB db.SqlClient
}


func (s *service) CreateRecordCoreTeam(bl *model.Blogs)(model.BlogData, error){
    record, err := s.sqlDB.CreateBlogRecord(bl)
    if err != nil {
        log.Info("Failure: mot getting data from Table", err)
        return model.BlogData{}, err
    }
    return record, nil
}

func (s *service) GetRecordSetPost(id string) (model.Blogs, error) {
    record, err := s.sqlDB.GetBlogs(id)
    if err != nil {
        log.Info("Failure: not getting data from Table", err)
        return model.Blogs{}, err
    }
    return record, nil
}

func (s *service) UpdateBlog(id string, bl *model.Blogs) (model.BlogData, error) {
    record, err := s.sqlDB.UpdateBlogs(id, bl)
    if err != nil {
        log.Info("Failure: mot getting data from Table", err)
        return model.BlogData{}, err
    }
    return record, nil
}

func (s *service) DeleteBlogs(id string)(string, error) {
    record, err := s.sqlDB.DeleteBlog(id)
    if err != nil {
        log.Info("Failure: not getting data from Table", err)
        return "record is not available in the system", err
    }
    return record, nil

}
