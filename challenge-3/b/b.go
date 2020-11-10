package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"github.com/wesleywillians/go-rabbitmq/queue"
)

type Order struct {
	ID       uuid.UUID
	Coupon   string
	CcNumber string
}

func NewOrder() Order {
	return Order{ID: uuid.NewV4()}
}

const (
	InvalidCoupon   = "invalid"
	ValidCoupon     = "valid"
	ConnectionError = "connection error"
)

func init() {
	err := godotenv.Load("/Users/elisson/workspaces/avanca-dev/filas/b/.env")

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()

	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	for msg := range messageChannel {
		process(msg)
	}
}

type Result struct {
	Status string
}

func process(msg amqp.Delivery) {
	order := NewOrder()
	json.Unmarshal(msg.Body, &order)

	resultCoupon := makeHTTPCall("http://localhost:9092", order.Coupon)

	switch resultCoupon.Status {
	case InvalidCoupon:
		log.Println("Order: ", order.ID, ": invalid coupon!")
	case ConnectionError:
		msg.Reject(false)
		log.Println("Order: ", order.ID, ": could not process the messsage!, Connection Error")
	case ValidCoupon:
		log.Println("Order: ", order.ID, ": Processed!")

	}
}

func makeHTTPCall(urlMicroservice string, coupon string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)

	res, err := http.PostForm(urlMicroservice, values)

	if err != nil {
		result := Result{Status: ConnectionError}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		result := Result{Status: "Servidor fora do ar"}
		return result
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result
}
