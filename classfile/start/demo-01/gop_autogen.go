// Code generated by gop (Go+); DO NOT EDIT.

package main

import "fmt"

const _ = true

type Rect struct {
	Width  int
	Height int
}
//line Rect.gox:5:1
func (this *Rect) Area() int {
//line Rect.gox:6:1
	return this.Width * this.Height
}
//line hello.gop:1
func main() {
//line hello.gop:1:1
	rect := &Rect{20, 20}
//line hello.gop:2:1
	fmt.Println(rect.Area())
//line hello.gop:4:1
	fmt.Println(rect.Area() + rect.Area())
}
