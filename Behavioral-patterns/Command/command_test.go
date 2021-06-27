package Command

import "testing"

// 命令——对象行为型模式
func TestCreateCommand(t *testing.T) {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}
