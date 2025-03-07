package core

import (
	"testing"

	"github.com/ethereum-optimism/superchain-registry/superchain"

	"github.com/oswaldindex/op-geth/core/rawdb"
	"github.com/oswaldindex/op-geth/trie"
)

func TestOPStackGenesis(t *testing.T) {
	for id := range superchain.OPChains {
		gen, err := LoadOPStackGenesis(id)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("chain: %d, genesis block hash: %s", id, gen.ToBlock().Hash())
	}
}

func TestRegistryChainConfigOverride(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	genesis, err := LoadOPStackGenesis(10)
	if err != nil {
		t.Fatal(err)
	}
	if genesis.Config.RegolithTime == nil {
		t.Fatal("expected non-nil regolith time")
	}
	expectedRegolithTime := *genesis.Config.RegolithTime
	genesis.Config.RegolithTime = nil

	// initialize the DB
	tdb := trie.NewDatabase(db, newDbConfig(rawdb.PathScheme))
	genesis.MustCommit(db, tdb)
	bl := genesis.ToBlock()
	rawdb.WriteCanonicalHash(db, bl.Hash(), 0)
	rawdb.WriteBlock(db, bl)

	// create chain config, even with incomplete genesis input: the chain config should be corrected
	chainConfig, _, err := SetupGenesisBlockWithOverride(db, tdb, genesis, &ChainOverrides{
		ApplySuperchainUpgrades: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	// check if we have a corrected chain config
	if chainConfig.RegolithTime == nil {
		t.Fatal("expected regolith time to be corrected, but time is still nil")
	} else if *chainConfig.RegolithTime != expectedRegolithTime {
		t.Fatalf("expected regolith time to be %d, but got %d", expectedRegolithTime, *chainConfig.RegolithTime)
	}
}
