package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type GatewayActivate struct {
	GatewayConf GatewayConf `yaml:"Gateway"`
}

type GatewayConf struct {
	AccessId string `yaml:"AccessId"`
	AccessKey string `yaml:"AccessKey"`
}

func TestReadFile(t *testing.T) {
	path := "/usr/local/var/tuya/gateway.conf"
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return
	}
	defer f.Close()
	gateway := &GatewayActivate{}
	yaml.NewDecoder(f).Decode(gateway)
	fmt.Println(gateway)
	fmt.Println(gateway.GatewayConf.AccessId)
	fmt.Println(f.Name())
}

/**
 * 中国预发：https://openapi-cn.wgine.com
 * 中国线上：https://openapi.tuyacn.com
 */
func TestGetToken(tt *testing.T){
	clientId := "whynhm4ma0ptficy4aiz"
	key := "ab079313a49044b9bfaf769b756858f2"
	url := "https://openapi-cn.wgine.com/v1.0/token?grant_type=1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := strconv.FormatInt(time.Now().Unix()*1000, 10)
	s := strings.ToUpper(SHA265Sign([]byte(clientId+t), []byte(key)))
	req.Header.Add("client_id", clientId)
	req.Header.Add("sign", s)
	req.Header.Add("t", t)
	req.Header.Add("sign_method", "HMAC-SHA256")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type PostBody struct {
	Mac string `json:"mac"`
}

func TestGatewayActivate(tt *testing.T) {
	clientId := "whynhm4ma0ptficy4aiz"
	key := "ab079313a49044b9bfaf769b756858f2"
	url := "https://openapi-cn.wgine.com/v1.0/pass/devices/paring/auto"
	mac := "888888888"
	client := &http.Client{}
	msg := PostBody{
		Mac: mac,
	}
	msgJ, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", url, bytes.NewReader(msgJ))
	if err != nil {
		fmt.Println(err)
		return
	}
	token := "8d4abdec25e90c0b359261e2d78fcddb"
	t := strconv.FormatInt(time.Now().Unix()*1000, 10)
	s := strings.ToUpper(SHA265Sign([]byte(clientId+token+t), []byte(key)))
	req.Header.Add("client_id", clientId)
	req.Header.Add("access_token", token)
	req.Header.Add("sign", s)
	req.Header.Add("t", t)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("Content0-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body)+"s")
	return
}

/**
 * hujun@tuya.com
 * highway请求header参数加密
 */
func SHA265Sign(msg []byte, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(msg)
	return hex.EncodeToString(h.Sum(nil))
}

