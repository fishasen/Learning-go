package main

import "fmt"

type Painter interface {
	Draw()
}

type Xiaowang struct {
}

func (Xiaowang) Draw() {
	fmt.Println("I am drawing a paper.")
}

func main() {
	var xw Xiaowang
	var painter Painter = xw
	painter.Draw()
	var painter2 Painter = &xw
	painter2.Draw()
	painter3 := Painter(xw)
	painter3.Draw()
	painter4 := Painter(&xw)
	painter4.Draw()
}
