// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type Outcome struct {
	ID                      string   `json:"id"`
	WinnerWizardID          string   `json:"winnerWizardId"`
	PowerTransfer           string   `json:"powerTransfer"`
	HomeMoves               []string `json:"homeMoves"`
	AwayMoves               []string `json:"awayMoves"`
	HomeWizardFinalPower    string   `json:"homeWizardFinalPower"`
	AwayWizardFinalPower    string   `json:"awayWizardFinalPower"`
	HomeWizardStartingPower string   `json:"homeWizardStartingPower"`
	AwayWizardStartingPower string   `json:"awayWizardStartingPower"`
}
