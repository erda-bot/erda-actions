package archive

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/erda-project/erda/pkg/customhttp"
	"github.com/pkg/errors"
)

func RequestPost(url string, payload []byte, header http.Header) ([]byte, *http.Response, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	request, err := customhttp.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	for k, values := range header {
		for _, v := range values {
			request.Header.Add(k, v)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, response, nil
}

func RequestGet(url string, header http.Header) ([]byte, *http.Response, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	request, err := customhttp.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	request.Header = header

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, response, nil
}
