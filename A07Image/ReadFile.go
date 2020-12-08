package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

func main() {

	url := "https://cetus-cn.wgine.com/v1.0/file/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	ff, _ := ioutil.ReadFile("A07Image/flow.jpg")
	fmt.Printf("%T", payload) // *bytes.Buffer
	writer := multipart.NewWriter(payload)
	//file, errFile1 := os.Open("A07Image/flow.jpg")
	//defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file","A07Image/flow.jpg")
	//_, errFile1 = io.Copy(part1, file)
	if errFile1 !=nil {
		fmt.Println(errFile1)
	}
	fmt.Printf("part1: %v", part1)
	//err := writer.Close()
	//if err != nil {
	//	fmt.Println(err)
	//}
	client := &http.Client {
	}

	m := make(map[string]string)
	m["file"] = string(ff)
	r, _ := json.Marshal(m)
	fmt.Printf("%T", r)
	//req, err := http.NewRequest(method, url, payload)
	req, err := http.NewRequest(method, url, strings.NewReader("file="+string(ff)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("req: %v", req)
	req.Header.Add("file_upload_token", "4be6f63341ca93a230d0f52ddb02c406")
	req.Header.Add("User-Agent", "PostmanRuntime/7.23.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	//req.Header.Add("Postman-Token", "16f969fe-c57f-4d43-b9ad-fee26f5899ae")
	req.Header.Add("Host", "cetus-cn.wgine.com")
	//req.Header.Add("Content-Type", "multipart/form-data; boundary=--------------------------021727496917787682338797")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Length", "14769")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Charset","utf-8")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=--------------------------021727496917787682338797")
	fmt.Printf("%v\n",writer.FormDataContentType() )
	res, err := client.Do(req)
	defer res.Body.Close()
	var reader io.ReadCloser
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return
		}
	} else {
		reader = res.Body
	}
	body, err := ioutil.ReadAll(reader)
	fmt.Printf("%s, %d\n", res.Status, res.StatusCode)
	fmt.Printf("%s\n", res.Body)
	fmt.Printf("%s\n", body)
	fmt.Println(string(body))
}