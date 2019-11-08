// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package main

import (
	"context"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/awishformore/cancoillotte/adaptor/alchemy"
	"github.com/awishformore/cancoillotte/adaptor/database"
	"github.com/awishformore/cancoillotte/adaptor/gatekeeper"
	"github.com/awishformore/cancoillotte/adaptor/resolver"
	"github.com/awishformore/cancoillotte/adaptor/tournament"
	"github.com/awishformore/cancoillotte/adaptor/website"
	"github.com/awishformore/cancoillotte/adaptor/wizardguild"
	"github.com/awishformore/cancoillotte/core"
	"github.com/awishformore/cancoillotte/state/duels"
	"github.com/awishformore/cancoillotte/state/players"
	"github.com/awishformore/cancoillotte/state/secrets"
	"github.com/awishformore/cancoillotte/state/wizards"
	"github.com/awishformore/cancoillotte/state/world"
	"github.com/awishformore/cancoillotte/strategy"
)

// TODO: implement alternative moveset selection strategy (exploitative);
//		 polarize range between strongest wins and movesets that win against
//		 what the other player would select against those movesets
// TODO: implement ascension strategy
// TODO: add support for multiple key stores
// TODO: re-sync duels and wizards after resolution window
// TODO: gift power when maximum challenge slot wizards reached
// TODO: weight outgoing challenges by responsiveness of player

func main() {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)

	var config string
	var simulate bool
	pflag.StringVarP(&config, "cfg", "c", "cancoillotte.yml", "path to config file")
	pflag.BoolVarP(&simulate, "sim", "s", false, "whether to only simulate challenges")
	pflag.Parse()

	log.Info().Msg("cancoillotte starting")

	data, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal().Err(err).Msg("could not read config file")
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse config file")
	}

	data, err = ioutil.ReadFile(cfg.Ethereum.Keystore)
	if err != nil {
		log.Fatal().Err(err).Msg("could not read keystore file")
	}

	key, err := keystore.DecryptKey(data, cfg.Ethereum.Password)
	if err != nil {
		log.Fatal().Err(err).Msg("could not decode private key")
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)

	opts := badger.DefaultOptions(cfg.Badger.Datadir)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal().Err(err).Msg("could not open badger database")
	}

	cli, err := ethclient.Dial(cfg.Ethereum.Endpoint)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect ethereum client")
	}

	web, err := website.NewClient(cfg.Website.URL, cfg.Website.API, cfg.Website.Token)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize websites client")
	}

	alch, err := alchemy.NewClient(cfg.Alchemy.URL, cfg.Alchemy.Token, cfg.Alchemy.Email)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize alchemy client")
	}

	reso, err := resolver.NewBinding(cfg.Contracts.Resolver, cli)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize resolver binding")
	}

	guild, err := wizardguild.NewBinding(cfg.Contracts.Wizardguild, cli)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize wizardguild binding")
	}

	gate, err := gatekeeper.NewBinding(cfg.Contracts.Gatekeeper, cli, auth)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize gatekeeper binding")
	}

	tour, err := tournament.NewBinding(cfg.Contracts.Tournament, cli, auth)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize tournament binding")
	}

	kv := database.NewKV(db)

	play, err := players.NewRepo(log, kv, web)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize players repository")
	}

	wiz, err := wizards.NewRepo(log, tour, guild, play)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize wizard repository")
	}

	duel, err := duels.NewRepo(log, alch, kv, wiz)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize duel repository")
	}

	sec, err := secrets.NewRepo(log, kv)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize commit repository")
	}

	info, err := world.NewInfo(cli, tour, duel, play, auth.From)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize world info")
	}

	ctx, cancel := context.WithCancel(context.Background())

	strat, err := strategy.NewStandard(log, ctx, guild, gate, tour, reso, wiz, duel, sec, play, info, web, key, cfg.Strategy, simulate)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize standard strategy")
	}

	sub, err := core.NewSubscriber(log, cli, guild.Contract(), tour.Contract(), strat)
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize core subscriber")
	}

	log.Info().Msg("cancoillotte started")

	<-sig

	log.Info().Msg("cancoillotte stopping")

	go func() {
		<-sig
		os.Exit(1)
	}()

	cancel()

	err = db.Close()
	if err != nil {
		log.Warn().Err(err).Msg("could not close badger database")
	}

	log.Info().Msg("database persisted")

	sub.Stop()

	log.Info().Msg("cancoillotte stopped")

	os.Exit(0)
}
