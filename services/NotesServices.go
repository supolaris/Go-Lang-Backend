package services

import "gorm.io/gorm"

type NotesService struct {
	db *gorm.DB
}
type Notes struct {
	Id    int
	Title string
}

func (n *NotesService) GetNotesService() []Notes {
	data := []Notes{
		{Id: 1, Title: "Harry Potter"},
		{Id: 2, Title: "Puzzle Quest"},
	}
	return data
}

func (n *NotesService) PostNotesService() []Notes {
	data := []Notes{
		{Id: 3, Title: "Hello world"},
	}
	return data
}