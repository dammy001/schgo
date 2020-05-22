package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Job struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255;not null;unique" json:"title"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	Benefit     string    `gorm:"size:255;not null;unique" json:"benefit"`
	Experience  string    `gorm:"size:255;not null;" json:"experience"`
	Price       uint64    `gorm:"not null;" json:"price"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Job) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Benefit = html.EscapeString(strings.TrimSpace(p.Benefit))
	p.Experience = html.EscapeString(strings.TrimSpace(p.Experience))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Job) Validate() error {
	if p.Title == "" {
		return errors.New("Title field required")
	}
	if p.Description == "" {
		return errors.New("Description field required")
	}
	if p.Benefit == "" {
		return errors.New("Benefit field required")
	}
	if p.Experience == "" {
		return errors.New("Experience field is required")
	}
	if p.Price < 1 {
		return errors.New("Price field is required")
	}
	return nil
}

func (p *Job) SaveJob(db *gorm.DB) (*Job, error) {

	var err error
	err = db.Debug().Model(&Job{}).Create(&p).Error

	if err != nil {
		return &Job{}, err
	}

	return p, nil
}

func (p *Job) allJobsListing(db *gorm.DB) (*[]Job, error) {
	var err error
	jobs := []Job{}

	err = db.Debug().Model(&Job{}).Find(&jobs).Error
	if err != nil {
		return &[]Job{}, err
	}
	return &jobs, nil
}

func (p *Job) FindJobByID(db *gorm.DB, pid uint64) (*Job, error) {
	var err error
	err = db.Debug().Model(&Job{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Job{}, err
	}
	return p, nil
}

func (p *Job) UpdateJob(db *gorm.DB) (*Job, error) {

	var err error

	err = db.Debug().Model(&Job{}).Where("id = ?", p.ID).Updates(Job{Title: p.Title, Description: p.Description, Benefit: p.Benefit, Experience: p.Experience, Price: p.Price, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Job{}, err
	}
	return p, nil
}

func (p *Job) DeleteJob(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Job{}).Where("id = ?", pid).Take(&Job{}).Delete(&Job{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Job not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
