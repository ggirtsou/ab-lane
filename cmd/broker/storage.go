package main

// Storage is responsible for interacting with different data store implementations.
type Storage interface {
	// Insert saves a serialized Message in storage
	Insert(topic string, message Message) error
	// Get streams messages for the specified topic and consumer group.
	Get(topic string, consumerGroup string) (Message, error)
}

// Message represents how
type Message struct {
	Attributes map[string]interface{}
	Payload    []byte
}
