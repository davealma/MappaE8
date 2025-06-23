package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HitDors(cookieVal string) (string, string) {
	
	req, err := http.NewRequest("Post", "https://makers-challenge.altscore.ai/v1/s1/e8/actions/door", nil)
	if err != nil {
		panic(err)
	}
	if cookieVal != "" {
		cookie := &http.Cookie{
			Name: "gryffindor",
			Value: cookieVal,
		}
		req.AddCookie(cookie)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("API-KEY", "ca7d669ffc1746079404cb4d5fb92a3d")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "gryffindor" {
			return cookie.Value, string(bodyResp)
		}
	}
	return "", string(bodyResp)
}

func PostSolution (message string) {
	body := []byte(`{
		"hidden_message": "`+message+`"
	}`)

	req, err := http.NewRequest("Post", "https://makers-challenge.altscore.ai/v1/s1/e8/solution", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("API-KEY", "ca7d669ffc1746079404cb4d5fb92a3d")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Solution Response: ", string(bodyResp))
}