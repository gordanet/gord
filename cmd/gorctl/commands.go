package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gordanet/gord/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.GordMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.GordMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.GordMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.GordMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.GordMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.GordMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.GordMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.GordMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.GordMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.GordMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.GordMessage_BanRequest{}),
	reflect.TypeOf(protowire.GordMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
