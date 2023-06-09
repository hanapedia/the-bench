package core

// EgressAdapter provides common interface for all the external service resouce.
// Example resources include:
// - REST API routes
// - gRPC methods
// - Kafka topic
// - Database table
//
// It is intended to represent the individual endpoints on each exteranl service,
// not the services themselves; hence the name `EgressAdapter`
type EgressAdapter interface {
	Call() (string, error)
}

// Used to reuse connections to other serivces
// Wrapper interface, so the struct to implement this should have pointer to actual connection
// types to implement this interface should have some sort of pointer to connections 
type EgressConnection interface {
	Close()
}

type EgressAdapterError struct {
	EgressAdapter *EgressAdapter
	Error             error
}
