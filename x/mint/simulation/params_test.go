package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/mint/simulation"
)

func TestParamChangest(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"mint/InflationRateChange", "InflationRateChange", "\"0.230000000000000000\"", "mint"},
		{"mint/InflationMax", "InflationMax", "\"0.200000000000000000\"", "mint"},
		{"mint/InflationMin", "InflationMin", "\"0.070000000000000000\"", "mint"},
		{"mint/GoalBonded", "GoalBonded", "\"0.670000000000000000\"", "mint"},
		{"mint/DistributionProportions", "DistributionProportions", "{\"staking\":\"0.250000000000000000\",\"incentives\":\"0.210000000000000000\",\"community_pool\":\"0.540000000000000000\"}", "mint"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 5)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}

}
