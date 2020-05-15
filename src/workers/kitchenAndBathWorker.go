package workers

import "fmt"

type KitchenAndBathWorker struct {}

func (w KitchenAndBathWorker)Decorate() {
	fmt.Println("Decorate kitchen and bath...")
}