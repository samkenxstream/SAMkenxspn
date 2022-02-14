package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// Setup initializes a new App
func Setup(isCheckTx bool) *App {
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()
	encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)

	cmdApp := New(logger, db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, encoding, simapp.EmptyAppOptions{})

	var a *App
	switch c := cmdApp.(type) {
	case *App:
		a = c
	default:
		panic("simapp imported incorrectly")
	}

	if !isCheckTx {
		genesisState := NewDefaultGenesisState(encoding.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		a.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)

	}

	return a
}
