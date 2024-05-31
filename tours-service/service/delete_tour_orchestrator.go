package service

import (
	"database-example/saga/events"

	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type DeleteTourOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteTourOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteTourOrchestrator, error) {
	o := &DeleteTourOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// mozda ali mozda treba komanda prva da bude delete encounters jer
// on je vec mozda obriso ture pa salje sad da se brisu enkaunteri
func (o *DeleteTourOrchestrator) Start(tourId int64) error {
	event := &events.DeleteTourCommand{
		Type:   events.DeleteEncounter,
		TourId: tourId,
	}
	return o.commandPublisher.Publish(event)
}

func (o *DeleteTourOrchestrator) handle(reply *events.DeleteTourReply) {
	command := events.DeleteTourCommand{TourId: reply.TourId}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteTourOrchestrator) nextCommandType(reply events.DeleteTourReplyType) events.DeleteTourCommandType {
	switch reply {
	case events.TourDeleted:
		return events.DeleteEncounter
	case events.EncountersNotDeleted:
		return events.RollbackTourDelete
	case events.EncountersDeleted:
		return events.Confirm
	default:
		return events.UnknownCommand
	}
}
