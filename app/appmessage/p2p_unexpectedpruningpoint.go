package appmessage

// MsgUnexpectedPruningPoint represents a gor UnexpectedPruningPoint message
type MsgUnexpectedPruningPoint struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgUnexpectedPruningPoint) Command() MessageCommand {
	return CmdUnexpectedPruningPoint
}

// NewMsgUnexpectedPruningPoint returns a new gor UnexpectedPruningPoint message
func NewMsgUnexpectedPruningPoint() *MsgUnexpectedPruningPoint {
	return &MsgUnexpectedPruningPoint{}
}
