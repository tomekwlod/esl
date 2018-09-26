package esl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Env struct {
	*Client
}

func (e *Env) SendLog(log *Log) (jr *JSONResp, err error) {

	// maybe there is a way to avoid doing this??
	if e.Project == "" || e.Url == "" {
		return nil, errors.New("Cannot send a log without a proper configuration")
	}

	// override the log.Poject
	log.Project = e.Project

	logBytes, err := json.Marshal(log)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(logBytes)

	req, err := http.NewRequest("POST", e.Client.Url+":"+strconv.Itoa(e.Client.Port)+"/logger/doc", body)
	if err != nil {
		return nil, err
	}

	auth := basicAuthToken(e.Client.User, e.Client.Password)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// contentString := string(content)
	// contentType := resp.Header.Get("Content-Type")

	if resp.StatusCode >= 300 {
		return nil, errors.New("Error code received [" + strconv.Itoa(resp.StatusCode) + "]")
	}
	// fmt.Println(jr)
	json.Unmarshal(content, &jr)
	if jr.ID == "" {
		return jr, errors.New("Data error")
	}
	// fmt.Println(jr)
	return jr, err
}

func basicAuthToken(user string, pass string) string {
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
}
