package config

import (
	"errors"
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
