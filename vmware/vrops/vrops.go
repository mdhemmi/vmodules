package vrops

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get_vrops_token(vrops string, vropsusername string, vropspassword string, authSource string, debug bool, insecure bool) string {
	var token string
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	url := "https://" + vrops + "/suite-api/api/auth/token/acquire"
	loginjson := "{\"username\": \"" + vropsusername + "\",\"password\": \"" + vropspassword + "\",\"authSource\": \"" + authSource + "\"}"
	var jsonStr = []byte(loginjson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	xmlstring := string(body)

	token = Between(xmlstring, "<ops:token>", "</ops:token>")
	if debug {
		funcname := GetCurrentFuncName()
		Debug("", "", "start")
		Debug(funcname, "string", "print")
		//Debug(xmlstring, "string", "print")
		Debug(token, "string", "print")
		Debug("", "", "end")
	}

	return token
}

// ==============================================================================================================================

func Get_vrops_resource(vropshost string, token string, vropsname string, insecure bool) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	// https://192.168.0.29/suite-api/api/resources?adapterKind=VMWARE&name=skyline&page=0&pageSize=1000&resourceKind=VirtualMachine&_no_links=true
	url := "https://" + vropshost + "/suite-api/api/resources?name=" + vropsname + "&resourceKind=VirtualMachine&adapterKind=VMWARE"
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	vropstoken := "vRealizeOpsToken " + token
	req.Header.Set("Authorization", vropstoken)
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
	//spew.Dump(jsonstring)
	resource := Between(jsonstring, "identifier=\"", "\"><ops:resourceKey>")
	//fmt.Println(resource)
	return resource
}

// ==============================================================================================================================

func Post_vrops_metric(vropshost string, vropsobjectid string, token string, endpoint string, subitem string, timestamp string, state string, insecure bool) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	content := "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?><ops:stat-contents xmlns:ops=\"http://webservice.vmware.com/vRealizeOpsMgr/1.0/\" xmlns:xs=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"><ops:stat-content statKey=\"Skyline|" + endpoint + "|" + subitem + "\"><ops:timestamps>" + timestamp + "</ops:timestamps><ops:values>" + state + "</ops:values></ops:stat-content></ops:stat-contents>"
	url := "https://" + vropshost + "/suite-api/api/resources/" + vropsobjectid + "/stats"
	var str = []byte(content)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
	vropstoken := "vRealizeOpsToken " + token
	req.Header.Set("Authorization", vropstoken)
	req.Header.Set("Content-Type", "application/xml")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//jsonstring := string(body)
	//fmt.Println(jsonstring)
}

// ==============================================================================================================================

func Send_vROPs(vrops string, vropsusername string, vropspassword string, authSource string, debug bool, insecure bool, item string, state string, vmname string) {

	check := Check_Service_Availability(vrops, 443)
	if check {
		vrops_token := Get_vrops_token(vrops, vropsusername, vropspassword, authSource, debug, insecure)
		vropsobjectid := Get_vrops_resource(vrops, vrops_token, vmname, false)
		timestamp := fmt.Sprint((time.Now().UTC().UnixNano() / int64(time.Millisecond)))
		fmt.Println(endpoint, subitem, timestamp, state)
		Post_vrops_metric(vrops, vropsobjectid, vrops_token, endpoint, item, timestamp, state, insecure)
	}
}
