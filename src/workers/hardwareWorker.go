package workers

import "fmt"

type HardwareWorker struct{}

func (w HardwareWorker) Decorate() {
	fmt.Println("Decorate window and balcony door...")
}
