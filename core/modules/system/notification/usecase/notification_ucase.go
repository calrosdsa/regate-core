package usecase

import (
	"context"
	r "regate-core/domain/repository"
	"time"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"

)

type notificationU struct {
	timeout time.Duration
	notificationR r.NotificationRepository
	utilU r.UtilUseCase
	kafkaW      *kafka.Writer
}

func NewUseCase(timeout time.Duration,notificationR r.NotificationRepository,utilU r.UtilUseCase) r.NotificationuseCase{
	w := &kafka.Writer{
		Addr:     kafka.TCP(viper.GetString("kafka.host")),
		Topic:    "notification-diffusion",
		Balancer: &kafka.LeastBytes{},
	}
	return &notificationU{
		timeout: timeout,
		notificationR: notificationR,
		utilU: utilU,
		kafkaW: w,
	}
}

func (u *notificationU)SendNotification(ctx context.Context,d r.RequestNotificationDiffusion)(err error){
	// ctx,cancel := context.WithTimeout(ctx,u.timeout)
	// defer cancel()
	if d.Notification.Image == "" {
		d.Notification.Image = "https://cdn-icons-png.flaticon.com/64/4239/4239989.png"
	}
	go u.utilU.SendMessageToKafka(u.kafkaW,d,"notification")
	return
}

func (u *notificationU) GetInfoText(ctx context.Context,id int)(res r.InfoText,err error){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	// res,err = u.notificationU.GetInfoText(ctx,id)
	if err != nil {
		u.utilU.LogError("GetInfoText","info_usecase",err.Error())
	}
	return
}


