package mock

import (
	"context"
	"net/http"

	"github.com/morkid/paginate"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
)

type NFTBalanceRepository struct {
	nftBalances []*eventindexer.NFTBalance
}

func NewNFTBalanceRepository() *NFTBalanceRepository {
	return &NFTBalanceRepository{}
}

func (r *NFTBalanceRepository) IncreaseAndDecreaseBalancesInTx(
	ctx context.Context,
	increaseOpts eventindexer.UpdateNFTBalanceOpts,
	decreaseOpts eventindexer.UpdateNFTBalanceOpts,
) (increasedBalance *eventindexer.NFTBalance, decreasedBalance *eventindexer.NFTBalance, err error) {
	return nil, nil, nil
}

func (r *NFTBalanceRepository) FindByAddress(ctx context.Context,
	req *http.Request,
	address string,
	chainID string,
) (paginate.Page, error) {
	var balances []*eventindexer.NFTBalance

	for _, b := range r.nftBalances {
		if b.Address == address {
			balances = append(balances, b)
		}
	}

	return paginate.Page{
		Items: balances,
	}, nil
}
