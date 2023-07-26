package keeper_test

import (
	sdk "github.com/Finschia/finschia-sdk/types"

	"github.com/Finschia/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	expParams := types.DefaultParams()
	res, _ := suite.chainA.GetSimApp().ICAControllerKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().Equal(&expParams, res.Params)
}
