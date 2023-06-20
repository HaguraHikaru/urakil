package main

import (
	"bufio"
	"fmt"
	//"io"
	//"io/ioutil"
	//"log"
	//"net/http"
	"os"
	"path/filepath"
	//"strings"

	flag "github.com/spf13/pflag"
	"github.com/HaguraHikaru/urakil"
)

const VERSION = "0.1.4"

//const defURL = "https://news.google.com/home?hl=ja&gl=JP&ceid=JP%3Aja"

func versionString(args []string) string {
	prog := "urakil"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

func helpMessage(args []string) string {
	prog := "urakil"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
	OPTIONS:  
	-t --token          BitlyのAPIトークンを指定します。このオプションは必須です
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

/*
type options struct {
	token string
	//qrcode string
	//config string
	input_file bool
	help       bool
	version    bool
}*/

type flags struct {
	deleteFlag    bool
	listGroupFlag bool
	helpFlag      bool
	versionFlag   bool
	input_fileFlag    bool
}

type runOpts struct {
	token  string
	qrcode string
	config string
	group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) urakil.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return urakil.ListGroup
	case len(args) == 0:
		return urakil.List
	case opts.flagSet.deleteFlag:
		return urakil.Delete
	case opts.runOpt.qrcode != "":
		return urakil.QRCode
	default:
		return urakil.Shorten
	}
}
/*
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
}*/

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "BitlyのAPIトークンを指定します。このオプションは必須です")
	flags.StringVarP(&opts.runOpt.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	flags.StringVarP(&opts.runOpt.config, "config", "c", "", "specify the configuration file.")
	flags.StringVarP(&opts.runOpt.group, "group", "g", "", "specify the group name for the service. Default is \"urleap\"")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. This is hidden option.")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "ヘルプの表示")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "バージョン確認")
	flags.BoolVarP(&opts.flagSet.input_fileFlag, "input_file", "f", false, "ファイルを指定し,変換したURLを一括で標準出力")
	return opts, flags
}

func parseOptions(args []string) (*options, []string, *UrakilError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &UrakilError{statusCode: 3, message: "no token was given"}
	}
	if opts.flagSet.input_fileFlag {
		fileName := flags.Args()[0]
		if fileName[len(fileName)-4:] == ".txt" {
			return opts, fileRead(&fileName), nil
		}
	}
	return opts, flags.Args(), nil
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
	fmt.Println(urls)
	return urls
}
/*
func retrieveLinks(opts *options) ([]byte, error) {
	request, err := http.NewRequest("GET", "https://api-ssl.bitly.com/v4/bitlinks/bit.ly/12a4b6c", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opts.token))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data, err

}*/

/*
func bitlyRequest(opts *options, long_url *string) ([]byte, error) {
	//fmt.Printf("long_url: %s\n", *long_url)
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
}*/

func shortenEach(bitly *urakil.Bitly, config *urakil.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func deleteEach(bitly *urakil.Bitly, config *urakil.Config, url string) error {
	return bitly.Delete(config, url)
}

func listUrls(bitly *urakil.Bitly, config *urakil.Config) error {
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}

func listGroups(bitly *urakil.Bitly, config *urakil.Config) error {
	groups, err := bitly.Groups(config)
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}

func performImpl(args []string, executor func(url string) error) *UrakilError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *UrakilError {
	bitly := urakil.NewBitly(opts.runOpt.group)
	config := urakil.NewConfig(opts.runOpt.config, opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case urakil.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case urakil.ListGroup:
		err := listGroups(bitly, config)
		return makeError(err, 2)
	case urakil.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case urakil.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(bitly, config, url)
		})
	}
	return nil
}


/*
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
}*/

func makeError(err error, status int) *UrakilError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*UrakilError)
	if ok {
		return ue
	}
	return &UrakilError{statusCode: status, message: err.Error()}
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

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
