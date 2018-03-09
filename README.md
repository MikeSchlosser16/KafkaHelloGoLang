# Using KStreams in Java a lot, want to test Confluent's Golang Client for Apache Kafka.


### Build / Run
```
cd $HOME/go/src/helloKafka
go build
./helloKafka
```

### Dependencies
```
brew install librdkafka
brew install pkg-config
go get -u github.com/confluentinc/confluent-kafka-go/kafka
```

##### References
* https://github.com/confluentinc/confluent-kafka-go
* https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies
* https://kafka.apache.org/documentation/
* https://golang.org/pkg/fmt/
* https://tour.golang.org/concurrency/1
* https://gobyexample.com/goroutines
* https://golangbot.com/goroutines/
* https://golang.org/doc/effective_go.html#blank
