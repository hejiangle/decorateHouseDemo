package thirdParty

import (
	"../workers"
	"errors"
	"fmt"
)

func RecommendWorker(workType string) (workers.Worker, error) {
	switch workType {
	case "Hardware":
		fmt.Println("HardwareWorker call me...")
		return workers.HardwareWorker{}, nil
	case "Floor":
		fmt.Println("FloorFinisher call me...")
		return workers.FloorFinisher{}, nil
	case "KitchenAndBath":
		fmt.Println("KitchenAndBathWorker call me...")
		return workers.KitchenAndBathWorker{}, nil
	case "Wall":
		fmt.Println("KitchenAndBathWorker call me...")
		return workers.PaperHanger{}, nil
	case "Paint":
		fmt.Println("KitchenAndBathWorker call me...")
		return workers.Painter{}, nil
	default:
		return nil, errors.New("no this type worker in our labor market")
	}
}
