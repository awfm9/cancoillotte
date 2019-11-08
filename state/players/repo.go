// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package players

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/awishformore/cancoillotte/adaptor/database"
	"github.com/awishformore/cancoillotte/adaptor/website"
)

type Repo struct {
	sync.Mutex
	kv      *database.KV
	web     *website.Client
	players map[string]string
}

func NewRepo(log zerolog.Logger, kv *database.KV, web *website.Client) (*Repo, error) {
	players, err := kv.Players()
	if err != nil {
		return nil, errors.Wrap(err, "could not load players")
	}
	set := make(map[string]struct{})
	for _, player := range players {
		set[player] = struct{}{}
	}
	log.Info().Int("owners", len(players)).Int("players", len(set)).Msg("players loaded from disk")
	r := &Repo{
		kv:      kv,
		web:     web,
		players: players,
	}
	return r, nil
}

func (r *Repo) Get(owner string, wizID uint64) (string, error) {
	r.Lock()
	defer r.Unlock()
	player, ok := r.players[owner]
	if ok {
		return player, nil
	}
	player, err := r.web.Player(wizID)
	if err == nil {
		err = r.kv.Player(owner, player)
		if err != nil {
			return "", errors.Wrap(err, "could not store player")
		}
	}
	r.players[owner] = player
	return player, nil
}
