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

func TestCreatePost(t *testing.T) {

	testCases := []struct {
		name     string
		postReq  model.CreatePostReq
		expResp  *model.PostListEl
		behavior func(td *testDeps, post model.CreatePostReq) error
	}{
		{
			name: "status OK",
			postReq: model.CreatePostReq{
				Name:            "test name",
				Content:         "test content",
				Author:          "test author",
				CommentsAllowed: true,
			},
			expResp: &model.PostListEl{
				ID:      1,
				Name:    "test name",
				Author:  "test author",
				Content: "test content",
			},
			behavior: func(td *testDeps, post model.CreatePostReq) error {
				td.service.EXPECT().
					CreatePost(gomock.Any(), post).
					Return(&model.PostListEl{ID: 1, Name: post.Name,
						Author:  post.Author,
						Content: post.Content},
						nil)
				return nil
			},
		},
		{
			name: "internal server error",
			postReq: model.CreatePostReq{
				Name:            "test name",
				Content:         "test content",
				Author:          "test author",
				CommentsAllowed: true,
			},
			behavior: func(td *testDeps, post model.CreatePostReq) error {
				td.service.EXPECT().
					CreatePost(gomock.Any(), post).
					Return(nil, models.ErrInternalServerResolver)
				return models.ErrInternalServerResolver
			},
		},
		{
			name: "validation error",
			postReq: model.CreatePostReq{
				Name:            "",
				Content:         "test content",
				Author:          "test author",
				CommentsAllowed: true,
			},
			behavior: func(td *testDeps, post model.CreatePostReq) error {
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

			expErr := tc.behavior(td, tc.postReq)

			resp, err := srv.Mutation().CreatePost(context.Background(), tc.postReq)

			require.Equal(t, expErr, err)
			require.Equal(t, tc.expResp, resp)
		})

	}
}
