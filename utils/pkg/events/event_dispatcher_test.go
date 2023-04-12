package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (t *TestEvent) GetName() string {
	return t.Name
}

func (t *TestEvent) GetTimestamp() time.Time {
	return time.Now()
}

func (t *TestEvent) GetPayload() interface{} {
	return t.Payload
}

type TestEventHandler struct {
	ID int
}

func (t *TestEventHandler) Handle(data *Event) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      *TestEvent
	event2     *TestEvent
	handler    *TestEventHandler
	handler2   *TestEventHandler
	handler3   *TestEventHandler
	dispatcher *EventDispatcherImpl
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.handler = &TestEventHandler{}
	suite.handler2 = &TestEventHandler{}
	suite.handler3 = &TestEventHandler{}
	suite.event = &TestEvent{Name: "test", Payload: "test"}
	suite.event2 = &TestEvent{Name: "test2", Payload: "test"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.Register(suite.event.GetName(), suite.handler)
	_, ok := suite.dispatcher.handlers[suite.event.GetName()]

	suite.Nil(err)
	suite.True(ok)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))
	suite.Equal(suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterTwiceNotAllowed() {
	err := suite.dispatcher.Register(suite.event.GetName(), suite.handler)

	suite.Nil(err)

	err = suite.dispatcher.Register(suite.event.GetName(), suite.handler)

	suite.NotNil(err)
	suite.Equal("handler already registered", err.Error())
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.dispatcher.Register(suite.event.GetName(), suite.handler)
	suite.Nil(err)
	suite.True(suite.dispatcher.Has(suite.event.GetName(), suite.handler))
	suite.False(suite.dispatcher.Has(suite.event.GetName(), suite.handler2))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.dispatcher.Register(suite.event.GetName(), suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event2.GetName(), suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event2.GetName()]))

	suite.dispatcher.Clear()
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event.GetName()]))
	suite.Equal(0, len(suite.dispatcher.handlers[suite.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
