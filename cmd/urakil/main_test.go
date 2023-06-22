package main

import "testing"

func _example_main() {
	goMain([]string{"./urakil", "-t", "token"})
	// Output:
	// 過去に作成した短縮済みURLの一覧
}

func example_help() {
	goMain([]string{"./urakil", "--help"})
	// Output:
	// OPTIONS
	// 	OPTIONS:
	// 	-t --token          BitlyのAPIトークンを指定します。このオプションは必須です
	// 	-h --help           ヘルプの表示
	// 	-v --version        バージョン確認
	// 	-f, --input-file    ファイルを指定し,変換したURLを一括で標準出力
	// ARGUMENT
	// URL		短縮したいURLを指定します。引数は複数のURLを指定することができます。
	// 			引数が指定されていない場合、過去に作成した短縮済みURLの一覧を出力します。
}

func test_main(t *testing.T) {
	if status := goMain([]string{"./urakil", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
