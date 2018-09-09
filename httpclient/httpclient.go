// httpclient
package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"g"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"net/url"
)

func Client(method string, url string, data []byte) []byte {
	sinInfo := getsig()

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		log.Fatal("NewRequest failed,err:%v", err.Error())
	}

	var apitoken = models.ApiToken{
		Name: sinInfo.Name,
		Sig:  sinInfo.Sig,
	}
	tocken, _ := json.Marshal(apitoken)
	req.Header.Set("Apitoken", string(tocken))
	req.Header.Set("X-Forwarded-For", "127.0.0.1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("name", sinInfo.Name)
	req.Header.Set("sig", sinInfo.Sig)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read body failed,err:%v", err.Error())
		return nil
	}
	return body

}

func getsig() *models.Sigure {
	var login_url = fmt.Sprintf("http://%s/api/v1/user/login", g.Config().ApiAddr)
	resp, err := http.PostForm(login_url, url.Values{"name": {"admin"}, "password": {"Ebupt@2017cf"}})
	if err != nil {
		log.Fatal("get sigure failed,err:%v", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read body failed,err:%v", err.Error())
	}
	var s = &models.Sigure{}
	json.Unmarshal(body, s)
	return s
}
