// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package model

import (
	"encoding/hex"
	"fmt"
	"sort"
)

type Secret struct {
	ChaID  string
	WizID  uint64
	Set    Set
	Commit Hash
	Salt   Hash
}

func (s Secret) Key(wizID1 uint64, wizID2 uint64) string {
	full := s.Set.Full()
	key := fmt.Sprintf("CW::revealKey::v1::%d::%d::%d::%x::%x",
		wizID1,
		wizID2,
		s.WizID,
		s.Salt,
		full,
	)
	return key
}

type Set [5]byte

func (s Set) Full() [32]byte {
	var full [32]byte
	copy(full[0:5], s[0:5])
	return full
}

func Compact(set [32]byte) Set {
	var compact Set
	copy(compact[0:5], set[0:5])
	return compact
}

func Distance(s1 Set, s2 Set) uint {
	var d, i, j uint
	sort.Sort(s1)
	sort.Sort(s2)
	for i < 5 && j < 5 {
		if s1[i] < s2[j] {
			i++
			d++
		} else if s1[i] > s2[j] {
			j++
			d++
		} else {
			i++
			j++
		}
	}
	return d
}

func (s Set) String() string {
	return hex.EncodeToString(s[:])
}

func (s Set) Distance(o Set) uint {
	d := uint(0)
	s1 := s
	s2 := o
	sort.Sort(s1)
	sort.Sort(s2)
	i := 0
	j := 0
	for i < 5 && j < 5 {
		if s1[i] < s2[j] {
			i++
			d++
		} else if s1[i] > s2[j] {
			j++
			d++
		} else {
			i++
			j++
		}
	}
	return d
}

func (s Set) Len() int {
	return 5
}

func (s Set) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s Set) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

type Hash [32]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

type SecretList []Secret
