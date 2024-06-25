package test

import (
	"testing"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/stretchr/testify/require"
)

func TestCreateCommentValidate(t *testing.T) {
	testCases := []struct {
		name    string
		comment model.CreateCommentReq
		expErr  bool
	}{
		{
			name: "empty fields",
			comment: model.CreateCommentReq{
				Author:  "",
				Content: "",
			},
			expErr: true,
		},
		{
			name: "to long content",
			comment: model.CreateCommentReq{
				Author:  "",
				Content: string(make([]rune, 2001)),
			},
			expErr: true,
		},
		{
			name: "to long author",
			comment: model.CreateCommentReq{
				Author:  string(make([]rune, 65)),
				Content: "",
			},
			expErr: true,
		},
		{
			name: "ok",
			comment: model.CreateCommentReq{
				Author:  "test ok",
				Content: "test ok",
			},
			expErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := validation.CreateCommentValidate(tc.comment)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
