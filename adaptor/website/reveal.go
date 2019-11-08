// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type RevealResponse struct {
	Data RevealData `json:"data"`
}

type RevealData struct {
	Challenge Challenge `json:"revealMoves"`
}

const RevealBody = `{"operationName":"RevealMovesMutation","variables":{"wizardId":"%s","challengeID":"%s","key":"%s","moveSet":"%s"},"query":"mutation RevealMovesMutation($challengeID: String!, $wizardId: String!, $moveSet: String!, $key: String!) {\n  revealMoves(input: {challengeID: $challengeID, wizardId: $wizardId, moveSet: $moveSet, key: $key}) {\n    id\n    homeWizardId\n    awayWizardId\n    __typename\n  }\n}\n"}`
