package graph

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCommentsForPost(t *testing.T) {
	var (
		limit     = 2
		offset    = 0
		refPostId = 1
	)
	testCases := []struct {
		name     string
		post     *model.Post
		expResp  []*model.Comment
		behavior func(td *testDeps, post *model.Post, expResp []*model.Comment) error
	}{
		{
			name: "status OK",
			post: &model.Post{
				ID: refPostId,
			},
			expResp: []*model.Comment{
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
			behavior: func(td *testDeps, post *model.Post, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetCommentsForPost(gomock.Any(), post, gomock.Any(), gomock.Any()).
					Return(expResp, nil)
				return nil
			},
		},
		{
			name: "eror not nill",
			behavior: func(td *testDeps, post *model.Post, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetCommentsForPost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, models.ErrInternalServerResolver)
				return models.ErrInternalServerResolver
			},
		},
		{
			name: "eror not found",
			behavior: func(td *testDeps, post *model.Post, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetCommentsForPost(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, models.ErrNotFound)
				return models.ErrNotFoundCommentsResolver
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			srv := td.newSrv()

			expErr := tc.behavior(td, tc.post, tc.expResp)

			resp, err := srv.Post().Comments(context.Background(), tc.post, &limit, &offset)

			require.Equal(t, expErr, err)
			require.Equal(t, tc.expResp, resp)
		})

	}
}
