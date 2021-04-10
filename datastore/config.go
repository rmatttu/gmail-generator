package config

import (
	"errors"
	"runtime"
	"time"
)

type Browser struct {
	OpenBrowser       bool `yaml:"openBrowser"`
	CustomBrowserPath bool `yaml:"customBrowserPath"`
	Path              string
	Option            string
}

type ReplacementData struct {
	ReplaceTarget string `yaml:"replaceTarget"`
	Method        string
	Option        string
}

type Template struct {
	Name        string
	Endpoint    string
	Replacement []ReplacementData
	TO          string
	CC          string
	BCC         string
	Subject     string
	Body        string
}

type Config struct {
	Browser  Browser
	Template []Template
}

func (u *ReplacementData) ReplaceNew() (*string, error) {
	if u.Method != "DATETIME" {
		return nil, errors.New("not datetime error")
	}
	t := time.Now().Format(u.Option)

	// TODO: go-sed sample
	// engine, err := sed.New(strings.NewReader(u.SED))
	// if err != nil {
	// 	return nil, err
	// }
	// output, err := engine.RunString(t)
	// if err != nil {
	// 	return nil, err
	// }
	// // Remove line feed
	// if output[len(output)-1] == '\n' {
	// 	output = output[:len(output)-1]
	// }

	return &t, nil
}

func Default() Config {
	path := ""
	switch runtime.GOOS {
	case "linux":
		path = "/usr/bin/google-chrome"
	case "windows":
		path = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	case "darwin":
		path = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}

	browser := Browser{
		OpenBrowser:       true,
		CustomBrowserPath: false,
		Path:              path,
		Option:            "",
	}
	replaceMents := []ReplacementData{}
	item := ReplacementData{
		ReplaceTarget: "___DATETIME___",
		Method:        "DATETIME",
		Option:        "2006-01-02",
	}
	replaceMents = append(replaceMents, item)
	item = ReplacementData{
		ReplaceTarget: "___DATETIME_JP___",
		Method:        "DATETIME",
		Option:        "2006年1月2日",
	}
	replaceMents = append(replaceMents, item)
	templates := []Template{}
	mainTemplate := Template{
		Name:        "main",
		Endpoint:    "https://mail.google.com/mail/u/2",
		Replacement: replaceMents,
		TO:          "to@example.com, to2@example.com, mailing-list@example.com",
		CC:          "cc@example.com",
		BCC:         "",
		Subject:     "業務日報(〇〇) ___DATETIME___",
		Body: `関係各位
以下、本日、___DATETIME_JP___の日報です。
	* ああああ
	* いいいい

以上

-- 
署名`,
	}
	templates = append(templates, mainTemplate)
	return Config{Browser: browser, Template: templates}
}
