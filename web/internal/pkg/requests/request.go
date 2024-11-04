package requests

import (
	"fmt"
	"io"
	"net/http"
	"webapp/internal/pkg/cookies"
)

func DoRequestWithAuth(r *http.Request, method, url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie.Token))

	client := &http.Client{}
	return client.Do(request)
}
