package key

import (
	"sync"
)

var importKeyOnce sync.Once
var importKeyInstance *ImportKey

type ImportKey struct {
	PriArmor string
	PubArmor string
}

func GwKey() *ImportKey {
	importKeyOnce.Do(func() {
		importKeyInstance = &ImportKey{}
	})
	return importKeyInstance
}

func (i *ImportKey) Set(
	PriArmor string,
	PubArmor string) {

	i.PriArmor = PriArmor
	i.PubArmor = PubArmor
}

func (i *ImportKey) GetPriKey() (string) {
	return i.PriArmor
}

func (i *ImportKey) GetPubKey() (string) {
	return i.PubArmor
}