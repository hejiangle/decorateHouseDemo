package hosts

import (
	"../thirdParty"
	"fmt"
)

type HardcoverHost struct{}

func (host HardcoverHost) TellMyStory() {
	host.introduce()
	buyAHardcoverHouse(thirdParty.Decorate)
	moveToNewHouse()
}

func (host HardcoverHost) introduce(){
	fmt.Println("I'm the host of hardcover house...")
}

func buyAHardcoverHouse(decorate func()) {
	fmt.Println("I bought a hardcover house from developer...")
	fmt.Println("so that i don't need take care of decoration...")
	decorate()
}