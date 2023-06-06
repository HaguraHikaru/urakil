package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.3"

//const defURL = "https://news.google.com/home?hl=ja&gl=JP&ceid=JP%3Aja"

type options struct {
	token string
	//qrcode string
	//config string
	input_file bool
	help       bool
	version    bool
}

func fileRead(fileName *string) []string {
	urls := []string{}
	readFile, err := os.Open(*fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		urls = append(urls, fileScanner.Text())
	}
	readFile.Close()
	return urls
}

func bitlyRequest(opts *options, long_url *string) ([]byte, error) {
	fmt.Printf("long_url: %s\n", *long_url)
	//data := "test"

	json := fmt.Sprintf(`{"long_url": "%s", "domain": "bit.ly"}`, *long_url)
	requestBody := strings.NewReader(json)
	request, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", requestBody)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opts.token))
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return data, err
}

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.token, "token", "t", "", "specify the token for the service. This option is mandatory")
	//flags.StringVarP(&opts.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	//flags.StringVarP(&opts.config, "config", "c", "", "specify the configuration file")
	flags.BoolVarP(&opts.input_file, "input_file", "f", false, "ファイルを指定し,変換したURLを一括で標準出力")
	flags.BoolVarP(&opts.help, "help", "h", false, "ヘルプの表示")
	flags.BoolVarP(&opts.version, "version", "v", false, "バージョン確認")
	return opts, flags
}

func perform(opts *options, args []string) *UrakilError {
	if opts != nil {
		fmt.Printf("Token: %s\n", opts.token)
	}
	for _, long_url := range args {
		data, err := bitlyRequest(opts, &long_url)
		if opts != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%s\n", data)
		}

	}
	return nil
}

func parseOptions(args []string) (*options, []string, *UrakilError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args))
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	} else if opts.version {
		fmt.Println("version: " + VERSION)
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}

	if opts.token == "" {
		return nil, nil, &UrakilError{statusCode: 3, message: "no token was given"}
	}

	if opts.input_file {
		fileName := flags.Args()[0]
		if fileName[len(fileName)-4:] == ".txt" {
			return opts, fileRead(&fileName), nil
		}
	}
	return opts, flags.Args(), nil
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}

	return 0
}

func helpMessage(args []string) string {
	prog := "urakil"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
	OPTIONS:  
	-t --token          BitlyのAPIトークンを指定
	-h --help           ヘルプの表示  
	-v --version        バージョン確認
	-f, --input-file    ファイルを指定し,変換したURLを一括で標準出力 
ARGUMENT
URL     specify the url for shortening. this arguments accept multiple values.
	if no arguments were specified, urakil prints the list of available shorten urls.`, prog)
}

type UrakilError struct {
	statusCode int
	message    string
}

func (e UrakilError) Error() string {
	return e.message
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
