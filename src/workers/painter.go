package workers

import "fmt"

type Painter struct{}

func (w Painter) Decorate() {
	fmt.Println("Decorate door and kick line...")
}
