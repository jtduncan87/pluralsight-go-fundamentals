package concurrentProgrammingInGo

import (
	"fmt"
	"strings"
)

//BasicChannelMain Acts as both sender and receiver of a channel message
func BasicChannelMain() {
	ch := make(chan string, 1)
	ch <- "Roll Tide"
	fmt.Println(<-ch)
}

//BufferedChannelMain demonstrates buffered channels
func BufferedChannelMain() {
	phrase := "I'd rather be fishing.\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch, " ")
	}
}

//ClosingAChannel demonstrates closing a channel
func ClosingAChannel() {
	phrase := "I'd rather be fishing.\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println("The loop still prints because closing only closes the sending side\nIf we try to add something to it now, we'll get a panic:")
	ch <- "Roll Tide"
}

//RangingOverAChannel has a while loop that receives channel input
func RangingOverAChannel() {
	phrase := "I'd rather be fishing.\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	// this loop is checking the channel's status
	// for {
	// 	if msg, ok := <-ch; ok {
	// 		fmt.Print(msg, " ")
	// 	} else {
	// 		break
	// 	}
	// }
	// the same can be accomplished more succinctly here:
	for msg := range ch {
		fmt.Print(msg, " ")
	}
}

//SwitchBetweenChannelsExample demonstrates channel switching
func SwitchBetweenChannelsExample() {
	msgCh := make(chan message, 1)
	errCh := make(chan failedMessage, 1)

	/*msg := message{
		To:      []string{"dunscild@rohan.me"},
		From:    "strider@gondor.me",
		Content: "Gondor calls for aid!",
	}
	msgCh <- msg

	failedMsg := failedMessage{
		ErrorMessage:    "The beacons failed to light!",
		OriginalMessage: message{},
	}
	errCh <- failedMsg*/

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedErr := <-errCh:
		fmt.Println(receivedErr)
	default:
		fmt.Println("No messages in Middle Earth today")
	}
}

type message struct {
	To      []string
	From    string
	Content string
}
type failedMessage struct {
	ErrorMessage    string
	OriginalMessage message
}
