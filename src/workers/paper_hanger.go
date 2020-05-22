package workers

import "fmt"

type PaperHanger struct{}

func (w PaperHanger) Decorate() {
	fmt.Println("PaperHanger decorate wall...")
}
