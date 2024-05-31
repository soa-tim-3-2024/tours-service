package handler

import (
	"database-example/saga"
	"database-example/saga/events"
	"database-example/service"
)

type DeleteTourCommandHandler struct {
	tourService       *service.TourService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteTourCommandHandler(service *service.TourService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteTourCommandHandler, error) {
	o := &DeleteTourCommandHandler{
		tourService:       service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteTourCommandHandler) handle(command *events.DeleteTourCommand) {
	reply := events.DeleteTourReply{TourId: command.TourId}
	switch command.Type {
	case events.RollbackTourDelete:
		err := handler.tourService.RollbackTourDelete(command.TourId)
		if err != nil {
			return
		}
		reply.Type = events.UnknownReply
	case events.Confirm:
		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
