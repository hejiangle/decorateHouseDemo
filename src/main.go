package main

import "fmt"
import "../src/hosts"

func main() {
	fmt.Println("Here is a story of roughcast host:")

	hosts.RoughcastHost{}.TellMyStory()

	fmt.Println("****************")

	fmt.Printf("Here is a story of smart roghcast host:")

	fmt.Println("****************")
}


