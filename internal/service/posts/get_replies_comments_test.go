package posts

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRepliesComments(t *testing.T) {
	var refCommentId int = 1
	testCases := []struct {
		name         string
		comment      model.Comment
		commentsResp []model.Comment
		behavior     func(td *testDeps, comment model.Comment, commentsResp []model.Comment) error
	}{
		{
			name:    "test case: OK",
			comment: model.Comment{ID: refCommentId},
			commentsResp: []model.Comment{
				{
					ID:      1,
					Author:  "author 1",
					Content: "content 1",
					Post:    1,
					ReplyTo: &refCommentId,
				},
				{
					ID:      2,
					Author:  "author 2",
					Content: "content 2",
					Post:    1,
					ReplyTo: &refCommentId,
				},
			},
			behavior: func(td *testDeps, comment model.Comment, commentsResp []model.Comment) error {
				td.repo.EXPECT().
					GetRepliesComments(gomock.Any(), &comment, gomock.Any(), gomock.Any()).
					Return(commentsResp, nil)
				return nil
			},
		},
		{
			name:         "test case: repo func err",
			comment:      model.Comment{ID: 1},
			commentsResp: []model.Comment{},
			behavior: func(td *testDeps, comment model.Comment, commentsResp []model.Comment) error {
				td.repo.EXPECT().
					GetRepliesComments(gomock.Any(), &comment, gomock.Any(), gomock.Any()).
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
				tc.comment,
				tc.commentsResp,
			)

			service := td.newService()
			resp, err := service.GetRepliesComments(context.Background(), &tc.comment, 0, 0)

			if resp != nil {
				require.Equal(t, len(tc.commentsResp), len(resp))
				for i := range tc.commentsResp {
					require.Equal(t, tc.commentsResp[i], *resp[i])
				}
			}
			require.ErrorIs(t, err, expErr)

		})
	}
}
