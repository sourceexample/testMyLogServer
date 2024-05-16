package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const G_APPID = "testapp1"
const G_Endpoint = "https://gatlinglogserver.onrender.com/"

func logInfo(log string) error {
	url := G_Endpoint + "info/" + G_APPID

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(log)))
	if err != nil {
		return err
	}
	req.Header.Set("X_API_KEY", "aaXAPIKEYaa")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("success")
	} else {
		fmt.Println("failed: ", resp)
		return errors.New("failed")
	}

	return nil

}
func logError(log string) error {
	url := G_Endpoint + "error/" + G_APPID

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(log)))
	if err != nil {
		return err
	}
	req.Header.Set("X_API_KEY", "aaXAPIKEYaa")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("success")
	} else {
		fmt.Println("failed: ", resp)
		return errors.New("failed")
	}

	return nil

}

func main() {
	logError("error client example running 2")
}
