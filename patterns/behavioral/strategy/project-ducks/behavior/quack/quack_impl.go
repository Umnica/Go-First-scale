package quack

// NormalQuack - обычное кряканье
type NormalQuack struct{}

func (q NormalQuack) Quack() string {
	return "Кря-кря!"
}

// Squeak - писк (резиновая утка)
type Squeak struct{}

func (q Squeak) Quack() string {
	return "Пи-пи-пи!"
}

// MuteQuack - молчит
type MuteQuack struct{}

func (q MuteQuack) Quack() string {
	return "..."
}
