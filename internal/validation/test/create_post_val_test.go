package test

import (
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/stretchr/testify/require"
)

func TestCreatePostValidate(t *testing.T) {
	testCases := []struct {
		name   string
		post   model.CreatePostReq
		expErr bool
	}{
		{
			name: "empty fields",
			post: model.CreatePostReq{
				Name:    "",
				Author:  "",
				Content: "",
			},
			expErr: true,
		},
		{
			name: "to long name",
			post: model.CreatePostReq{
				Name:    string(make([]rune, validation.MaxPostNameLength+1)),
				Author:  "test author",
				Content: "test content",
			},
			expErr: true,
		},
		{
			name: "to long author",
			post: model.CreatePostReq{
				Name:    "test name",
				Author:  string(make([]rune, validation.MaxAuthorLength+1)),
				Content: "test content",
			},
			expErr: true,
		},
		{
			name: "ok",
			post: model.CreatePostReq{
				Name:    "test ok",
				Author:  "test ok",
				Content: "test ok",
			},
			expErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := validation.CreatePostValidate(tc.post)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
