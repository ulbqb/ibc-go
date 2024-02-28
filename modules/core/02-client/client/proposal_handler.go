package client

import (
	govclient "github.com/Finschia/finschia-sdk/x/gov/client"

	"github.com/cosmos/ibc-go/v4/modules/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
)
