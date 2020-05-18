package hosts

import "fmt"
import "../workers"

type RoughcastHost struct{}

func (host RoughcastHost) TellMyStory() {
	host.introduce()
	buyARoughcastHouse()
	host.decorateHouse()
	moveToNewHouse()
}

func (host RoughcastHost) introduce() {
	fmt.Println("I'm the experienced host of roughcast house...")
}

func (host RoughcastHost) decorateHouse() {
	fmt.Println("I want to decorate my house...")
	hardwareWorker := findHardwareWorker()
	hardwareWorker.Decorate()

	floorFinisher := findFloorFinisher()
	floorFinisher.Decorate()

	kitchenAndBathWorker := findKitchenAndBathWorker()
	kitchenAndBathWorker.Decorate()

	paperhanger := findPaperHanger()
	paperhanger.Decorate()

	painter := findPainter()
	painter.Decorate()

	fmt.Println("Decorate house done...")
}

func findPainter() workers.Painter {
	fmt.Println("Find painter...")
	return workers.Painter{}
}

func findPaperHanger() workers.PaperHanger {
	fmt.Println("Find paper hanger...")
	return workers.PaperHanger{}
}

func findHardwareWorker() workers.HardwareWorker {
	fmt.Println("Find hardware worker...")
	return workers.HardwareWorker{}
}

func findKitchenAndBathWorker() workers.KitchenAndBathWorker {
	fmt.Println("Find kitchen and bath Worker...")
	return workers.KitchenAndBathWorker{}
}

func findFloorFinisher() workers.FloorFinisher {
	fmt.Println("Find floor finisher")
	return workers.FloorFinisher{}
}
