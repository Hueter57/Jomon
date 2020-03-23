package model

import (
	"github.com/gofrs/uuid"
	storagePkg "github.com/traPtitech/Jomon/storage"
	"io"
)

type ApplicationsImage struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key"`
	ApplicationID uuid.UUID `gorm:"type:char(36);not null"`
	MimeType      string    `gorm:"type:text;not null" json:"-"`
}

type ApplicationsImageRepository interface {
	CreateApplicationsImage(applicationId uuid.UUID, src io.Reader, ext string) error
	DeleteApplicationsImage(id uuid.UUID) error
}

type applicationsImageRepository struct {
	storage storagePkg.Storage
}

func NewApplicationsImageRepository(storage storagePkg.Storage) ApplicationsImageRepository {
	return &applicationsImageRepository{storage: storage}
}

func (repo *applicationsImageRepository) CreateApplicationsImage(applicationId uuid.UUID, src io.Reader, ext string) error {
	panic("implement me")
}

func (repo *applicationsImageRepository) DeleteApplicationsImage(id uuid.UUID) error {
	panic("implement me")
}
