package workers

import "fmt"

type KitchenAndBathWorker struct{}

func (w KitchenAndBathWorker) Decorate() {
	fmt.Println("KitchenAndBathWorker decorate kitchen and bath...")
}
