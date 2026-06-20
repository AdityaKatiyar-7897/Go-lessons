package main

import "fmt"

type Ball struct {
	x float64
	y float64
}

func main() {
	ball := Ball{
		x: 10,
		y: 20,
	}

	fmt.Println(ball.x)
	fmt.Println(ball.y)
}
