// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package database

import (
	"encoding/binary"
	"encoding/json"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"

	"github.com/awishformore/cancoillotte/model"
)

type KV struct {
	db *badger.DB
}

func NewKV(db *badger.DB) *KV {
	k := &KV{
		db: db,
	}
	return k
}

func (k *KV) Duel(duel model.Duel) error {
	key := duel.ID[:]
	err := k.set(DuelCode, key, duel)
	if err != nil {
		return errors.Wrap(err, "could not save duel")
	}
	return nil
}

func (k *KV) Duels() (model.DuelList, error) {
	var duels model.DuelList
	err := k.scan(DuelCode, func(key []byte, val []byte) error {
		var duel model.Duel
		err := json.Unmarshal(val, &duel)
		if err != nil {
			return errors.Wrap(err, "could not decode duel")
		}
		duels = append(duels, duel)
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve duels")
	}
	return duels, nil
}

func (k *KV) Wizard(wiz model.Wizard) error {
	key := make([]byte, 8)
	binary.LittleEndian.PutUint64(key, wiz.ID)
	err := k.set(WizardCode, key, wiz)
	if err != nil {
		return errors.Wrap(err, "could not save wizard")
	}
	return nil
}

func (k *KV) Wizards() (model.WizardList, error) {
	var wizards model.WizardList
	err := k.scan(WizardCode, func(key []byte, val []byte) error {
		var wizard model.Wizard
		err := json.Unmarshal(val, &wizard)
		if err != nil {
			return errors.Wrap(err, "could not decode wizard")
		}
		wizards = append(wizards, wizard)
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve wizards")
	}
	return wizards, nil
}

func (k *KV) Secret(sec model.Secret) error {
	key := []byte(sec.ChaID)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, sec.WizID)
	key = append(key, buf...)
	err := k.set(SecretCode, key, sec)
	if err != nil {
		return errors.Wrap(err, "could not save secret")
	}
	return nil
}

func (k *KV) Secrets() (model.SecretList, error) {
	var secrets model.SecretList
	err := k.scan(SecretCode, func(key []byte, val []byte) error {
		var secret model.Secret
		err := json.Unmarshal(val, &secret)
		if err != nil {
			return errors.Wrap(err, "could not decode secret")
		}
		secrets = append(secrets, secret)
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve wizards")
	}
	return secrets, nil
}

func (k *KV) Player(owner string, player string) error {
	key := []byte(owner)
	err := k.set(PlayerCode, key, player)
	if err != nil {
		return errors.Wrap(err, "could not save player")
	}
	return nil
}

func (k *KV) Players() (map[string]string, error) {
	players := make(map[string]string)
	var player string
	err := k.scan(PlayerCode, func(key []byte, val []byte) error {
		owner := string(key[1:])
		err := json.Unmarshal(val, &player)
		if err != nil {
			return errors.Wrap(err, "could not decode player")
		}
		players[owner] = player
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve players")
	}
	return players, nil
}

func (k *KV) set(code uint8, key []byte, v interface{}) error {
	key = append([]byte{code}, key...)
	err := k.db.Update(func(tx *badger.Txn) error {
		val, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = tx.Set(key, val)
		return err

	})
	if err != nil {
		return errors.Wrap(err, "could not save data")
	}
	return nil
}

func (k *KV) scan(code uint8, next func(key []byte, val []byte) error) error {
	prefix := []byte{code}
	err := k.db.View(func(tx *badger.Txn) error {
		it := tx.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			key := item.Key()
			err := item.Value(func(val []byte) error {
				err := next(key, val)
				return err
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "could not scan code")
	}
	return nil
}
