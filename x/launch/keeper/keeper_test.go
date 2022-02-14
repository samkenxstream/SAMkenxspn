package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/tendermint/spn/app"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *app.App
}
