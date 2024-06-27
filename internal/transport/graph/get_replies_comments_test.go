package graph

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetRepliesComments(t *testing.T) {
	var (
		limit         = 2
		offset        = 0
		refPostId     = 1
		refCommmentId = 1
	)
	testCases := []struct {
		name     string
		comment  *model.Comment
		expResp  []*model.Comment
		behavior func(td *testDeps, comment *model.Comment, expResp []*model.Comment) error
	}{
		{
			name: "status OK",
			comment: &model.Comment{
				ID: refCommmentId,
			},
			expResp: []*model.Comment{
				{
					ID:      1,
					Author:  "comment author 1",
					Content: "comment content 1",
					ReplyTo: &refCommmentId,
					Post:    refPostId,
				},
				{
					ID:      2,
					Author:  "comment author 2",
					Content: "comment content 2",
					ReplyTo: &refCommmentId,
					Post:    refPostId,
				},
			},
			behavior: func(td *testDeps, comment *model.Comment, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetRepliesComments(gomock.Any(), comment, gomock.Any(), gomock.Any()).
					Return(expResp, nil)
				return nil
			},
		},
		{
			name: "eror not nill",
			behavior: func(td *testDeps, comment *model.Comment, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetRepliesComments(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, models.ErrInternalServerResolver)
				return models.ErrInternalServerResolver
			},
		},
		{
			name: "eror not found",
			behavior: func(td *testDeps, comment *model.Comment, expResp []*model.Comment) error {
				td.service.EXPECT().
					GetRepliesComments(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
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

			expErr := tc.behavior(td, tc.comment, tc.expResp)

			resp, err := srv.Comment().Replies(context.Background(), tc.comment, &limit, &offset)

			require.Equal(t, expErr, err)
			require.Equal(t, tc.expResp, resp)
		})

	}
}
