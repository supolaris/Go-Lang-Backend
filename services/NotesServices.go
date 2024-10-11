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

func (n *NotesService) GetSingleNotesService(id int64) ([]*internal.Notes, error) {
	var note []*internal.Notes
	if err := n.db.Where("id = ?", id).Find(&note).Error; err !=nil {
		return nil,err
	} else {
		return note,nil
	}
}

func (n *NotesService) GetNotesService(status *bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	query := n.db
	if status != nil {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&notes).Error; err !=nil {
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

func (n *NotesService) UpdateNotesService(title string, status bool, id int) (*internal.Notes, error) {
	var note internal.Notes
	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
			return nil, err
	}
	note.Title = title
	note.Status = status
	if err := n.db.Save(&note).Error; err != nil {
			fmt.Println("Error!", err)
			return nil, err
	}
	return &note, nil
}

func (n *NotesService) DeleteNotesService(id int64) (error) {
	var note internal.Notes
	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return err
	} 
	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		fmt.Println(err);
		return err
	}
	return nil
}