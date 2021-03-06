package main

import (
	"io"
	"log"
)

// NoopStore does nothing with battery saves
type NoopStore struct {
}

func NewNoopStore() *NoopStore {
	return new(NoopStore)
}

func (n *NoopStore) Open(game string) (io.ReadCloser, error) {
	return new(NoopContent), nil
}

func (n *NoopStore) Create(game string) (io.WriteCloser, error) {
	return new(NoopContent), nil
}

type NoopContent struct{}

func (c *NoopContent) Write(b []byte) (int, error) {
	log.Println("Saving to a NOOP store")
	log.Println(string(b))

	return len(b), nil
}

func (c *NoopContent) Read(p []byte) (int, error) {
	return len(p), nil
}

func (c *NoopContent) Close() error {
	return nil
}
