package observer

import (
	"fmt"
	"time"
)

type ObserverInterface interface {
	GetObserverName() string
	Display(topic, message string)
}

type Observer struct {
	name string
}

func NewObserver(name string) ObserverInterface {
	return &Observer{
		name: name,
	}
}

func (o *Observer) GetObserverName() string {
	return o.name
}

func (o *Observer) Display(topic, message string) {
	fmt.Printf("%s - [%s] %s: %s \n", time.Now().Format(time.ANSIC), o.name, topic, message)
}
