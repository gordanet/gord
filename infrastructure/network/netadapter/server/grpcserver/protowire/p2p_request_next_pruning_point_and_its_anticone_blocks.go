package protowire

import (
	"github.com/gordanet/gord/app/appmessage"
	"github.com/pkg/errors"
)

func (x *GordMessage_RequestNextPruningPointAndItsAnticoneBlocks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GordMessage_DonePruningPointAndItsAnticoneBlocks is nil")
	}
	return &appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks{}, nil
}

func (x *GordMessage_RequestNextPruningPointAndItsAnticoneBlocks) fromAppMessage(_ *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks) error {
	return nil
}
