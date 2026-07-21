package main

import (
	"fmt"
	"patterns/behavioral/strategy/project-ducks/behavior/fly"
	"patterns/behavioral/strategy/project-ducks/factory"
)

func main() {
	fmt.Printf("### Паттерн Стратегия - Утки ###\n\n")

	rocketDuck := factory.CreateRocketDuck("РАКЕТНАЯ УТКА")

	rocketDuck.PerformFly()
	rocketDuck.PerformQuack()

	fmt.Printf("Меняем поведение Fly\n\n")
	rocketDuck.SetFlyBehavior(fly.FlyNoWay{})

	rocketDuck.PerformFly()
	rocketDuck.PerformQuack()
}
