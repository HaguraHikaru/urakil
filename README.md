# urakil
URLを短縮させるCLIソフトウェア

## Description
URLは他者とWebページを共有する際に便利な形式であるが,ページによってURLが非常に長くなるものがある.
Bit.lyなどURLを短縮させるWebサービスがあるが,WebブラウザからURLを入力することは少々手間である.
本ソフトウェアでは入力したURLを,Bit.lyのAPIを利用し,短縮させて出力する.
CLIで動作させ,オプションによって複数のURLを一括で変換できるようにする.

#Usage
--help

-i
ファイル内のURLを変換

-o
指定したファイル名で出力

urakil

input:
https://www.kyoto-su.ac.jp/entrance/index-ksu.html

output:
https://www.kyoto-su.ac.jp
