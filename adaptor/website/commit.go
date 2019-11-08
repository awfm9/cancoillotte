// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type CommitResponse struct {
	Data CommitData `json:"data"`
}

type CommitData struct {
	Challenge Challenge `json:"submitAndSignMoves"`
}

const CommitBody = `{"operationName":"SubmitAndSignMovesMutation","variables":{"input":{"walletAddress":"%s","currentWizardId":"%s","challengeId":"%s","commit":"%s","signature":"0x%s"}},"query":"mutation SubmitAndSignMovesMutation($input: SubmitAndSignDuelMovesInput!) {\n  submitAndSignMoves(input: $input) {\n    id\n    homeWizardId\n    awayWizardId\n    status\n    __typename\n  }\n}\n"}`
