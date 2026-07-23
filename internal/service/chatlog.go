package service

import (
	"context"
)

func (svc *Service) ServiceStoreToChatLogDB(ctx context.Context, name string, text string) error {
	return svc.DBChatLog.StoreToChatLogDB(ctx, name, text)
}

func (svc *Service) ServiceReadFromChatLogDB(ctx context.Context, val int64) (string, error) {
	return svc.DBChatLog.ReadFromChatLogDB(ctx, val)
}
