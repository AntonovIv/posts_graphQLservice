package test

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllPost(t *testing.T) {
	testCases := []struct {
		name     string
		allPost  []model.PostListEl
		behavior func(td *testDeps, allPost []model.PostListEl) error
	}{
		{
			name: "test case: OK",
			allPost: []model.PostListEl{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			behavior: func(td *testDeps, allPost []model.PostListEl) error {
				td.repo.EXPECT().
					GetAllPosts(gomock.Any()).
					Return(allPost, nil)
				return nil
			},
		},
		{
			name:    "test case: get all posts failed",
			allPost: []model.PostListEl{},
			behavior: func(td *testDeps, allPost []model.PostListEl) error {
				td.repo.EXPECT().
					GetAllPosts(gomock.Any()).
					Return(allPost, assert.AnError)
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
				tc.allPost,
			)

			service := td.newService()
			postsResp, err := service.GetAllPosts(context.Background())

			if postsResp != nil {
				require.Equal(t, tc.allPost[0], *postsResp[0])
			}
			require.Equal(t, len(tc.allPost), len(postsResp))
			require.ErrorIs(t, err, expErr)
		})
	}
}
