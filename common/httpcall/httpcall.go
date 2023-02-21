package httpcall

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mdhemmi/vmodules/common/getfuncname"
	"github.com/mdhemmi/vmodules/common/vdebug"
)

func HTTPcall(target string, token string, call string, method string, content string, insecure bool, debug bool) string {
	authtoken := "Bearer " + token
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	url := "https://" + target + call
	req, err := http.NewRequest(method, url, nil)

	if content == "null" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		var jsonStr = []byte(content)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	}
	if debug {
		funcname := getfuncname.GetCurrentFuncName()
		vdebug.Debug("", "", "start")
		vdebug.Debug(funcname, "string", "print")
		fmt.Println(authtoken)
		fmt.Println(err)
		fmt.Println(req)
		vdebug.Debug("", "", "end")

	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Frame_options", "Deny")
	req.Header.Set("Authorization", authtoken)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonstring := string(body)
	return jsonstring
}
