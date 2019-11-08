// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type ChallengesResponse struct {
	Data ChallengesData `json:"data"`
}

type ChallengesData struct {
	Challenges Challenges `json:"myChallenges"`
}

type Challenges struct {
	Sent     []Challenge `json:"sent"`
	Received []Challenge `json:"received"`
	Active   []Challenge `json:"active"`
}

const ChallengesBody = `{"operationName":"MyDuelsQuery","variables":{},"query":"query MyDuelsQuery {\n  myChallenges {\n    sent {\n      id\n      homeWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      awayWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      status\n      createdAt\n      lastActionAt\n      outcome {\n        id\n        winnerWizardId\n        powerTransfer\n        homeMoves\n        awayMoves\n        homeWizardFinalPower\n        awayWizardFinalPower\n        homeWizardStartingPower\n        awayWizardStartingPower\n        __typename\n      }\n      __typename\n    }\n    received {\n      id\n      homeWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      awayWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      status\n      createdAt\n      lastActionAt\n      outcome {\n        id\n        winnerWizardId\n        powerTransfer\n        homeMoves\n        awayMoves\n        homeWizardFinalPower\n        awayWizardFinalPower\n        homeWizardStartingPower\n        awayWizardStartingPower\n        __typename\n      }\n      __typename\n    }\n    active {\n      id\n      homeWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      awayWizard {\n        id\n        name\n        ownerAddress\n        affinityType\n        imageUrl\n        power\n        owner {\n          id\n          nickname\n          avatarURL\n          twitterAvatarURL\n          status\n          __typename\n        }\n        __typename\n      }\n      status\n      createdAt\n      lastActionAt\n      outcome {\n        id\n        winnerWizardId\n        powerTransfer\n        homeMoves\n        awayMoves\n        homeWizardFinalPower\n        awayWizardFinalPower\n        homeWizardStartingPower\n        awayWizardStartingPower\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}`
