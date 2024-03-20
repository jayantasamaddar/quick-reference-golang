package lib

import (
	"encoding/json"
	"fmt"
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

/****************************************************************************/
// (1a) Message Queue
/****************************************************************************/
// The Message Queue is the Mediator and contains topics
type MessageQueue struct {
	ID               string
	Name             string
	Topics           map[string]*Topic
	closed           bool
	unsubscriptionCh chan notifyUnsubscribe
}

func (m *MessageQueue) GetQueueEvents() chan notifyUnsubscribe {
	return m.unsubscriptionCh
}

// Initialize a new Message Queue
func NewMessageQueue(name string) *MessageQueue {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	return &MessageQueue{
		ID:               id,
		Name:             name,
		Topics:           map[string]*Topic{},
		closed:           false,
		unsubscriptionCh: make(chan notifyUnsubscribe),
	}
}

// Check if MessageQueue is closed
func (m *MessageQueue) IsClosed() bool {
	return m.closed
}

// Check if Topic exists in Message Queue
func (m *MessageQueue) TopicExists(t string) bool {
	if _, ok := m.Topics[t]; ok {
		return true
	}
	return false
}

// Get Topic
func (m *MessageQueue) GetTopic(topicName string) *Topic {
	for _, topic := range m.Topics {
		if topic.Name == topicName {
			return topic
		}
	}
	return nil
}

// Add topic to the MessageQueue if it exists. If it doesn't exist, skip.
func (m *MessageQueue) addTopic(t *Topic) {
	if !m.TopicExists(t.Name) {
		m.Topics[t.Name] = t
	}
}

// Add an Event to the Queue
func (m *MessageQueue) addEvent(event *Event) {
	// If topic doesn't exist, create a Topic
	if !m.TopicExists(event.TopicName) {
		topic := NewTopic(event.TopicName)

		// Add event to Topic
		topic.AddEvent(event)

		// Add topic to MessageQueue
		m.addTopic(topic)
		return
	}

	// Topic already exists
	for _, topic := range m.Topics {
		if topic.Name == event.TopicName {
			// Add event to topic
			topic.AddEvent(event)
		}
	}
}

// Add a Consumer to a Topic. If Topic doesn't exist, create a topic and then add Consumer.
func (m *MessageQueue) addConsumer(topicName string, c *Consumer) {
	// If topic doesn't exist
	if !m.TopicExists(topicName) {
		// Create topic
		topic := NewTopic(topicName)

		// Add topic to MessageQueue
		m.addTopic(topic)

		// Add event to Topic
		topic.AddConsumer(c)
		return
	}

	// If topic exists
	m.Topics[topicName].AddConsumer(c)
}

// Remove a Consumer from a Topic. If Topic doesn't exist, skip.
func (m *MessageQueue) removeConsumer(topicName string, c *Consumer) {
	// If topic doesn't exist, skip
	if m.TopicExists(topicName) {
		m.Topics[topicName].RemoveConsumer(c)
	}
}

/****************************************************************************/
// (1b) Topic: MessageQueue component
/****************************************************************************/
// The Topic
type Topic struct {
	Name      string
	consumers map[string]*Consumer
	events    chan Event
}

type notifyUnsubscribe struct {
	ConsumerID string
	TopicName  string
}

// Topic Constructor
func NewTopic(name string) *Topic {
	return &Topic{
		Name:      name,
		consumers: make(map[string]*Consumer),
		events:    make(chan Event),
	}
}

// AddEvent to the events channel
func (t *Topic) AddEvent(event *Event) {
	go func() {
		t.events <- *event
	}()
}

// Add Consumer to the Topic
func (t *Topic) AddConsumer(c *Consumer) {
	t.consumers[c.id] = c
}

// Remove Consumer from the Topic
func (t *Topic) RemoveConsumer(c *Consumer) {
	delete(t.consumers, c.id)
}

// Get an iterable channel that emits consumers
func (t *Topic) GetConsumers() chan<- *Consumer {
	ch := make(chan *Consumer)
	go func() {
		defer close(ch)
		for _, obs := range t.consumers {
			ch <- obs
		}
	}()
	return ch
}

/****************************************************************************/
// (2) Observable / Producer
/****************************************************************************/

// The Producer is the Observable
type Observable interface {
	Publish()
	JSON() string
}

// An Event is an Observable
type Event struct {
	Queue     *MessageQueue
	TopicName string
	Data      interface{}
}

func NewEvent(queue *MessageQueue, topic string, data interface{}) Observable {
	return &Event{queue, topic, data}
}

func (e *Event) Publish() {
	e.Queue.addEvent(e)
}

func (e *Event) JSON() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic("failed to marshal event")
	}
	return string(b)
}

