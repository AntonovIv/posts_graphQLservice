package posts

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPost(t *testing.T) {
	testCases := []struct {
		name        string
		postid      int
		postSrvResp *model.Post
		behavior    func(td *testDeps, id int, postSrvResp *model.Post) error
	}{
		{
			name:   "test case: OK",
			postid: 1,
			postSrvResp: &model.Post{
				ID:              1,
				Name:            "test name",
				Author:          "test author",
				Content:         "test content",
				CommentsAllowed: true,
			},
			behavior: func(td *testDeps, id int, postSrvResp *model.Post) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), id).
					Return(*postSrvResp, nil)
				return nil
			},
		},
		{
			name:        "test case: repo get post err",
			postid:      1,
			postSrvResp: nil,
			behavior: func(td *testDeps, id int, postSrvResp *model.Post) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), id).
					Return(model.Post{}, assert.AnError)
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
				tc.postid,
				tc.postSrvResp)

			service := td.newService()
			post, err := service.GetPostByID(context.Background(), tc.postid)

			require.Equal(t, tc.postSrvResp, post)
			require.ErrorIs(t, err, expErr)
		})
	}
}
