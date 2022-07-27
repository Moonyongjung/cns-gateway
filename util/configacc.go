package util

import (
	"sync"
)

var acconce sync.Once
var accInstance *ConfigAcc

type ConfigAcc struct {
	AccountNumber    string
	AccountSequence  string	
}

func GetConfigAcc() *ConfigAcc {
	acconce.Do(func() {
		accInstance = &ConfigAcc{}
	})
	return accInstance
}

func (c *ConfigAcc) Set(
	accountNumber string,
	accountSequence string) {

	c.AccountNumber = accountNumber
	c.AccountSequence = accountSequence
}

func (c *ConfigAcc) Get() (string, string) {
	return c.AccountNumber, c.AccountSequence
}
