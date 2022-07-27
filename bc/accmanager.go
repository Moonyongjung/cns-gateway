package bc

import (
	"strconv"
	"sync"

	"github.com/Moonyongjung/cns-gw/util"
)

var accSequenceInstance *AccSequenceStruct
var accSequenceOnce sync.Once

//-- Manage account Sequence
type AccSequenceStruct struct {
	AccountSequence string
}

func AccSequenceMng() *AccSequenceStruct {
	accSequenceOnce.Do( func() {
		accSequenceInstance = &AccSequenceStruct{}
	})
	return accSequenceInstance
}

func (n *AccSequenceStruct) NewAccSeq() {
	_, accountSequence := util.GetConfigAcc().Get()
	n.AccountSequence = accountSequence
}

func (n *AccSequenceStruct) NowAccSeq() string {
	return n.AccountSequence
}

func (n *AccSequenceStruct) AddAccSeq() {
	accountSequence := n.AccountSequence
	temp, err := strconv.Atoi(accountSequence)
	if err != nil {
		util.LogGw(err)
	}
	
	n.AccountSequence = strconv.Itoa(temp+1)
}