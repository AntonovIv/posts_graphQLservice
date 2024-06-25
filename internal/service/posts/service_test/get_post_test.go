package test

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPost(t *testing.T) {
	testCases := []struct {
		name        string
		postid      int
		postSrvResp *model.Post
		behavior    func(td *testDeps, id int) error
	}{
		{
			name:   "test case: OK",
			postid: 1,
			postSrvResp: &model.Post{
				ID: 1,
			},
			behavior: func(td *testDeps, id int) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), id).
					Return(model.Post{
						ID: 1,
					}, nil)
				return nil
			},
		},
		{
			name:        "test case: get post by id  failed",
			postid:      1,
			postSrvResp: nil,
			behavior: func(td *testDeps, id int) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), id).
					Return(model.Post{}, assert.AnError)
				return assert.AnError
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			expErr := tc.behavior(
				td,
				tc.postid,
			)

			service := td.newService()
			post, err := service.GetPostByID(context.Background(), tc.postid)

			require.Equal(t, tc.postSrvResp, post)
			require.ErrorIs(t, err, expErr)
		})
	}
}
