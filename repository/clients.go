package repository

import "context"

type Client interface {
	GetByID(context.Context, int64) ClientModel
	Get(context.Context) []ClientModel
}

type ClientModel struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

var _ Client = ClientMock{}

type ClientMock struct{}

func (cm ClientMock) GetByID(ctx context.Context, ID int64) ClientModel {
	return ClientModel{
		ID:   ID,
		Name: "El watton pelado",
		Mail: "wattonpelado@gomaster.org",
	}
}

func (cm ClientMock) Get(context.Context) []ClientModel {
	return []ClientModel{
		{
			ID:   1,
			Name: "El watton pelado",
			Mail: "wattonpelado@gomaster.org",
		},
		{
			ID:   2,
			Name: "El programador pobre",
			Mail: "pobreprog@gomaster.org",
		},
		{
			ID:   3,
			Name: "El bailarín de semáforos",
			Mail: "dancer@gomaster.org",
		},
		{
			ID:   4,
			Name: "El undercoder",
			Mail: "undercoder@gomaster.org",
		},
	}
}
