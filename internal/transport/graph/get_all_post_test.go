package graph

import (
	"context"
	"log/slog"
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	mockSrv "github.com/AntonovIv/post_graphQlservice/internal/service/posts/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type testDeps struct {
	service *mockSrv.MockpostsSrv
}

func newTestDeps(t *testing.T) *testDeps {
	t.Helper()

	ctr := gomock.NewController(t)
	t.Cleanup(func() {
		ctr.Finish()
	})

	return &testDeps{
		service: mockSrv.NewMockpostsSrv(ctr),
	}
}

func (td *testDeps) newSrv() *Resolver {
	return New(td.service, slog.New(slog.Default().Handler()))
}

func TestGetAllPosts(t *testing.T) {

	testCases := []struct {
		name     string
		resp     []*model.PostListEl
		behavior func(td *testDeps, reps []*model.PostListEl) error
	}{
		{
			name: "status OK",
			resp: []*model.PostListEl{
				{
					ID:      1,
					Content: "test",
				},
				{
					ID:      2,
					Content: "test",
				},
			},
			behavior: func(td *testDeps, resp []*model.PostListEl) error {
				td.service.EXPECT().
					GetAllPosts(gomock.Any()).
					Return(resp, nil)
				return nil
			},
		},
		{
			name: "eror not nill",
			resp: nil,
			behavior: func(td *testDeps, resp []*model.PostListEl) error {
				td.service.EXPECT().
					GetAllPosts(gomock.Any()).
					Return(nil, &gqlerror.Error{
						Message: "internal server error",
					})
				return &gqlerror.Error{
					Message: "internal server error",
				}
			},
		},
		{
			name: "eror not found",
			resp: nil,
			behavior: func(td *testDeps, resp []*model.PostListEl) error {
				td.service.EXPECT().
					GetAllPosts(gomock.Any()).
					Return(nil, models.ErrNotFound)
				return &gqlerror.Error{
					Message: "posts not found",
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

			expErr := tc.behavior(td, tc.resp)

			resp, err := srv.Query().GetAllPosts(context.Background())

			require.Equal(t, err, expErr)
			require.Equal(t, tc.resp, resp)
		})

	}
}
