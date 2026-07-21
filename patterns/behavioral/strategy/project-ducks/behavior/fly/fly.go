package fly

type FlyBehavior interface {
	Fly(duckName string) string
}
