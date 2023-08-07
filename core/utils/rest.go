package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func SimpleFetch(url string) ([]byte, error) {
	client := &http.Client{}
	return executeSimpleFetch(url, client)
}

func Post(url string, request interface{}, client *http.Client) ([]byte, error) {
	buf, err := CreateHttpRequestBody(request)
	if err != nil {
		return nil, err
	}

	log.Infof("Post: %s", url)
	res, err := client.Post(url, "application/json", &buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		err = fmt.Errorf("%s, body: %s", res.Status, body)
		log.Error(err)
		return body, err
	}

	return body, nil
}

func CreateHttpRequestBody(request interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		log.Fatal(err)
	}
	return buf, err
}

func executeSimpleFetch(url string, client *http.Client) ([]byte, error) {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
