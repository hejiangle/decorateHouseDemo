package hosts

import "fmt"
import "../third_party"

type ExperiencedRoughcastHost struct {
}

func (host ExperiencedRoughcastHost) TellMyStory() {
	host.introduce()
	buyARoughcastHouse()
	host.decorateHouse()
	moveToNewHouse()
}

func (host ExperiencedRoughcastHost) introduce() {
	fmt.Println("I'm the experienced host of roughcast house...")
}

func (host ExperiencedRoughcastHost) decorateHouse() {
	fmt.Println("I want to decorate my house...")
	fmt.Println("I call the labor contractor...")

	hardwareWorker, _ := third_party.RecommendWorker("Hardware")
	hardwareWorker.Decorate()

	floorFinisher, _ := third_party.RecommendWorker("Floor")
	floorFinisher.Decorate()

	kitchenAndBathWorker, _ := third_party.RecommendWorker("KitchenAndBath")
	kitchenAndBathWorker.Decorate()

	paperhanger, _ := third_party.RecommendWorker("Wall")
	paperhanger.Decorate()

	painter, _ := third_party.RecommendWorker("Paint")
	painter.Decorate()

	fmt.Println("Decorate house done...")
}
