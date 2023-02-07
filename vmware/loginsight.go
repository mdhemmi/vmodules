package vmodules

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func Login(username string, password string, provider string, vrli string) string {

	// curl -k -X POST https://loginsight.example.com:9543/api/v1/sessions \
	//     -d '{"username":"admin","password":"Secret!","provider":"Local"}'

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var requestbody = "{\"username\":\"" + username + "\",\"password\":\"" + password + "\",\"provider\":\"" + provider + "\"}"

	body := strings.NewReader(requestbody)
	req, err := http.NewRequest("POST", "https://"+vrli+":9543/api/v1/sessions", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)
	//spew.Dump(result)
	sessionId := fmt.Sprint(result["sessionId"])

	return sessionId

}

func Query(vrli string, sessionId string) {
	// GET /api/v1/aggregated-events/{+path}

	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	body := strings.NewReader("")
	req, err := http.NewRequest("GET", "https://"+vrli+":9543/api/v1/events", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+sessionId)

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)
	spew.Dump(result)
}
