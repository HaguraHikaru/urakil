package main

import "testing"

func Example_completion() {
	goMain([]string{"./urakil", "--generate-completions"})
	// Output:
}

func Example_d() {
	goMain([]string{"./urakil", "-d"})
	// Output:
	// トークンが与えられていません
}

func Example_delete() {
	goMain([]string{"./urakil", "--delete"})
	// Output:
	// トークンが与えられていません
}

func Example_l() {
	goMain([]string{"./urakil", "-L"})
	// Output:
	// トークンが与えられていません
}

func Example_list_group() {
	goMain([]string{"./urakil", "--list-group"})
	// Output:
	// トークンが与えられていません
}

func Example_token() {
	goMain([]string{"./urakil", "--token"})
	// Output:
	// トークンが与えられていません
}

func Example_main() {
	goMain([]string{"./urakil", "-t"})
	// Output:
	// トークンが与えられていません
}

func Example_help() {
	goMain([]string{"./urakil", "--help"})
	// Output:
	// urakil [OPTIONS] [URLs...]
	// OPTIONS
	//	OPTIONS:
	//	-t --token          BitlyのAPIトークンを指定します。このオプションは必須です
	//	-L --list-group     グループIDを取得
	//	-d --delete         短縮したURLを削除
	//	-h --help           ヘルプの表示
	//	-v --version        バージョン確認
	//	-f --input-file     ファイルを指定し,変換したURLを一括で標準出力
	// ARGUMENT
	// URL		短縮したいURLを指定します。引数は複数のURLを指定することができます
	//		引数が指定されていない場合、過去に作成した短縮済みURLの一覧を出力します
}

func Test_main(t *testing.T) {
	if status := goMain([]string{"./urakil", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
