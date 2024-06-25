package test

import (
	"testing"

	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/stretchr/testify/require"
)

func ptr(i int) *int {
	return &i
}
func TestPagingValidate(t *testing.T) {
	type args struct {
		limit  *int
		offset *int
	}
	testCases := []struct {
		name      string
		args      args
		expLimit  int
		expOffset int
	}{
		{
			name:      "Both nil",
			args:      args{nil, nil},
			expLimit:  0,
			expOffset: 0,
		},
		{
			name: "ok",
			args: args{
				limit:  ptr(2),
				offset: ptr(2)},
			expLimit:  2,
			expOffset: 4,
		},
		{
			name: "limit <0",
			args: args{
				limit:  ptr(-1),
				offset: ptr(2)},
			expLimit:  0,
			expOffset: 0,
		},
		{
			name: "offset <0",
			args: args{
				limit:  ptr(2),
				offset: ptr(-1)},
			expLimit:  2,
			expOffset: 0,
		},
		{
			name: "offset <0 limit <0",
			args: args{
				limit:  ptr(-1),
				offset: ptr(-1)},
			expLimit:  0,
			expOffset: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			resLimit, resOffset := validation.PagingValidate(tc.args.limit, tc.args.offset)
			require.Equal(t, tc.expLimit, resLimit)
			require.Equal(t, tc.expOffset, resOffset)
		})
	}
}
