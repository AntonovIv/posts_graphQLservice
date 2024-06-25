package test

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCommentsForPost(t *testing.T) {
	testCases := []struct {
		name     string
		post     model.Post
		comments []model.Comment
		behavior func(td *testDeps, post model.Post, comments []model.Comment) error
	}{
		{
			name: "test case: OK",
			post: model.Post{ID: 1},
			comments: []model.Comment{
				{
					ID:   1,
					Post: 1,
				},
				{
					ID:   2,
					Post: 1,
				},
			},
			behavior: func(td *testDeps, post model.Post, comments []model.Comment) error {
				td.repo.EXPECT().
					GetCommentsForPost(gomock.Any(), &post, gomock.Any(), gomock.Any()).
					Return(comments, nil)
				return nil
			},
		},
		{
			name:     "test case: get all posts failed",
			post:     model.Post{ID: 1},
			comments: []model.Comment{},
			behavior: func(td *testDeps, post model.Post, comments []model.Comment) error {
				td.repo.EXPECT().
					GetCommentsForPost(gomock.Any(), &post, gomock.Any(), gomock.Any()).
					Return(comments, assert.AnError)
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
				tc.post,
				tc.comments,
			)

			service := td.newService()
			postsResp, err := service.GetCommentsForPost(context.Background(), &tc.post, 0, 0)

			if postsResp != nil {
				require.Equal(t, len(tc.comments), len(postsResp))
				for i := range tc.comments {
					require.Equal(t, tc.comments[i], *postsResp[i])
					require.Equal(t, tc.post.ID, postsResp[i].Post)
				}
			}
			require.Equal(t, len(tc.comments), len(postsResp))
			require.ErrorIs(t, err, expErr)
		})
	}
}
