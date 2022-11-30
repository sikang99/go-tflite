package main

import (
	"fmt"
	"log"

	"github.com/sikang99/go-tflite"
)

func bin(n int, num_digits int) []float32 {
	f := make([]float32, num_digits)
	for i := uint(0); i < uint(num_digits); i++ {
		f[i] = float32((n >> i) & 1)
	}
	return f[:]
}

func dec(b []float32) int {
	for i := 0; i < len(b); i++ {
		if b[i] > 0.4 {
			return i
		}
	}
	panic("Sorry, I'm wrong")
}

func display(v []float32, i int) {
	switch dec(v) {
	case 0:
		fmt.Println(i)
	case 1:
		fmt.Println("Fizz")
	case 2:
		fmt.Println("Buzz")
	case 3:
		fmt.Println("FizzBuzz")
	}
}

func main() {
	model := tflite.NewModelFromFile("fizzbuzz_model.tflite")
	if model == nil {
		log.Println("cannot load model")
		return
	}
	defer model.Delete()

	interpreter := tflite.NewInterpreter(model, nil)
	defer interpreter.Delete()

	interpreter.AllocateTensors()

	for i := 1; i <= 100; i++ {
		buf := bin(i, 7)
		interpreter.GetInputTensor(0).CopyFromBuffer(&buf[0])
		interpreter.Invoke()
		display(interpreter.GetOutputTensor(0).Float32s(), i)
	}
}
