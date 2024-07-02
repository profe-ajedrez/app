package services

import (
	"context"

	"github.com/profe-ajedrez/app/repository"
)

type ClientResponse struct {
	Data  []repository.ClientModel `json:"data"`
	Count int                      `json:"count"`
}

func GetClients(c context.Context) ClientResponse {
	container := repository.GetContainer()
	cli := container.Clients().Get(c)
	count := len(cli)

	return ClientResponse{Data: cli, Count: count}
}
