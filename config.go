// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package main

import "github.com/awishformore/cancoillotte/strategy"

type Config struct {
	Ethereum  Ethereum
	Badger    Badger
	Website   Website
	Alchemy   Alchemy
	Contracts Contracts
	Strategy  strategy.Config
}

type Ethereum struct {
	Endpoint string
	Keystore string
	Password string
}

type Badger struct {
	Datadir string
}

type Website struct {
	URL   string
	API   string
	Token string
}

type Alchemy struct {
	URL   string
	Token string
	Email string
}

type Contracts struct {
	Resolver    string
	Wizardguild string
	Gatekeeper  string
	Timekeeper  string
	Tournament  string
}
