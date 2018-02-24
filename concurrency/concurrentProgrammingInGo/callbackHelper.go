package concurrentProgrammingInGo

type purchaseOrder struct {
	Number int
	Value  float64
}

func savePO(po *purchaseOrder, callback chan *purchaseOrder) {
	po.Number = 23

	callback <- po
}
