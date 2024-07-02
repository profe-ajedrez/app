package repository

import "sync"

var container ContainerInt
var s *sync.Once = &sync.Once{}

type ContainerInt interface {
	Clients() Client
}

type Container struct {
	client Client
}

func GetContainer() ContainerInt {
	return newContainer()
}

func newContainer() ContainerInt {
	s.Do(func() {
		container = &Container{}
	})

	return container
}

func (c *Container) Clients() Client {
	if c.client == nil {
		c.client = ClientMock{}
	}

	return c.client
}
