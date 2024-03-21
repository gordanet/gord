package main

import (
	"context"
	"fmt"

	"github.com/gordanet/gord/cmd/gorwallet/daemon/client"
	"github.com/gordanet/gord/cmd/gorwallet/daemon/pb"
	"github.com/gordanet/gord/cmd/gorwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatGor(addressBalance.Available), utils.FormatGor(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, GOR %s %s%s\n", utils.FormatGor(response.Available), utils.FormatGor(response.Pending), pendingSuffix)

	return nil
}
