package test

import (
	"context"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	mockRepo "github.com/AntonovIv/post_graphQlservice/internal/repository/postgre/mock"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testDeps struct {
	repo *mockRepo.MockpostRepo
	obs  *mockRepo.MockObserverPool
}

func newTestDeps(t *testing.T) *testDeps {
	t.Helper()

	ctl := gomock.NewController(t)
	t.Cleanup(func() {
		ctl.Finish()
	})

	return &testDeps{
		repo: mockRepo.NewMockpostRepo(ctl),
		obs:  mockRepo.NewMockObserverPool(ctl),
	}
}

func (td *testDeps) newService() *posts.PostService {
	return posts.New(td.repo, td.obs)
}

func TestCreateComment(t *testing.T) {
	testCases := []struct {
		name     string
		input    model.CreateCommentReq
		behavior func(td *testDeps, input model.CreateCommentReq) error
	}{
		{
			name: "test case: OK",
			input: model.CreateCommentReq{
				Author:  "test",
				Content: "test",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, input model.CreateCommentReq) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), input.Post).
					Return(model.Post{ID: 1, CommentsAllowed: true}, nil)

				td.repo.EXPECT().
					CreateComment(gomock.Any(), input).
					Return(model.Comment{Author: input.Author,
						Content: input.Content,
						Post:    input.Post,
						ReplyTo: input.ReplyTo}, nil)

				return nil
			},
		},
		{
			name: "test case: getpost in transaction failed",
			input: model.CreateCommentReq{
				Author:  "test",
				Content: "test",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, input model.CreateCommentReq) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), input.Post).
					Return(model.Post{}, assert.AnError)
				return assert.AnError
			},
		},
		{
			name: "test case: comment not alowed",
			input: model.CreateCommentReq{
				Author:  "test",
				Content: "test",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, input model.CreateCommentReq) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), input.Post).
					Return(model.Post{ID: 1, CommentsAllowed: false}, nil)
				return models.ErrBadPostId
			},
		},
		{
			name: "test case: create comment in transaction failed",
			input: model.CreateCommentReq{
				Author:  "test",
				Content: "test",
				Post:    1,
				ReplyTo: nil,
			},
			behavior: func(td *testDeps, input model.CreateCommentReq) error {
				td.repo.EXPECT().
					GetPostByID(gomock.Any(), input.Post).
					Return(model.Post{ID: 1, CommentsAllowed: true}, nil)

				td.repo.EXPECT().
					CreateComment(gomock.Any(), input).
					Return(model.Comment{}, assert.AnError)

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
				tc.input,
			)

			service := td.newService()
			_, err := service.CreateComment(context.Background(), tc.input)

			require.ErrorIs(t, err, expErr)
		})
	}
}
