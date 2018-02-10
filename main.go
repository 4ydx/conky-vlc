package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type VLCResponse struct {
	Information Information
}

type Information struct {
	Category Category
}

type Category struct {
	Meta Meta
}

type Meta struct {
	Title      string
	NowPlaying string `json:"now_playing"`
	Genre      string
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/requests/status.json", nil)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	req.SetBasicAuth("", os.Getenv("VLC_PASSWORD"))

	// make request.  on failure display nothing when appropriate.
	resp, err := client.Do(req)
	if err != nil {
		m, e := regexp.MatchString("connection refused", err.Error())
		if e != nil {
			fmt.Fprint(os.Stdout, err.Error())
			return
		}
		if m {
			fmt.Fprint(os.Stdout, "--")
			return
		}
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	defer resp.Body.Close()

	// read response, fill the VLCResponse object, and return text to conky
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	vlc := &VLCResponse{}
	err = json.Unmarshal(body, vlc)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	tmpl, err := template.New("test").Parse("{{.Information.Category.Meta.NowPlaying}}")
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	err = tmpl.Execute(os.Stdout, vlc)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
	}
}
