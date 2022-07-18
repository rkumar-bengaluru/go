package model

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewModel),
)

func NewModel() IOrder {
	var m IOrder
	m = &InMemOrder{}
	return m
}
