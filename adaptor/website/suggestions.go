// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type SuggestionResponse struct {
	Data SuggestionData `json:"data"`
}

type SuggestionData struct {
	Challenges []Challenge `json:"getQuickMatches"`
}

const SuggestionBody = `{"operationName":"DuelBlockPage","variables":{"wizardId": null},"query": "query DuelBlockPage($wizardId: ID) {\n  myProfile {\n    user {\n      id\n      wizards {\n        ...WizardFragments\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  getQuickMatches(wizardId: $wizardId) {\n    id\n    awayWizardId\n    homeWizardId\n    homeWizard {\n      ...WizardFragments\n      __typename\n    }\n    awayWizard {\n      ...WizardFragments\n      __typename\n    }\n    status\n    __typename\n  }\n  tournament: currentTournament {\n    address\n    tournamentWindow: window\n    blueWallPower\n    nextBlueWallPower\n    calendar: fightWindows {\n      tournamentWindow: window\n      start\n      windowStartBlock\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment WizardFragments on Wizard {\n  id\n  name\n  ownerAddress\n  affinityType\n  imageUrl\n  power\n  status\n  owner {\n    avatarURL\n    id\n    nickname\n    status\n    twitterAvatarURL\n    twitterHandle\n    wallets {\n      address\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n"}`
