package main

import(
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main(){
  fmt.Printf("Hello, world\n")

  consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
      "boostrap.servers": "localhost",
      "group.id": "myGroup",
      "auto.offset.reset": "earliest"})

  if err != nil{
    panic(err)
  }

  consumer.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

  for {
    msg, err := consumer.ReadMessage(-1)
    if err == nil {
      fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
    } else {
      fmt.Printf("Consumer error: %v (%v)\n", err, msg)
  }

  consumer.Close()
}
