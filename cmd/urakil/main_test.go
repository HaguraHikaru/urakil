package main

import (
	"os"
	"testing"
)

func _Example_Main() {
	goMain([]string{"./urleap", "-t", "token"})
	// Output:
	// Hello World
}

func Example_Help() {
	goMain([]string{"./urakil", "--help"})
	// Output:
	// urleap [OPTIONS] [URLs...]
	// OPTIONS
	//     -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
	//     -q, --qrcode <FILE>      include QR-code of the URL in the output.
	//     -c, --config <CONFIG>    specify the configuration file.
	//     -g, --group <GROUP>      specify the group name for the service. Default is "urleap"
	//     -d, --delete             delete the specified shorten URL.
	//     -h, --help               print this mesasge and exit.
	//     -v, --version            print the version and exit.
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, urleap prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./urleap", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

func TestbitlyRequest(t *testing.T) {
	config := NewConfig(os.Getenv("URAKIL_TOKEN"), Shorten) //bitly := NewBitly("")
	testdata := []struct {
		giveUrl          string
		wontShortenError bool
		wontDeleteError  bool
	}{
		{"https://tamadalab.github.io/", false, false},
	}
	for _, td := range testdata {
		result, err := bitly.Shorten(config, td.giveUrl)
		if (err == nil) != td.wontShortenError {
			t.Errorf("shorten %s wont error %t, but got %t", td.giveUrl, td.wontShortenError, !td.wontShortenError)
		}
		err = bitly.Delete(config, result.Shorten)
		if (err == nil) != td.wontDeleteError {
			t.Errorf("delete %s wont error %t, but got %t", result.Shorten, td.wontDeleteError, !td.wontDeleteError)
		}
	}
}
