package main

import "fmt"
import "../src/hosts"

func main() {
	fmt.Println("Here is a story of roughcast host:")

	hosts.RoughcastHost{}.TellMyStory()

	fmt.Println("****************")

	fmt.Println("Here is a story of experienced roughcast host:")

	hosts.ExperiencedRoughcastHost{}.TellMyStory()

	fmt.Println("****************")
}


