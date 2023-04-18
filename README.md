# urakil
URLを短縮させるCLIソフトウェア

## Description
URLは他者とWebページを共有する際に便利な形式であるが,ページによってURLが非常に長くなるものがある.
Bit.lyなどURLを短縮させるWebサービスがあるが,WebブラウザからURLを入力することは少々手間である.
本ソフトウェアでは入力したURLを,Bit.lyのAPIを利用し,短縮させて出力する.
CLIで動作させ,オプションによって複数のURLを一括で変換できるようにする.
標準入出力でURLを受け取る.

## 使い方

    Usage:  
        urakil [OPTIONS] [FILE]   
        urakil <SUBCOMMAND>  
  
    OPTIONS:  
      -f, --input-file  
          ファイルを指定し,変換したURLを標準出力  
      --help  
          ヘルプの表示  
## インストール方法 
  git clone  
  Makefile  

## プロジェクトについて
  開発者 HaguraHikaru   
  ライセンス MIT LICENSE   
  アイコン icon.svg  
  名前の由来 Hikaruを逆順にし, 最後をlに変えただけ  
  バージョン履歴 0  
