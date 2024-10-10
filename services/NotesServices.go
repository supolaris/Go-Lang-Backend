package services

import (
	"fmt"
	internal "goProject/internals/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}
type Notes struct {
	Id    int
	Title string
}

func (n *NotesService) InitService(database *gorm.DB){
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

func (n *NotesService) GetNotesService(status bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	if err := n.db.Where("status = ?", status).Find(&notes).Error; err !=nil {
		return nil,err
	} else {
		return notes,nil
	}
}

func (n *NotesService) PostNotesService() []Notes {
	data := []Notes{
		{Id: 3, Title: "Hello world"},
	}
	return data
}

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note :=  &internal.Notes{
		Title: title,
		Status: status,
	}
	if err := n.db.Create(note).Error; err !=nil {
		fmt.Println("Error!", err)

		return nil,err
	}
	return note,nil;
}