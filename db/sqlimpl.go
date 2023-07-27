package db

import (
	"errors"
	"fmt"

	"github.com/go-chi-micro/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)


func createBlog(bl *model.Blogs)(model.BlogData, error){
    db:= GetDBConnection()
    if err := db.Table("blogs").Create(&bl).Error; err != nil{
        log.Info("failure", model.Blogs{}, err)
    }
    record := model.BlogData{
        Blog: *bl,
        Message: "data saved",
    }
    return record, nil

}

func getAllBlog(id string) (model.Blogs, error){
    var record []model.Blogs
    db:= GetDBConnection()
    if err := db.Table("blogs").Where("id=?", id).Find(&record).Error; err != nil {
        log.Info("failure", []model.Blogs{})
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return model.Blogs{}, fmt.Errorf("blog with IDs %s not found", id)
        }
        return model.Blogs{}, fmt.Errorf("failed to get blog: %w", err)
    }
    if len(record) == 0 {
        return model.Blogs{}, gorm.ErrRecordNotFound
    }
    return record[0], nil
}

func updateBlog(id string, bl *model.Blogs) (model.BlogData, error){

    db:= GetDBConnection()
    if err := db.Table("blogs").Where("id=?", id).Updates(&bl).Error; err != nil {
        log.Info("failure", []model.Blogs{})
    }
    record := model.BlogData{
        Blog: *bl,
        Message: "record updated successfully",
    }
    return record, nil
}

func deleteBlog(id string) (string, error){
    var bl model.Blogs
    db:= GetDBConnection()
    if err := db.Table("blogs").Where("id=?", id).Delete(&bl).Error; err != nil {
        log.Info("failure", []model.Blogs{})
        return "not able to delete", err
    }
   
    return "deleted successfully", nil
}
