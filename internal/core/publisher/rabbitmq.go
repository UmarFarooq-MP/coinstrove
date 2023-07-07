package publisher

import (
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type rabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQPublisher(amqpURL string) (ports.Publisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQPublisher{conn: conn, channel: channel}, nil
}

func (p *rabbitMQPublisher) Publish(message domain.Response) {
	body, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error While Marshaling in Publish Method with Message :%v", err)
		return
	}

	err = p.channel.Publish(
		"",
		"prices",
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		log.Fatalf("Error While publishing data in Publish Method with Message :%v", err)
		return
	}

}

func (p *rabbitMQPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}

func (p *rabbitMQPublisher) Init() {

	// We create a Queue to send the message to.
	_, err := p.channel.QueueDeclare(
		"prices", // name
		false,    // durable
		true,     // delete when unused
		false,    // exclusive
		true,     // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue with error : %v", err)
	}
}
