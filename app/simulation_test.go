package app_test

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"stellartalk/app"
	"testing"
)

func init() {
	simapp.GetSimulatorFlags()
}

// Running as go benchmark test:
// `go test -benchmem -run=^$ -bench ^BenchmarkSimulation ./app -NumBlocks=200 -BlockSize 50 -Commit=true -Verbose=true -Enabled=true`
func BenchmarkSimulation(b *testing.B) {
	simapp.FlagEnabledValue = true
	simapp.FlagCommitValue = true

	config, db, dir, logger, _, err := simapp.SetupSimulation("goleveldb-newApp-sim", "Simulation")
	require.NoError(b, err, "simulation setup failed")

	b.Cleanup(func() {
		err := db.Close()
		if err != nil {
			log.Println("Error closing DB")
		}
		err = os.RemoveAll(dir)
		require.NoError(b, err)
	})

	encoding := app.MakeEncodingConfig()

	newApp := app.New(
		logger,
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)

	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		b,
		os.Stdout,
		newApp.BaseApp,
		simapp.AppStateFn(newApp.AppCodec(), newApp.SimulationManager()),
		simulationtypes.RandomAccounts,
		simapp.SimulationOperations(newApp, newApp.AppCodec(), config),
		newApp.ModuleAccountAddrs(),
		config,
		newApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(newApp, config, simParams)
	require.NoError(b, err)
	require.NoError(b, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}
}
