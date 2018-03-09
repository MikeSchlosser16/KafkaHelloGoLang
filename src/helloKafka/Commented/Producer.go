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

  /* A goroutine is a lightweight thread managed by the Go runtime.
  go f(x,y,z) starts a new goroutine running f(x,y,z).
  They are functions that run concurrently with other functions,
  Basically threads, but much more light weight and tiny... */

  /* This means here, the below function will run concurrently with the main function
  which lets us send events and then check if they've failed. Main also runs in
  its own Goroutine called the main Goroutine.
  -- When new Goroutine started, goroutine call returns immediately
  -- Main Goroutine should be running for any other Goroutines to run
     (dont let it die first) */
  go func(){
    for event := range producer.Events(){ // range iterates over elements in a variety of data structures, so in this case over each event
      switch ev := event.(type){ // Switch in Go
      case *kafka.Message:
        if ev.TopicPartition.Error != nil {
          fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
        } else {
          fmt.Printf("Successfully delievered message to: %v\n", ev.TopicPartition)
        }
      }
    }
  }()


  // Produce messages to topic we are listening on with consumer (asynchronously)
  topic := "myTopic"
  /* Blank identifier - The use of a blank identifier in a for range loop is a
  special case of a general situation: multiple assignment.
  Avoids creating a dummy variable, and makes it clear those values can be discarded.
  range on arrays and slices provides both the index and value for each entry.
  Below we don't need the index, so we ignored it with the blank identifier _.
  Only care about the word */
  for _, word := range []string{"Hello", "this", "is", "a", "golang", "kafka",
    "learning", "experience"}{
      producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{ Topic: &topic, Partition: kafka.PartitionAny},
        Value:         []byte(word),
      }, nil)}


    // Waiting for message deliveries
    producer.Flush(15 * 1000)
}
