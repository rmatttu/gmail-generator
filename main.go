package main

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	config "gmail-generator/datastore"
)

// Args is commandline args
type Args struct {
	Template *int
	Dryrun   *bool
}

func openDefaultBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	}
	return errors.New("No such as GOOS")
}

func open(url string, execPath string, options ...string) error {
	args := append([]string{url}, options...)
	return exec.Command(execPath, args...).Start()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func outputDefaultConfig(filename string) (*config.Config, error) {
	defaultCfg := config.Default()
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal("outputDefaultConfig os.Open err:", err)
		return nil, err
	}
	err = yaml.NewEncoder(writer).Encode(defaultCfg)
	return &defaultCfg, err
}

func loadConfigForYaml() (*config.Config, error) {
	target := "config.yml"
	if !fileExists(target) {
		defaultCfg, err := outputDefaultConfig(target)
		return defaultCfg, err
	}
	f, err := os.Open(target)
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func main() {
	var args Args
	args.Template = flag.Int("template", 0, "Using template index")
	args.Dryrun = flag.Bool("dryrun", false, "Log output only")
	flag.Parse()

	cfg, err := loadConfigForYaml()
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		panic(err)
	}

	mainTemplate := cfg.Template[*args.Template]
	subject := mainTemplate.Subject
	body := mainTemplate.Body
	for _, targetVariable := range mainTemplate.Replacement {
		replaceNew, _ := targetVariable.ReplaceNew()
		subject = strings.Replace(subject, targetVariable.ReplaceTarget, *replaceNew, -1)
		body = strings.Replace(body, targetVariable.ReplaceTarget, *replaceNew, -1)
	}
	log.Print("Subject: " + subject)
	log.Print("Body: " + body)

	u, err := url.Parse(mainTemplate.Endpoint)
	if err != nil {
		log.Fatal("Error on endpoint url:", err)
		panic(err)
	}

	q := u.Query()
	q.Set("view", "cm")
	q.Set("fs", "1")
	q.Set("tf", "1")
	q.Set("to", mainTemplate.TO)
	q.Set("cc", mainTemplate.CC)
	q.Set("bcc", mainTemplate.BCC)
	q.Set("su", subject)
	q.Set("body", body)
	u.RawQuery = q.Encode()
	log.Print(u.String())

	if *args.Dryrun {
		os.Exit(0)
	}

	err = nil
	if cfg.Browser.CustomBrowserPath {
		log.Print("openCustomPath")
		err = open(u.String(), cfg.Browser.Path, cfg.Browser.Option)
	} else {
		log.Print("openDefaultBrowser")
		err = openDefaultBrowser(u.String())
	}
	if err != nil {
		log.Fatal("Failed open browser:", err)
		log.Panic(err)
	}

}
