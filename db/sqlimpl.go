package db

import (
    "fmt"

    "github.com/go-chi-micro/model"
    log "github.com/sirupsen/logrus"
)


func createBlog(bl *model.Blog)(model.Blog, error){
    
    db:= GetDBConnection()
    if err := db.Table("blog").Create(&bl).Error; err != nil{
        log.Info("failure", model.Blog{}, err)
    }
    recor := model.Blog{
        Blogname: "saved",
        BlogDetails: "done",
    }
    return recor, nil

}

func getAllBlog(id string) (model.Blog, error){
    var record []model.Blog
    db:= GetDBConnection()
    if err := db.Table("blog").Where("id=?", id).Find(&record).Error; err != nil {
        log.Info("failure", []model.Blog{})
    }
    fmt.Println("print ====>", record[0])
    return record[0], nil

}
