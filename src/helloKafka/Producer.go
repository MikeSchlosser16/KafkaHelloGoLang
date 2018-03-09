package main

import (
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
  producer, err := kafka.NewProducer(&kafka.ConfigMap{
      "bootstrap.servers" : "localhost"})
  if err != nil{
    panic(err)
  }

  go func(){
    for event := range producer.Events(){
      switch ev := event.(type){
      case *kafka.Message:
        if ev.TopicPartition.Error != nil {
          fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
        } else {
          fmt.Printf("Successfully delievered message to: %v\n", ev.TopicPartition)
        }
      }
    }
  }()

  topic := "myTopic"
  for _, word := range []string{"Hello", "this", "is", "a", "golang", "kafka",
    "learning", "experience"}{
      producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{ Topic: &topic, Partition: kafka.PartitionAny},
        Value:         []byte(word),
      }, nil)}

    producer.Flush(15 * 1000)
}
