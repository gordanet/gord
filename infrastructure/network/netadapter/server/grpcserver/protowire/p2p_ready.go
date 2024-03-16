package protowire

import (
	"github.com/gordanet/gord/app/appmessage"
	"github.com/pkg/errors"
)

func (x *GordMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GordMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *GordMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
