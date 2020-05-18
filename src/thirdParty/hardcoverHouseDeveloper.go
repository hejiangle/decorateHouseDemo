package thirdParty

import "../workers"

func Decorate() {
	workerTeam := append(
		[]workers.Worker{},
		workers.HardwareWorker{},
		workers.KitchenAndBathWorker{},
		workers.PaperHanger{},
		workers.FloorFinisher{},
		workers.Painter{})

	for _, w := range workerTeam {
		w.Decorate()
	}
}
