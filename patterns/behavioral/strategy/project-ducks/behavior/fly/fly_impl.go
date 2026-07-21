package fly

import "fmt"

// FlyWithWings - полет с крыльями (обычная утка)
type FlyWithWings struct{}

func (f FlyWithWings) Fly(duckName string) string {
	return fmt.Sprintf("Я %s лечу на крыльях!", duckName)
}

// FlyNoWay - не умеет летать (резиновая утка)
type FlyNoWay struct{}

func (f FlyNoWay) Fly(duckName string) string {
	return fmt.Sprintf("Я %s не умею летать :(", duckName)
}

// FlyWithRocket - полет с реактивным двигателем
type FlyWithRocket struct{}

func (f FlyWithRocket) Fly(duckName string) string {
	return fmt.Sprintf("Я %s лечу на реактивной тяге! 🚀", duckName)
}
