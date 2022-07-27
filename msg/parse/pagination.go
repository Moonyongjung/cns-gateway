package parse

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)
func defaultPagination() *query.PageRequest{
	return &query.PageRequest {
		Key: []byte(""),
		Offset: 0,
		Limit: 0,
		CountTotal: false,
		Reverse: false,
	}
}