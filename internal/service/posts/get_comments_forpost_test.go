package posts

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCommentsForPost(t *testing.T) {
	var refPostId int = 1
	testCases := []struct {
		name            string
		post            model.Post
		expCommentsResp []model.Comment
		behavior        func(td *testDeps, post model.Post, expCommentsResp []model.Comment) error
	}{
		{
			name: "test case: OK",
			post: model.Post{ID: refPostId},
			expCommentsResp: []model.Comment{
				{
					ID:      1,
					Author:  "comment author 1",
					Content: "comment content 1",
					ReplyTo: nil,
					Post:    refPostId,
				},
				{
					ID:      2,
					Author:  "comment author 2",
					Content: "comment content 2",
					ReplyTo: nil,
					Post:    refPostId,
				},
			},
			behavior: func(td *testDeps, post model.Post, commentsResp []model.Comment) error {
				td.repo.EXPECT().
					GetCommentsForPost(gomock.Any(), &post, gomock.Any(), gomock.Any()).
					Return(commentsResp, nil)
				return nil
			},
		},
		{
			name:            "test case: repo func get comments for post err",
			post:            model.Post{ID: refPostId},
			expCommentsResp: []model.Comment{},
			behavior: func(td *testDeps, post model.Post, commentsResp []model.Comment) error {
				td.repo.EXPECT().
					GetCommentsForPost(gomock.Any(), &post, gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
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
				tc.expCommentsResp,
			)

			service := td.newService()
			resp, err := service.GetCommentsForPost(context.Background(), &tc.post, 0, 0)

			if resp != nil {
				require.Equal(t, len(tc.expCommentsResp), len(resp))
				for i := range tc.expCommentsResp {
					require.Equal(t, tc.expCommentsResp[i], *resp[i])
				}
			}
			require.ErrorIs(t, err, expErr)
		})
	}
}
