package main

import (
	"fmt"
	"os"

	"github.com/catspa3/catspad/util"

	"github.com/catspa3/catspad/version"

	"github.com/pkg/errors"

	_ "net/http/pprof"

	"github.com/catspa3/catspad/infrastructure/os/signal"
	"github.com/catspa3/catspad/util/panics"
	"github.com/catspa3/catspad/util/profiling"
)

func main() {
	defer panics.HandlePanic(log, "MAIN", nil)
	logoCats()

	interrupt := signal.InterruptListener()

	cfg, err := parseConfig()
	if err != nil {
		printErrorAndExit(errors.Errorf("Error parsing command-line arguments: %s", err))
	}
	defer backendLog.Close()

	fmt.Println(`:3 Welcome to Cats Network! We thank you for running a node and supporting this project.`)
	fmt.Printf(`:3 This Miner address is` + " \033[32;5m %s \033[0m\n", cfg.MiningAddr)
	if cfg.NetworkFlags.Testnet {
		fmt.Println(`:3 Now,starting a client node on Cats Testnet`)
	} else if cfg.NetworkFlags.Simnet {
		fmt.Println(`:3 Now,starting a client node on Cats Simnet`)
	} else if cfg.NetworkFlags.Devnet {
		fmt.Println(`:3 Now,starting a client node on Cats Devnet`)
	} else {
		fmt.Println(`:3 Now,starting a client node on Cats Mainnet`)
	}

	// Show version at startup.
	log.Infof("Version %s", version.Version())

	// Enable http profiling server if requested.
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	client, err := newMinerClient(cfg)
	if err != nil {
		panic(errors.Wrap(err, "error connecting to the RPC server"))
	}
	defer client.Disconnect()

	miningAddr, err := util.DecodeAddress(cfg.MiningAddr, cfg.ActiveNetParams.Prefix)
	if err != nil {
		printErrorAndExit(errors.Errorf("Error decoding mining address: %s", err))
	}

	doneChan := make(chan struct{})
	spawn("mineLoop", func() {
		err = mineLoop(client, cfg.NumberOfBlocks, *cfg.TargetBlocksPerSecond, cfg.MineWhenNotSynced, miningAddr)
		if err != nil {
			panic(errors.Wrap(err, "error in mine loop"))
		}
		doneChan <- struct{}{}
	})

	select {
	case <-doneChan:
	case <-interrupt:
	}
}

func printErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	os.Exit(1)
}
