package usecase

import (
	"context"
	"encoding/json"
	"log"
	r "regate-core/domain/repository"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

type salaUseCase struct {
	salaRepo r.SalaRepository
	utilU r.UtilUseCase
	timeout time.Duration
	kafkaW       *kafka.Writer
}

func NewSalaUseCase(salaRepo r.SalaRepository,utilU r.UtilUseCase,timeout time.Duration) r.SalaUseCase{
	w := &kafka.Writer{
		Addr:     kafka.TCP(viper.GetString("kafka.url")),
		Topic:    "sala",
		Balancer: &kafka.LeastBytes{},
	}
	return &salaUseCase{
		salaRepo: salaRepo,
		utilU: utilU,
		kafkaW: w,
		timeout: timeout,
	}
}
func (u *salaUseCase)DisabledExpiredRooms(ctx context.Context){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	res,err := u.salaRepo.DisabledExpiredRooms(ctx)
	if err != nil {
		u.utilU.LogError("GetDepositosEmpresa","billing_usecase.go",err.Error())
		return 
	}
	log.Println(res)
	for _,salaId := range res {
		go u.sendSalaNotification(salaId)
	}
}

func (u *salaUseCase)DeleteUnAvailablesSalas(ctx context.Context){
	ctx,cancel := context.WithTimeout(ctx,u.timeout)
	defer cancel()
	u.salaRepo.DeleteUnAvailablesSalas(ctx)
}

func (u *salaUseCase)sendSalaNotification(salaId int){
	payload := r.MessageNotification {
		EnitityId: salaId,
		Title: "La vigencia de tu sala ha caducado.",
		Message: "La sala ha sido deshabilitada debido a que excedió el tiempo límite.",
	}
	json, err := json.Marshal(payload)
	if err != nil {
		u.utilU.LogError("sendSalaNotification_json.Marshal", "sala_use_ucase.go", err.Error())
	}
	err = u.kafkaW.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("sala"),
			Value: json,
		},
	)
	if err != nil {
		u.utilU.LogError("sendSalaNotification_kafka", "sala_use_case.go", err.Error())
	}
}