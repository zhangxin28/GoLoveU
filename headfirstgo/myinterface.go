package main

/*
An interface is a set of methods certain values are expected to have.
Any type that has all the methods listed in an interface definition
is said to satisfy that interface.
A type that satisfies an interface can be assigned to any variable
or function parameter that uses that interface as its type.
An interface type is an abstract type. Interfaces don't describe
what a value is: they don's say what its underlying type is or how
its data is stored. They only describe what a value can do:
what methods it has.
When you have a variable of an interface type, the only methods
you can call on it are those defined in the interface.
If you'ave assigned a value of a concrete type to a variable with an
interface type, you can use `type assertion` to get the concrete
type value back.
*/

import "fmt"

// MyInterfacer is a interface
type MyInterfacer interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

//MyType is a type which underlying type is int
type MyType int

// MethodWithoutParameters runs an interface method
func (m *MyType) MethodWithoutParameters() {
	m.MethodNotInInterface()
	fmt.Printf("MethodWithoutParameters called: %d\n", *m)
}

// MethodWithParameter runs an interface method
func (m *MyType) MethodWithParameter(f float64) {
	m.MethodNotInInterface()
	fmt.Printf("MethodWithParameter called: %f\n", float64(*m)*f)
}

// MethodWithReturnValue runs an interface method
func (m *MyType) MethodWithReturnValue() string {
	m.MethodNotInInterface()
	return fmt.Sprintf("Hi from MethodWithReturnValue: %d\n", *m)
}

// MethodNotInInterface runs a normal method not in interface
func (m *MyType) MethodNotInInterface() {
	fmt.Printf("MethodNotInInterface called: %d\n", *m)
}

func testMyInterface(mi MyInterfacer) {
	mi.MethodWithParameter(5.0)
	fmt.Print(mi.MethodWithReturnValue())
	mi.MethodWithoutParameters()
}

// NoiseMaker represents an interface
type NoiseMaker interface {
	MakeSound()
}

// Whistle represents a new type which underlying type is string
type Whistle string

// MakeSound repsents the method for Whistle.MakeSound
func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

// Horn represents a new type which underlying type is tring
type Horn string

func newHorn(value string) *Horn {
	h := Horn(value)
	return &h
}

// MakeSound represents the method of Horn.MakeSound
func (h *Horn) MakeSound() {
	fmt.Println("Honk!")
}

// Robot represents a new type which underlying type is string
type Robot string

func newRobot(value string) *Robot {
	r := Robot(value)
	return &r
}

// MakeSound represents the method of Robot.MakeSound
func (r *Robot) MakeSound() {
	fmt.Println("Beep Boop!")
}

// Walk represents the method of Robot.Walk
func (r *Robot) Walk() {
	fmt.Println("Powering legs!")
}

func play(n NoiseMaker) {
	n.MakeSound()
}

type player interface {
	play(string)
	stop()
}

type tapePlayer struct {
	batteries string
}

func (t tapePlayer) play(song string) {
	fmt.Println("Playing", song)
}

func (t tapePlayer) stop() {
	fmt.Println("Stopped!")
}

type tapeRecorder struct {
	microphones int
}

func (t tapeRecorder) play(song string) {
	fmt.Println("Playing", song)
}

func (t tapeRecorder) stop() {
	fmt.Println("Stopped!")
}

func playList(device player, songs []string) {
	for _, song := range songs {
		device.play(song)
	}

	device.stop()
}

type comedyError string

// Error represents a method for comedyError
func (c comedyError) Error() string {
	return string(c)
}

type overheatError float64

// Error represents a method for overheatError
func (o overheatError) Error() string {
	return fmt.Sprintf("Overheating by %.2f degress", o)
}

// RunMyInterface tests defining a type that satisfies an interface
func RunMyInterface() {
	var mi MyInterfacer
	m := MyType(5)
	n := m + 10
	mi = &m
	fmt.Printf("n = %d, n > 15: %t\n", n, n > 15)
	testMyInterface(mi)

	var noiseMaker NoiseMaker

	noiseMaker = newRobot("New Robot")
	play(noiseMaker)
	/*
		below is called `type assertion`,
		like saying "I know the noisemaker uses the interface
		type NoiseMaker, but I'm pretty sure this NoiseMaker
		is actually a pointer points to a Robot."
	*/
	r, ok := noiseMaker.(*Robot)
	// use the ok value to determine whether i can safely
	// convert to the right type, here is `*Robot`
	if ok {
		// method `Walk` is not part of NoiseMaker,
		// just for the Robot
		r.Walk()
	}

	noiseMaker = newHorn("New Horn")
	play(noiseMaker)
	if h, ok := noiseMaker.(*Horn); ok {
		h.MakeSound()
	}

	noiseMaker = Whistle("New Whistle")
	play(noiseMaker)
	if w, ok := noiseMaker.(Whistle); ok {
		w.MakeSound()
	}

	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var p player = tapePlayer{}
	playList(p, mixtape)
	p = tapeRecorder{}
	playList(p, mixtape)

	/*
		error here is an interface
		eg:

		type error interface {
			Error() string
		}
	*/
	var err error = comedyError("This is a comedy error")
	fmt.Println(err)

	err = func(actual float64, safe float64) error {
		excess := actual - safe
		if excess > 0 {
			return overheatError(excess)
		}
		return nil
	}(121.379, 100.0)

	if err != nil {
		fmt.Println(err)
	}
}
