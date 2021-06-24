package Composite

import "testing"

// 组合
func TestAthlete_Train(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestAnimal_Eat(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestSwimmerImpl_Swim(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}
