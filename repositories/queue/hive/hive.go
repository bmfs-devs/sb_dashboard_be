package hive

import (
	mHive "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Repository struct {
	Client mqtt.Client
}

func NewRepository(client mqtt.Client) *Repository {
	return &Repository{Client: client}
}

func (r *Repository) Publish(message mHive.HiveMessage) error {
	token := r.Client.Publish(message.Topic, 1, false, message.Payload)
	token.Wait()
	return token.Error()
}
