package types_test

import (
	"testing"

	"github.com/Finschia/finschia-sdk/store/iavl"
	"github.com/Finschia/finschia-sdk/store/rootmulti"
	storetypes "github.com/Finschia/finschia-sdk/store/types"
	"github.com/Finschia/ostracon/libs/log"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type MerkleTestSuite struct {
	suite.Suite

	store     *rootmulti.Store
	storeKey  *storetypes.KVStoreKey
	iavlStore *iavl.Store
}

func (suite *MerkleTestSuite) SetupTest() {
	db := dbm.NewMemDB()
	dblog := log.NewNopLogger()
	suite.store = rootmulti.NewStore(db, dblog)

	suite.storeKey = storetypes.NewKVStoreKey("iavlStoreKey")

	suite.store.MountStoreWithDB(suite.storeKey, storetypes.StoreTypeIAVL, nil)
	suite.store.LoadVersion(0)

	suite.iavlStore = suite.store.GetCommitStore(suite.storeKey).(*iavl.Store)
}

func TestMerkleTestSuite(t *testing.T) {
	suite.Run(t, new(MerkleTestSuite))
}
