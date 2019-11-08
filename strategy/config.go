// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package strategy

type Config struct {
	Reserve uint64
	Extra   uint64
	Ratio   Ratio
}

type Ratio struct {
	Fire    uint
	Water   uint
	Wind    uint
	Neutral uint
}
