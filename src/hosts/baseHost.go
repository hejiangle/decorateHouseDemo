package hosts

import "fmt"

type host interface {
	TellMyStory()
	decorateHouse()
}

func introduce(){
	fmt.Println("I'm the host of roughcast house...")
}

func buyARoughcastHouse() {
	fmt.Println("I bought a roughcast house...")
}

func moveToNewHouse(){
	fmt.Println("Now I can live in new house...")
}