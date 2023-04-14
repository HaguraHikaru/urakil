# urakil
URLを短縮させるCLIソフトウェア

## Description
URLは他者とWebページを共有する際に便利な形式であるが,ページによってURLが非常に長くなるものがある.
Bit.lyなどURLを短縮させるWebサービスがあるが,WebブラウザからURLを入力することは少々手間である.
本ソフトウェアでは入力したURLを,Bit.lyのAPIを利用し,短縮させて出力する.
CLIで動作させ,オプションによって複数のURLを一括で変換できるようにする.

## Usage
標準入出力でURLを受け取る.

--help
ヘルプの表示

-i　\n
指定したファイル名を受け取る

-o \n
txtファイルで出力.　引数はファイル名

urakil

URLを入力してください input:
https://www.kyoto-su.ac.jp/entrance/index-ksu.html

URLを短縮しました output:
https://www.kyoto-su.ac.jp
