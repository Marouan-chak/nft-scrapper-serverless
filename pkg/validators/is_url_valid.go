package validators

import "net/url"

func isURLValid(e string) bool {
	_, err := url.ParseRequestURI(e)
	if err != nil {
		return false
	}
	return true
}
