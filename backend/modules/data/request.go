package data

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/kaan-devoteam/firestore-security-demo/log"
)

func RequestPostWithToken(url, token string, body io.Reader) ([]byte, error) {
	// Makes request to firestore RestAPI url using Article model body (unmarshalled)
	// and the token comes from frontend
	req, errRequest := http.NewRequest(http.MethodPost, url, body)
	if errRequest != nil {
		log.Error(errRequest.Error())
		return nil, errRequest
	}
	req.Header.Add("Authorization", token)
	res, errResponse := http.DefaultClient.Do(req)
	if errResponse != nil {
		log.Error(errRequest.Error())
		return nil, errResponse
	}
	resBody, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Error(errBody.Error())
	}
	log.Info(fmt.Sprintf("Response is parsable: %s", string(resBody)))
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Oops. Firestore Api returned %v", res.StatusCode))
	}
	return resBody, errResponse
}

func RequestGetWithToken(url, token string) ([]byte, error) {
	req, errRequest := http.NewRequest(http.MethodGet, url, nil)
	if errRequest != nil {
		log.Error(errRequest.Error())
		return nil, errRequest
	}
	req.Header.Add("Authorization", token)
	res, errResponse := http.DefaultClient.Do(req)
	if errResponse != nil {
		log.Error(errRequest.Error())
		return nil, errResponse
	}
	resBody, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Error(errBody.Error())
	}
	log.Info(fmt.Sprintf("Response is parsable: %s", string(resBody)))
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Oops. Firestore Api returned %v", res.StatusCode))
	}
	return resBody, errResponse
}
