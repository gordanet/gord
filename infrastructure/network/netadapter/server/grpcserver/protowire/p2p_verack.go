package protowire

import (
	"github.com/gordanet/gord/app/appmessage"
	"github.com/pkg/errors"
)

func (x *GordMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GordMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *GordMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
