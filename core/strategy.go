// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package core

import "github.com/awishformore/cancoillotte/model"

type Strategy interface {
	OnBlock(hash string, height uint64) error
	OnCreation(wizID uint64, aff uint8, pow uint64) error
	OnTransfer(wizID uint64, from string, to string) error
	OnChallenge(wizID1 uint64, wizID2 uint64, nonce1 uint64, nonce2 uint64, commit1 model.Hash) error
	OnWithdrawal(wizID uint64) error
	OnFight(duelID model.Hash, wizID1 uint64, wizID2 uint64, timeout uint64, contest bool, commit1 model.Hash, commit2 model.Hash) error
	OnReveal(duelID model.Hash, wizID1 uint64, wizID2 uint64) error
	OnResult(duelID model.Hash, wizID1 uint64, wizID2 uint64, set1 model.Set, set2 model.Set, pow1 uint64, pow2 uint64) error
	OnTimeout(duelID model.Hash, wizID1 uint64, wizID2 uint64, pow1 uint64, pow2 uint64) error
	OnElimination(wizID uint64) error
	OnAttempt(wizID uint64) error
	OnPair(wizID1 uint64, wizID2 uint64) error
	OnContest(wizID1 uint64, wizID2 uint64, commit2 model.Hash) error
	OnSuccess(wizID uint64, pow uint64) error
	OnRevival(wizID uint64, pow uint64) error
	OnWipe(wizID1 uint64, wizID2 uint64, pow1 uint64, reason uint8) error
	OnClaim(wizID uint64, prize uint64) error
}
