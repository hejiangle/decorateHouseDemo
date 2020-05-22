package workers

import "fmt"

type Painter struct{}

func (w Painter) Decorate() {
	fmt.Println("Painter decorate door and kick line...")
}
