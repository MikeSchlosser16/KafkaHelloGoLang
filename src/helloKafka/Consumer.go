package main

import(
  "fmt" //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
  "github.com/confluentinc/confluent-kafka-go/kafka" //confluent-kafka-go is Confluent's Golang client for Apache Kafka

)

func main(){
  fmt.Printf("Hello, world\n")

  /*Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
    -- & here creates a pointer of the ConfigMap object
    -- Go has a built in error type. This checks if we fail to instantiate consumer, we then throw an error*/
  consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
      "boostrap.servers": "localhost",
      "group.id": "myGroup", // Specifies the name of the consumer group a Kafka consumer belongs to
      "auto.offset.reset": "earliest"}) // What to do when there is no initial offset in Kafka


  if err != nil{
    panic(err) /* Built in function that stops ordinary flow and begins panicking.
      A panic typically means something went unexpectedly wrong. Mostly we use
      it to fail fast on errors that shouldn’t occur during normal operation,
      or that we aren’t prepared to handle gracefully. A common use of panic
      is to abort if a function returns an error value that we don’t know how
      to (or want to) handle. */
  }

  // Subscribing our consumer to an array of topics, one that matches a reg expr
  consumer.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

  for {
    msg, err := consumer.ReadMessage(-1)
    // No error, consumer recieved message
    if err == nil {
      fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
    } else {
      fmt.Printf("Consumer error: %v (%v)\n", err, msg) //%v is a generic placeholder. It will automatically convert your variable into a string with some default options.
    }
  }

  consumer.Close()
}
