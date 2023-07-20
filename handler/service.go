package handler

import (
    "github.com/go-chi-micro/db"
    "github.com/go-chi-micro/model"
    log "github.com/sirupsen/logrus"
)

type Service interface {
    CreateRecordCoreTeam(bl *model.Blogs) (model.Blogs, error)
GetRecordSetPost(id string) (model.Blogs, error)
    
}

func NewService(sqlDB db.SqlClient) Service {
    return &service{
        sqlDB: sqlDB,
    }
}

type service struct {
    sqlDB db.SqlClient
}


func (s *service) CreateRecordCoreTeam(bl *model.Blogs)( model.Blogs, error){
    record, err := s.sqlDB.CreateBlogRecord(bl)
    if err != nil {
        log.Info("Failure: mot getting data from Table", err)
        return model.Blogs{}, err
    }
    return record, nil
}

func (s *service) GetRecordSetPost(id string) (model.Blogs, error) {
    record, err := s.sqlDB.GetBlogs(id)
    if err != nil {
        log.Info("Failure: mot getting data from Table", err)
        return model.Blogs{}, err
    }
    return record, nil
}