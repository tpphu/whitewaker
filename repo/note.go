package repo

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/helper"
	"github.com/tpphu/whitewalker/model"
)

// NoteRepo interface
type NoteRepo interface {
	Find(context.Context, uint) (*model.Note, error)
	List(helper.Pagination) ([]model.Note, error)
	Update(uint, model.Note) error
	Delete(uint) error
	Create(model.Note) (*model.Note, error)
}

// NoteRepoImpl struct
type NoteRepoImpl struct {
	DB *gorm.DB
}

// Create returns a note
func (noteRepo NoteRepoImpl) Create(note model.Note) (*model.Note, error) {
	err := noteRepo.DB.Create(&note).Error
	return &note, err
}

// Find a note
func (noteRepo NoteRepoImpl) Find(id uint) (*model.Note, error) {
	note := &model.Note{}
	err := noteRepo.DB.Find(note, id).Error
	return note, err
}

// List notes
func (noteRepo NoteRepoImpl) List(pagination helper.Pagination) ([]model.Note, error) {
	notes := []model.Note{}
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := noteRepo.DB.Offset(offset).
		Limit(limit).
		Find(&notes).
		Error
	return notes, err
}

// Update a note
func (noteRepo NoteRepoImpl) Update(id uint, note model.Note) error {
	err := noteRepo.DB.Model(&note).Where("id = ?", id).Update(&note).Error
	return err
}

// Delete a note
func (noteRepo NoteRepoImpl) Delete(id uint) error {
	err := noteRepo.DB.Where("id = ?", id).Delete(&model.Note{}).Error
	return err
}
