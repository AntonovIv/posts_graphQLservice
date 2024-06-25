package posts

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
					ID:      1,
					Name:    "test name",
					Author:  "test author",
					Content: "test content",
				},
				{
					ID:      2,
					Name:    "test name 2",
					Author:  "test author 2",
					Content: "test content 2",
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
			name:    "test case: repo func get all post err",
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
			resp, err := service.GetAllPosts(context.Background())

			if resp != nil {
				require.Equal(t, len(tc.allPost), len(resp))
				for i := range tc.allPost {
					require.Equal(t, tc.allPost[i], *resp[i])
				}
			}
			require.ErrorIs(t, err, expErr)
		})
	}
}
