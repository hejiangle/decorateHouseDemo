package workers

import "fmt"

type FloorFinisher struct{}

func (w FloorFinisher) Decorate() {
	fmt.Println("FloorFinisher decorate floor...")
}
