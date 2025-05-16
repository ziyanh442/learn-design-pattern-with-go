package subject

import (
	"2observer/observer"
	"fmt"
	"slices"
	"time"
)

type SubjectInterface interface {
	Register(observerInterface observer.ObserverInterface, topicName string)
	Deregister(observerInterface observer.ObserverInterface, topicName string)
	NotifySubscribers(topic, message string)
	// Display weather report every 1 second for a duration
	ReportWeather(duration time.Duration)
	// Display crime report every 5 second for a duration
	ReportCrime(duration time.Duration)
}

type Subscribers struct {
	observer observer.ObserverInterface
	topics   []string
}

type Subject struct {
	subscribers     []Subscribers
	availableTopics []string
}

func NewSubject() SubjectInterface {
	return &Subject{
		availableTopics: []string{
			"weather",
			"crime",
		},
	}
}

func (s *Subject) Register(observerInterface observer.ObserverInterface, topicName string) {
	if !slices.Contains(s.availableTopics, topicName) {
		fmt.Printf("topic %s is unavailable \n", topicName)
		return
	}

	for index, subscriber := range s.subscribers {
		if subscriber.observer.GetObserverName() != observerInterface.GetObserverName() {
			continue
		}

		if slices.Contains(subscriber.topics, topicName) {
			return
		}

		s.subscribers[index].topics = append(s.subscribers[index].topics, topicName)
		fmt.Printf("%s subscribed to %s report \n", observerInterface.GetObserverName(), topicName)
		return
	}

	s.subscribers = append(s.subscribers, Subscribers{
		observer: observerInterface,
		topics:   []string{topicName},
	})
	fmt.Printf("%s subscribed to %s report \n", observerInterface.GetObserverName(), topicName)
}

func (s *Subject) Deregister(observerInterface observer.ObserverInterface, topicName string) {
	for index, subscriber := range s.subscribers {
		if subscriber.observer.GetObserverName() != observerInterface.GetObserverName() {
			continue
		}

		if !slices.Contains(subscriber.topics, topicName) {
			return
		}

		if len(subscriber.topics) == 1 {
			s.subscribers = append(s.subscribers[:index], s.subscribers[index+1:]...)
			break
		}

		for topicIndex, topic := range subscriber.topics {
			if topic == topicName {
				s.subscribers[index].topics = append(subscriber.topics[:topicIndex], subscriber.topics[topicIndex+1:]...)
				break
			}
		}

		fmt.Printf("%s unsubscribed from %s report \n", observerInterface.GetObserverName(), topicName)
		return
	}
}

func (s *Subject) NotifySubscribers(topic, message string) {
	for _, subscriber := range s.subscribers {
		if !slices.Contains(subscriber.topics, topic) {
			continue
		}

		subscriber.observer.Display(topic, message)
	}
}
