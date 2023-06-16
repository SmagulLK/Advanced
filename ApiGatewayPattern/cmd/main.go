package main

import (
	"github.com/SmagulLK/APIGateway/pkg/auth"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"github.com/SmagulLK/APIGateway/pkg/product"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
	"log"
)

const (
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	queueName   = "msg"
)

type Message struct {
	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

type SendMessageResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

type MessengerService struct{}

func (s *MessengerService) SendMessage(ctx context.Context, message *Message) (*SendMessageResponse, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message.Body),
		},
	)
	if err != nil {
		return nil, err
	}

	return &SendMessageResponse{
		Success: true,
	}, nil
	response := &SendMessageResponse{
		Success: true,
	}
	return response, nil
}

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading config", err)

	}
	r := gin.Default()
	authService := auth.RegisterRouter(r, &c)
	product.ProductRoutes(r, &c, authService)
	r.Run(c.Port)

}
