package graph

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestGetPost(t *testing.T) {
	testCases := []struct {
		name     string
		reqId    int
		expResp  *model.Post
		behavior func(td *testDeps, expResp *model.Post, reqId int) error
	}{
		{
			name:  "status OK",
			reqId: 1,
			expResp: &model.Post{
				ID:              1,
				Name:            "test name",
				Author:          "test author",
				Content:         "test content",
				CommentsAllowed: true,
			},
			behavior: func(td *testDeps, expResp *model.Post, reqId int) error {
				td.service.EXPECT().
					GetPostByID(gomock.Any(), reqId).
					Return(expResp, nil)
				return nil
			},
		},
		{
			name:    "eror not nill",
			expResp: nil,
			reqId:   1,
			behavior: func(td *testDeps, expResp *model.Post, reqId int) error {
				td.service.EXPECT().
					GetPostByID(gomock.Any(), reqId).
					Return(nil, models.ErrInternalServerResolver)
				return models.ErrInternalServerResolver
			},
		},
		{
			name:    "eror not found",
			reqId:   1,
			expResp: nil,
			behavior: func(td *testDeps, expResp *model.Post, reqId int) error {
				td.service.EXPECT().
					GetPostByID(gomock.Any(), reqId).
					Return(nil, models.ErrNotFound)
				return models.ErrNotFoundPostResolver
			},
		},
		{
			name:    "invalid request",
			reqId:   -1,
			expResp: nil,
			behavior: func(td *testDeps, expResp *model.Post, reqId int) error {
				return &gqlerror.Error{
					Message: "bad request: " + "invalid id",
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

			expErr := tc.behavior(td, tc.expResp, tc.reqId)

			resp, err := srv.Query().GetPostByID(context.Background(), tc.reqId)

			require.Equal(t, expErr, err)
			require.Equal(t, tc.expResp, resp)
		})

	}
}
