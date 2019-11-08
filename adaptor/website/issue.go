// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type IssueResponse struct {
	Data IssueData `json:"data"`
}

type IssueData struct {
	Challenge Challenge `json:"sendChallenge"`
}

const IssueBody = `{"operationName":"SendChallengeMutation","variables":{"homeWizardId":"%s","awayWizardId":"%s"},"query":"mutation SendChallengeMutation($homeWizardId: ID!, $awayWizardId: ID!) {\n  sendChallenge(input: {homeWizardId: $homeWizardId, awayWizardId: $awayWizardId}) {\n    id\n    homeWizardId\n    awayWizardId\n    __typename\n  }\n}\n"}`
