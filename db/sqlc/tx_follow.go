package social

import "context"

type FollowsTransParam struct {
	FollowUserParams
	AfterFollow func(follow Follow) error
}

type FollowsTransResult struct {
	Follow Follow `json:"follow"`
}

func (store *SQLStore) FollowTx(ctx context.Context, arg FollowsTransParam) (FollowsTransResult, error) {
	var result FollowsTransResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Follow, err = q.FollowUser(ctx, arg.FollowUserParams)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
