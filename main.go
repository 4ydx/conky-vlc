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
	Filename   string
	NowPlaying string `json:"now_playing"`
	Genre      string
}

// Conky will periodically run this program, trying to access vlc's http api in order to find out
// what vlc is currently playing.  Make sure to have vlc playing something with the http api enabled.
// EG:
// cvlc ~/Music/*.mp3 --extraintf http --http-password <my-password>
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/requests/status.json", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
	req.SetBasicAuth("", os.Getenv("VLC_PASSWORD"))

	// make request.  on failure display nothing when appropriate.
	resp, err := client.Do(req)
	if err != nil {
		m, e := regexp.MatchString("connection refused", err.Error())
		if e != nil {
			fmt.Print(e)
			return
		}
		if m {
			fmt.Print("--")
			return
		}
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()

	// read response, fill the VLCResponse object, and return text to conky
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	vlc := &VLCResponse{}
	err = json.Unmarshal(body, vlc)
	if err != nil {
		fmt.Print(err)
		return
	}
	var tmpl *template.Template
	if vlc.Information.Category.Meta.NowPlaying != "" {
		tmpl, err = template.New("test").Parse("{{.Information.Category.Meta.NowPlaying}}")
		if err != nil {
			fmt.Print(err)
			return
		}
	} else {
		tmpl, err = template.New("test").Parse("{{.Information.Category.Meta.Title}}")
		if err != nil {
			fmt.Print(err)
			return
		}
	}
	err = tmpl.Execute(os.Stdout, vlc)
	if err != nil {
		fmt.Print(err)
	}
}
