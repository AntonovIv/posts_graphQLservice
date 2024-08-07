package posts

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	testCases := []struct {
		name        string
		post        model.CreatePostReq
		postSrvResp *model.PostListEl
		behavior    func(td *testDeps, post model.CreatePostReq) error
	}{
		{
			name: "test case: OK",
			post: model.CreatePostReq{
				Name:            "test name",
				Content:         "test content",
				Author:          "test author",
				CommentsAllowed: true,
			},
			postSrvResp: &model.PostListEl{ID: 1,
				Name:    "test name",
				Author:  "test author",
				Content: "test content"},
			behavior: func(td *testDeps, post model.CreatePostReq) error {
				td.repo.EXPECT().
					CreatePost(gomock.Any(), post).
					Return(model.PostListEl{
						ID:      1,
						Name:    post.Name,
						Author:  post.Author,
						Content: post.Content,
					}, nil)
				return nil
			},
		},
		{
			name:        "test case: create post failed",
			post:        model.CreatePostReq{},
			postSrvResp: nil,
			behavior: func(td *testDeps, post model.CreatePostReq) error {
				td.repo.EXPECT().
					CreatePost(gomock.Any(), post).
					Return(model.PostListEl{}, assert.AnError)
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
			)

			service := td.newService()
			resp, err := service.CreatePost(context.Background(), tc.post)

			require.Equal(t, tc.postSrvResp, resp)
			require.ErrorIs(t, err, expErr)
		})
	}
}
