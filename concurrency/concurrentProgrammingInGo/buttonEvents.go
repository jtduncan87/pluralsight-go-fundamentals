package concurrentProgrammingInGo

func makeButton() *button {
	result := new(button)
	result.eventListeners = make(map[string][]chan string)
	return result
}
func (btn *button) triggerEvent(event string, response string) {
	if _, present := btn.eventListeners[event]; present {
		for _, handler := range btn.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
func (btn *button) removeEventListener(event string, listenerCh chan string) {
	if _, present := btn.eventListeners[event]; present {
		for index := range btn.eventListeners[event] {
			if btn.eventListeners[event][index] == listenerCh {
				btn.eventListeners[event] = append(btn.eventListeners[event][:index], btn.eventListeners[event][index+1:]...)
				break
			}
		}
	}
}
func (btn *button) addEventListener(event string, responseCh chan string) {
	if _, present := btn.eventListeners[event]; present {
		btn.eventListeners[event] = append(btn.eventListeners[event], responseCh)
	} else {
		btn.eventListeners[event] = []chan string{responseCh}
	}
}

type button struct {
	eventListeners map[string][]chan string
}
