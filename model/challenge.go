// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

type Challenge struct {
	ID      string
	Wiz1    uint64
	Wiz2    uint64
	Owner1  string
	Owner2  string
	Player1 string
	Player2 string
	Online1 bool
	Online2 bool
	Status  string
}

type ChallengeFilter func(r Challenge) bool

type ChallengeList []Challenge

func (cl ChallengeList) Filter(filter ChallengeFilter) ChallengeList {
	var dup ChallengeList
	for _, c := range cl {
		if filter(c) {
			dup = append(dup, c)
		}
	}
	return dup
}

func (cl ChallengeList) Swap(filter ChallengeFilter) ChallengeList {
	for i, c := range cl {
		if filter(c) {
			cl[i].Wiz1, cl[i].Wiz2 = cl[i].Wiz2, cl[i].Wiz1
			cl[i].Owner1, cl[i].Owner2 = cl[i].Owner2, cl[i].Owner1
		}
	}
	return cl
}

func (cl ChallengeList) HomeIDs() []uint64 {
	set := make(map[uint64]struct{})
	for _, c := range cl {
		set[c.Wiz1] = struct{}{}
	}
	wizIDs := make([]uint64, 0, len(set))
	for wizID := range set {
		wizIDs = append(wizIDs, wizID)
	}
	return wizIDs
}

func (cl ChallengeList) AwayIDs() []uint64 {
	set := make(map[uint64]struct{})
	for _, c := range cl {
		set[c.Wiz2] = struct{}{}
	}
	wizIDs := make([]uint64, 0, len(set))
	for wizID := range set {
		wizIDs = append(wizIDs, wizID)
	}
	return wizIDs
}

func (cl ChallengeList) Mapping() map[string]string {
	mapping := make(map[string]string)
	for _, c := range cl {
		mapping[c.Owner1] = c.Player1
		mapping[c.Owner2] = c.Player2
	}
	return mapping
}

func (cl ChallengeList) Online() []string {
	set := make(map[string]struct{})
	for _, c := range cl {
		if c.Online1 {
			set[c.Player1] = struct{}{}
		}
		if c.Online2 {
			set[c.Player2] = struct{}{}
		}
	}
	players := make([]string, 0, len(set))
	for player := range set {
		players = append(players, player)
	}
	return players
}
