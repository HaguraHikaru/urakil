package main

import (
	"testing"
)

func _Example_Main() {
	goMain([]string{"./urakil", "-t", "token"})
	// Output:
	// Hello World
}

func Example_Help() {
	goMain([]string{"./urakil", "--help"})
	// Output:
	// urakil [OPTIONS] [URLs...]
	// OPTIONS
	//-t --token          BitlyのAPIトークンを指定
	//-h --help           ヘルプの表示
	//-v --version        バージョン確認
	//-f, --input-file    ファイルを指定し,変換したURLを一括で標準出力
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, urakil prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./urakil", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

/*
func TestbitlyRequest(t *testing.T) {
	config := NewConfig(os.Getenv("URAKIL_TOKEN"), Shorten) //bitly := NewBitly("")
	testdata := []struct {
		giveUrl          string
		wontShortenError bool
		wontDeleteError  bool
	}{
		{"https://news.google.com/home?hl=ja&gl=JP&ceid=JP%3Aja", false, false},
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

*/
