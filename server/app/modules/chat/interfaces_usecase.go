package chat

import (
	"context"

	entities "github.com/timPrachasri/app-chat/server/app/entities"
)

type Usecase interface {
	AddChat(ctx context.Context, message entities.Message) error
	ReadChat(ctx context.Context, roomName string) ([]*entities.Message, error)
}