func (e *Event) String() string {
	b, err := json.Marshal(e.Data)
	if err != nil {
		panic("failed to marshal event")
	}
	return fmt.Sprintf("%s\t{\"Topic\":\"%s\"}", string(b), e.TopicName)
	// return fmt.Sprintf("Event{ Topic: %s, Data: %+v }", e.TopicName, e.Data)
}

/****************************************************************************/
// (3) Observer / Consumer
/****************************************************************************/
// A Consumer is an Observer
type Observer interface {
	ID() string
	SetID(id string) Observer
	BindQueue(queue *MessageQueue) Observer
	Subscribe(topics ...string) Observer
	Unsubscribe(topics ...string) Observer
	isSubscribedTo(topicName string) (int, bool)
	GetTopics() []string
	Consume(topicName string, callback func(event Event), doneCh chan struct{})
	ConsumeAll(callback func(event Event), doneCh chan struct{})
}

// Consumer struct that is an Observer
type Consumer struct {
	id    string
	Name  string
	queue *MessageQueue
	// Topics consuming on
	topics []string
}

// Initialize New Consumer
func NewConsumer() Observer {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	return &Consumer{id: id}
}

// Returns the Customer ID
func (c *Consumer) ID() string {
	return c.id
}

// Sets the Customer ID
func (c *Consumer) SetID(id string) Observer {
	c.id = id
	return c
}

// Bind a Consumer to a Queue
func (c *Consumer) BindQueue(queue *MessageQueue) Observer {
	c.queue = queue
	return c
}

// Subscribe to a Topic in the queue
func (c *Consumer) Subscribe(topics ...string) Observer {
	for _, topic := range topics {
		c.queue.addConsumer(topic, c)
		c.topics = append(c.topics, topic)
	}
	return c
}

// Unsubscribe from topic
func (c *Consumer) Unsubscribe(topics ...string) Observer {
	var wg sync.WaitGroup

	for _, topic := range topics {
		if index, ok := c.isSubscribedTo(topic); ok {
			wg.Add(1) // Increment wait group for each topic
			c.queue.removeConsumer(topic, c)
			c.topics = append(c.topics[:index], c.topics[index+1:]...)

			go func(topic string) {
				defer wg.Done()
				c.queue.unsubscriptionCh <- notifyUnsubscribe{c.id, topic}
			}(topic)
		}
	}
	return c
}

// Private method to check if Consumer is subscribed to a Topic. Returns an index to the topic and an ok confirmation. Returns -1 as index if not found.
func (c *Consumer) isSubscribedTo(topicName string) (int, bool) {
	for index, t := range c.topics {
		if _, ok := c.queue.Topics[topicName]; ok && t == topicName {
			return index, true
		}
	}
	return -1, false
}

// Get all subscribed topics
func (c *Consumer) GetTopics() []string {
	return c.topics
}

// Goroutine to consume events from a subscribed topic. Listens for unsubscription notifications and modifies topic consumption on the fly.
func (c *Consumer) Consume(topicName string, callback func(event Event), doneCh chan struct{}) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case unsubscription := <-c.queue.unsubscriptionCh:
				fmt.Println("Unsubscription received:", unsubscription)
				if unsubscription.ConsumerID == c.id && unsubscription.TopicName == topicName {
					fmt.Println("Unsubscription received is valid (This is not being read):", unsubscription)
					if _, ok := c.isSubscribedTo(topicName); !ok {
						fmt.Println("This is not being read")
						break LOOP // Consumer is no longer subscribed, exit goroutine
					}
				}
			case event := <-c.queue.GetTopic(topicName).events:
				if _, ok := c.isSubscribedTo(topicName); ok {
					callback(event)
				}
			case <-doneCh:
				break LOOP // Exit goroutine when doneCh is closed
			}
		}
	}()
}

// Goroutine to Consume events for all subscribed Topics
func (c *Consumer) ConsumeAll(callback func(event Event), doneCh chan struct{}) {
	for topicName := range c.queue.Topics {
		go c.Consume(topicName, callback, doneCh)
	}
}
