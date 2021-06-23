package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

func main() {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim ()


	//===========
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

	//=============

}


type Animal struct{}
func (r *Animal)Eat() {
	println("Eating")
}

type Shark struct{
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type SwimmerImpl struct{}
func (s *SwimmerImpl) Swim(){
	println("Swimming!")
}
type CompositeSwimmerB struct{
	Trainer
	Swimmer
}

