package keyauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var sellerkey = os.Getenv("KAUTH_SELLERKEY")

const UNIQUE_ID_USERVAR = "uuid1"
const UNIQUE_ID4_USERVAR = "uuid4"

/*
// url query parameters. May be encoded with:
encoder := qs.NewEncoder()
values, _ := encoder.Values(query)
fmt.Println(values.Encode()) // (unescaped) output: "sellerkey=abcd12345&type=testing123"
*/
type KQuery struct {
	// api auth values
	SellerKey string `qs:"sellerkey,omitempty"`
	AppName   string `qs:"name,omitempty"`
	OwnerId   string `qs:"ownerid,omitempty"`
	SessionId string `qs:"sessionid,omitempty"`

	// various request value types for various request types
	// i.e. https://keyauthdocs.apidog.io/api/features/check-session
	Type     string `qs:"type"`
	UserName string `qs:"user,omitempty"`
	VarName  string `qs:"var,omitempty"`
	Data     string `qs:"data,omitempty"`
}

func send_keyauth_request(request_type, user, varname, data string) (map[string]interface{}, error) {
	jsonRes := make(map[string]interface{})
	url := fmt.Sprintf("https://keyauth.win/api/seller/?sellerkey=%s&type=%s&user=%s&var=%s&data=%s", sellerkey, request_type, user, varname, data)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		jsonRes["message"] = "failed to make request"
		jsonRes["success"] = false
		return jsonRes, err
	}
	req.Header.Add("User-Agent", "Aegis5/1.0.0 (https://aegisfive.xyz)")

	res, err := client.Do(req)
	if err != nil {
		jsonRes["message"] = "failed to perform request"
		jsonRes["success"] = false
		return jsonRes, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		jsonRes["message"] = "failed to read response bytes"
		jsonRes["success"] = false
		return jsonRes, err
	}
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		jsonRes["message"] = "failed to unmarshal json data"
		jsonRes["success"] = false
		return jsonRes, err
	}

	return jsonRes, nil
}

func CheckUserExists(username string) bool {
	result, err := send_keyauth_request("verifyuser", username, "", "")
	if err != nil {
		return false
	}
	return result["success"].(bool)
}

func SetVar(user, varname, data string) (map[string]interface{}, error) {
	jsonRes := make(map[string]interface{})
	url := fmt.Sprintf("https://keyauth.win/api/seller/?sellerkey=%s&type=setvar&user=%s&var=%s&data=%s", sellerkey, user, varname, data)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		jsonRes["message"] = "failed to make requst"
		jsonRes["success"] = false
		return jsonRes, err
	}
	req.Header.Add("User-Agent", "Aegis5/1.0.0 (https://aegisfive.xyz)")

	res, err := client.Do(req)
	if err != nil {
		jsonRes["message"] = "failed to perform requst"
		jsonRes["success"] = false
		return jsonRes, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		jsonRes["message"] = "failed to read response bytes"
		jsonRes["success"] = false
		return jsonRes, err
	}
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		jsonRes["message"] = "failed to unmarshal json data"
		jsonRes["success"] = false
		return jsonRes, err
	}

	return jsonRes, nil
}

func GetVar(user, varname string) (map[string]interface{}, error) {
	jsonRes := make(map[string]interface{})
	url := fmt.Sprintf("https://keyauth.win/api/seller/?sellerkey=%s&type=getvar&user=%s&var=%s", sellerkey, user, varname)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		jsonRes["message"] = "failed to make requst"
		jsonRes["success"] = false
		return jsonRes, err
	}
	req.Header.Add("User-Agent", "Aegis5/1.0.0 (https://aegisfive.xyz)")

	res, err := client.Do(req)
	if err != nil {
		jsonRes["message"] = "failed to perform requst"
		jsonRes["success"] = false
		return jsonRes, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		jsonRes["message"] = "failed to read response bytes"
		jsonRes["success"] = false
		return jsonRes, err
	}
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		jsonRes["message"] = "failed to unmarshal json data"
		jsonRes["success"] = false
		return jsonRes, err
	}

	return jsonRes, nil
}
