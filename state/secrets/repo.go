// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package secrets

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/awishformore/cancoillotte/adaptor/database"
	"github.com/awishformore/cancoillotte/model"
)

type Repo struct {
	sync.Mutex
	kv      *database.KV
	secrets map[string]model.Secret
}

func NewRepo(log zerolog.Logger, kv *database.KV) (*Repo, error) {
	secrets := make(map[string]model.Secret)
	kvSecrets, err := kv.Secrets()
	if err != nil {
		return nil, errors.Wrap(err, "could not bootstrap secrets")
	}
	for _, secret := range kvSecrets {
		secrets[secret.ChaID] = secret
	}
	log.Info().Int("secrets", len(secrets)).Msg("secrets loaded from disk")
	r := &Repo{
		kv:      kv,
		secrets: secrets,
	}
	return r, nil
}

func (r *Repo) Backup(sec model.Secret) error {
	r.Lock()
	defer r.Unlock()
	key := fmt.Sprintf("%v-%d", sec.ChaID, sec.WizID)
	err := r.kv.Secret(sec)
	if err != nil {
		return errors.Wrap(err, "could not save secret")
	}
	r.secrets[key] = sec
	return nil
}

func (r *Repo) Recover(chaID string, wizID uint64) (model.Secret, error) {
	r.Lock()
	defer r.Unlock()
	key := fmt.Sprintf("%v-%d", chaID, wizID)
	sec, ok := r.secrets[key]
	if !ok {
		return model.Secret{}, errors.Errorf("missing challenge key (%s)", key)
	}
	return sec, nil
}
