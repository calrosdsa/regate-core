package repository

import "context"

type Notification struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	EntityId   int        `json:"entity_id"`
	TypeEntity TypeEntity `json:"type_entity"`
	Read       bool       `json:"read"`
	Image      string     `json:"image"`
	ProfileId  int        `json:"profile_id"`
	CreatedAt  string     `json:"created_at,omitempty"`
}

type RequestNotificationDiffusion struct {
	Notification Notification `json:"notification"`
	Categories   []int        `json:"categories"`
	// Cities       []string     `json:"cities"`
}

type NotificationRepository interface {
	// SendNotification(ctx context.Context,d Notification)(err error)
}

type NotificationuseCase interface {
	SendNotification(ctx context.Context, d RequestNotificationDiffusion) (err error)
}
