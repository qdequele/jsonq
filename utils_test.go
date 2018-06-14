package jsonQuerry

import (
	"testing"
)

var querry = `{users:{username},topics:{topics:{title,fancy_title}}}`

func TestKeepRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		TestKeepRequestSuccess(t)
	})

	t.Run("error", func(t *testing.T) {
		// testParseRawNumberError(t, "xyz", "xyz")
	})
}

func TestKeepRequestSuccess(t *testing.T) {
	t.Helper()

	request, _ := NewKeepRequest(querry)
	stay, cont := getKeys(request)
	// fmt.Println(stay)
	// fmt.Println(cont)
}
