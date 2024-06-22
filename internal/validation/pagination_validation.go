package validation

func PagingValidate(limit *int, offset *int) (int, int) {

	hOff := 0
	hlim := 0

	if offset != nil && *offset > 0 {
		hOff = *offset
	}
	if limit != nil && *limit > 0 {
		hlim = *limit
	}

	hOff = hlim * hOff

	return hlim, hOff
}
