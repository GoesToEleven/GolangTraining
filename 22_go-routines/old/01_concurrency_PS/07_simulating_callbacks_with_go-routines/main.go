package main

import (
	"fmt"
)

type purchaseOrder struct {
	Number int
	Value  float64
}

func savePO(po *purchaseOrder, callbackChannel chan *purchaseOrder) {
	po.Number = 1234

	callbackChannel <- po
}

func main() {
	po := new(purchaseOrder)
	po.Value = 42.27

	ch := make(chan *purchaseOrder)

	go savePO(po, ch)

	newPo := <-ch
	fmt.Printf("PO Number: %d\n", newPo.Number)
}
