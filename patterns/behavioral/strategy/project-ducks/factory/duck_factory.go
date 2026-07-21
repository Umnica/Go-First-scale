package factory

import (
	"patterns/behavioral/strategy/project-ducks/behavior/fly"
	"patterns/behavioral/strategy/project-ducks/behavior/quack"
	"patterns/behavioral/strategy/project-ducks/duck"
)

func CreateMallradDuck(name string) *duck.Duck {
	return duck.New(
		name,
		&fly.FlyWithWings{},
		&quack.NormalQuack{},
	)
}

func CreateRubberDuck(name string) *duck.Duck {
	return duck.New(
		name,
		&fly.FlyNoWay{},
		&quack.MuteQuack{},
	)
}

func CreateRocketDuck(name string) *duck.Duck {
	return duck.New(
		name,
		&fly.FlyWithRocket{},
		&quack.NormalQuack{},
	)
}
