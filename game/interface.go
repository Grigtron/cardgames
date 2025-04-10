package game

type Game interface {
	HandleCommand(cmd string, args ...string) error
	Description() string
}