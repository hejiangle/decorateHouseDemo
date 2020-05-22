package workers

import "fmt"

type HardwareWorker struct{}

func (w HardwareWorker) Decorate() {
	fmt.Println("HardwareWorker decorate window and balcony door...")
}
