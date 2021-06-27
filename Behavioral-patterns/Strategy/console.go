package Strategy

type ConsoleSquare struct{}

func (c *ConsoleSquare) Print() error {
	println("Square")
	return nil
}
