package service

import (
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/model"
	storagePkg "github.com/traPtitech/Jomon/storage"
)

type Service interface {
	CreateFile(src io.Reader, mimetype string) (File, error)
	OpenFile(fileID uuid.UUID, mimetype string) (io.ReadCloser, error)
	DeleteFile(fileID uuid.UUID) error
	GetAccessToken(code string, codeVerifier string) (AuthResponse, error)
	GetClientId() string
}
type Services struct {
	Auth    Auth
	Storage storagePkg.Storage
}

func NewServices(repo model.Repository, storage storagePkg.Storage) (*Services, error) {
	traQClientID := os.Getenv("TRAQ_CLIENT_ID")
	/*
		webhookSecret := os.Getenv("WEBHOOK_SECRET")
		webhookChannelId := os.Getenv("WEBHOOK_CHANNEL_ID")
		webhookId := os.Getenv("WEBHOOK_ID")
	*/
	return &Services{
		Auth: Auth{
			ClientID: traQClientID,
		},
		Storage: storage,
	}, nil
}
