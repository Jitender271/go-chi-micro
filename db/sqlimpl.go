package db

import (
    "fmt"

    "github.com/go-chi-micro/model"
    log "github.com/sirupsen/logrus"
)


func createBlog(bl *model.Blogs)(model.Blogs, error){
    
    db:= GetDBConnection()
    if err := db.Table("blogs").Create(&bl).Error; err != nil{
        log.Info("failure", model.Blogs{}, err)
    }
    recor := model.Blogs{
        BlogName: "saved",
        BlogDetails: "done",
    }
    return recor, nil

}

func getAllBlog(id string) (model.Blogs, error){
    var record []model.Blogs
    db:= GetDBConnection()
    if err := db.Table("blog").Where("id=?", id).Find(&record).Error; err != nil {
        log.Info("failure", []model.Blogs{})
    }
    fmt.Println("print ====>", record[0])
    return record[0], nil

}
