package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HaguraHikaru/urakil"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.5"

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
	-L --list-group     グループIDを取得
	-d --delete         短縮したURLを削除
	-h --help           ヘルプの表示
	-v --version        バージョン確認
	-f --input-file     ファイルを指定し,変換したURLを一括で標準出力
ARGUMENT
URL		短縮したいURLを指定します。引数は複数のURLを指定することができます
		引数が指定されていない場合、過去に作成した短縮済みURLの一覧を出力します`, prog)
}

type UrakilError struct {
	statusCode int
	message    string
}

func (e UrakilError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag     bool
	listGroupFlag  bool
	helpFlag       bool
	versionFlag    bool
	input_fileFlag bool
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

var completions bool

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

func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }

	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions")
	flags.MarkHidden("generate-completions")

	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "BitlyのAPIトークンを指定します。このオプションは必須です")
	//flags.StringVarP(&opts.runOpt.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	//flags.StringVarP(&opts.runOpt.config, "config", "c", "", "specify the configuration file.")
	//flags.StringVarP(&opts.runOpt.group, "group", "g", "", "specify the group name for the service. Default is \"urakil\"")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "グループIDを取得")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "短縮したURLを削除")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "ヘルプの表示")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "バージョン確認")
	flags.BoolVarP(&opts.flagSet.input_fileFlag, "input-file", "f", false, "ファイルを指定し,変換したURLを一括で標準出力")
	return opts, flags
}

func parseOptions(args []string) (*options, []string, *UrakilError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if completions {
		GenerateCompletion(flags)
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &UrakilError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &UrakilError{statusCode: 3, message: "トークンが与えられていません"}
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
	return urls
}

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
