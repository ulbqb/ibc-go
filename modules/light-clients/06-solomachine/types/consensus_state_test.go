package types_test

import (
	"github.com/Finschia/ibc-go/v4/modules/core/exported"
	"github.com/Finschia/ibc-go/v4/modules/light-clients/06-solomachine/types"
	ibctesting "github.com/Finschia/ibc-go/v4/testing"
)

func (suite *SoloMachineTestSuite) TestConsensusState() {
	consensusState := suite.solomachine.ConsensusState()

	suite.Require().Equal(exported.Solomachine, consensusState.ClientType())
	suite.Require().Equal(suite.solomachine.Time, consensusState.GetTimestamp())
	suite.Require().Nil(consensusState.GetRoot())
}

func (suite *SoloMachineTestSuite) TestConsensusStateValidateBasic() {
	// test singlesig and multisig public keys
	for _, solomachine := range []*ibctesting.Solomachine{suite.solomachine, suite.solomachineMulti} {

		testCases := []struct {
			name           string
			consensusState *types.ConsensusState
			expPass        bool
		}{
			{
				"valid consensus state",
				solomachine.ConsensusState(),
				true,
			},
			{
				"timestamp is zero",
				&types.ConsensusState{
					PublicKey:   solomachine.ConsensusState().PublicKey,
					Timestamp:   0,
					Diversifier: solomachine.Diversifier,
				},
				false,
			},
			{
				"diversifier is blank",
				&types.ConsensusState{
					PublicKey:   solomachine.ConsensusState().PublicKey,
					Timestamp:   solomachine.Time,
					Diversifier: " ",
				},
				false,
			},
			{
				"pubkey is nil",
				&types.ConsensusState{
					Timestamp:   solomachine.Time,
					Diversifier: solomachine.Diversifier,
					PublicKey:   nil,
				},
				false,
			},
		}

		for _, tc := range testCases {
			tc := tc

			suite.Run(tc.name, func() {
				err := tc.consensusState.ValidateBasic()

				if tc.expPass {
					suite.Require().NoError(err)
				} else {
					suite.Require().Error(err)
				}
			})
		}
	}
}
