package hosts

import "fmt"
import "../thirdParty"

type ExperiencedRoughcastHost struct{
}

func (host ExperiencedRoughcastHost)TellMyStory() {
	introduce()
	buyARoughcastHouse()
	host.decorateHouse()
	moveToNewHouse()
}

func (host ExperiencedRoughcastHost) decorateHouse() {
	fmt.Println("I want to decorate my house...")
	fmt.Println("I call the labor contractor...")

	hardwareWorker, _ := thirdParty.RecommendWorker("Hardware")
	hardwareWorker.Decorate()

	floorFinisher, _ := thirdParty.RecommendWorker("Floor")
	floorFinisher.Decorate()

	kitchenAndBathWorker, _ := thirdParty.RecommendWorker("KitchenAndBath")
	kitchenAndBathWorker.Decorate()

	paperhanger, _ := thirdParty.RecommendWorker("Wall")
	paperhanger.Decorate()

	painter, _ := thirdParty.RecommendWorker("Paint")
	painter.Decorate()

	fmt.Println("Decorate house done...")
}