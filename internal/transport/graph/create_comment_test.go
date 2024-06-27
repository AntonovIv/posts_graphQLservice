package graph

import (
	"context"
	"fmt"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestCreateComment(t *testing.T) {

	testCases := []struct {
		name       string
		commentReq model.CreateCommentReq
		expResp    *model.Comment
		behavior   func(td *testDeps, comment model.CreateCommentReq, expResp *model.Comment) error
	}{
		{
			name: "status OK",
			commentReq: model.CreateCommentReq{
				Content: "test content",
				Author:  "test author",
				Post:    1,
				ReplyTo: nil,
			},
			expResp: &model.Comment{
				ID:      1,
				Author:  "test author",
				Content: "test content",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, comment model.CreateCommentReq, expResp *model.Comment) error {
				retResp := *expResp
				td.service.EXPECT().
					CreateComment(gomock.Any(), comment).
					Return(&retResp, nil)
				td.service.EXPECT().
					NotifyObservers(gomock.Any(), gomock.Any()).
					Return(nil)
				return nil

			},
		},
		{
			name: "internal server error",
			commentReq: model.CreateCommentReq{
				Content: "test content",
				Author:  "test author",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, comment model.CreateCommentReq, expResp *model.Comment) error {
				td.service.EXPECT().
					CreateComment(gomock.Any(), comment).
					Return(nil, models.ErrInternalServerResolver)
				return models.ErrInternalServerResolver
			},
		},
		{
			name: "comments not allowed",
			commentReq: model.CreateCommentReq{
				Content: "test content",
				Author:  "test author",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, comment model.CreateCommentReq, expResp *model.Comment) error {
				td.service.EXPECT().
					CreateComment(gomock.Any(), comment).
					Return(nil, models.ErrBadPostId)
				return models.ErrCommentsNotAllowedResolver
			},
		},
		{
			name: "validation error",
			commentReq: model.CreateCommentReq{
				Content: "",
				Author:  "test author",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, comment model.CreateCommentReq, expResp *model.Comment) error {
				return &gqlerror.Error{
					Message: "bad request: " + fmt.Errorf("forbidden empty field").Error(),
				}
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			srv := td.newSrv()

			expErr := tc.behavior(td, tc.commentReq, tc.expResp)

			resp, err := srv.Mutation().CreateComment(context.Background(), tc.commentReq)
			require.Equal(t, expErr, err)
			require.Equal(t, tc.expResp, resp)
		})

	}
}
