package duck

import (
	"fmt"
	"patterns/behavioral/strategy/project-ducks/behavior/fly"
	"patterns/behavioral/strategy/project-ducks/behavior/quack"
)

type Duck struct {
	Name          string
	FlyBehavior   fly.FlyBehavior
	QuackBehavior quack.QuackBehavior
}

func New(name string, flyBehavior fly.FlyBehavior, quackBehavior quack.QuackBehavior) *Duck {
	return &Duck{
		Name:          name,
		FlyBehavior:   flyBehavior,
		QuackBehavior: quackBehavior,
	}
}

func (d *Duck) SetFlyBehavior(fb fly.FlyBehavior) {
	d.FlyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb quack.QuackBehavior) {
	d.QuackBehavior = qb
}

func (d *Duck) PerformFly() {
	if d.FlyBehavior == nil {
		fmt.Println("Ошибка: поведение полёта не установлено")
	}

	fmt.Println(d.FlyBehavior.Fly(d.Name))
}

func (d *Duck) PerformQuack() {
	if d.QuackBehavior == nil {
		fmt.Println("Ошибка: поведение кряканья не установлено")
	}

	fmt.Println(d.QuackBehavior.Quack())
}
