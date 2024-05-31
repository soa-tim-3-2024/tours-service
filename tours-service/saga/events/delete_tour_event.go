package events

type DeleteTourCommandType int8

const (
	DeleteTour DeleteTourCommandType = iota
	RollbackTourDelete
	DeleteEncounter
	Confirm
	UnknownCommand
)

type DeleteTourCommand struct {
	TourId int64
	Type   DeleteTourCommandType
}

type DeleteTourReplyType int8

const (
	TourDeleted DeleteTourReplyType = iota
	EncountersDeleted
	EncountersNotDeleted
	UnknownReply
)

type DeleteTourReply struct {
	TourId int64
	Type   DeleteTourReplyType
}
