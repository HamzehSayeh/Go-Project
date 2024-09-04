package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://69ez1r.api.infobip.com/sms/2/text/advanced"
	method := "POST"

	payload := strings.NewReader(`{"messages":[{"destinations":[{"to":"970597527861"}],"from":"447491163443","text":"Congratulations on sending your first message.\nGo ahead and check the delivery report in the next step."}]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "App 51bd59a15eaa6fa6068811fff871dc98-d7913772-54ee-4de1-af79-524f8c11edef")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
