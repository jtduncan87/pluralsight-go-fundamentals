package concurrentProgrammingInGo

import (
	"errors"
)

type promise struct {
	successCh chan interface{}
	failureCh chan error
}

func savePOWithPromise(po *purchaseOrder, shouldFail bool) *promise {
	result := new(promise)
	result.successCh = make(chan interface{}, 1)
	result.failureCh = make(chan error, 1)
	go func() {
		if shouldFail {
			result.failureCh <- errors.New("Failed to save purchase order")
		} else {
			po.Number = 17
			result.successCh <- po
		}
	}()
	return result
}

func (this *promise) then(success func(interface{}) error, failure func(error)) *promise {
	result := new(promise)
	result.successCh = make(chan interface{}, 1)
	result.failureCh = make(chan error, 1)
	go func() {
		select {
		case obj := <-this.successCh:
			newErr := success(obj)
			if newErr == nil {
				result.successCh <- obj
			} else {
				result.failureCh <- newErr
			}
		case err := <-this.failureCh:
			failure(err)
			result.failureCh <- err
		}
	}()
	return result
}
