package State

import "testing"

// 状态——对象行为型模式
func TestState(t *testing.T)  {
	start := StartState{}
	game := GameContext{
		Next:&start,
	}
	for game.Next.executeState(&game) {}
}