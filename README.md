# urakil
![](https://img.shields.io/github/license/HaguraHikaru/urakil)

URLを短縮させるCLIソフトウェア

## Description
URLは他者とWebページを共有する際に便利な形式であるが,ページによってURLが非常に長くなるものがある.
Bit.lyなどURLを短縮させるWebサービスがあるが,WebブラウザからURLを入力することは少々手間である.
本ソフトウェアでは入力したURLを,Bit.lyのAPIを利用し,短縮させて出力する.
CLIで動作させ,オプションによって複数のURLを一括で変換できるようにする.

## 使い方
利用例  


    urakil 0.00
    URLを短縮させる

    USAGE:  
        urakil [OPTIONS] [FILE]   
        urakil <SUBCOMMAND>  
  
    OPTIONS:  
      -f, --input-file  
          ファイルを指定し,変換したURLを標準出力  
      --help  
          ヘルプの表示  
## インストール方法 
GitHubからコードをダウンロードする  
    `git clone git@github.com:HaguraHikaru/urakil.git`  
  
ダウンロードしたファイルurakilに移動しビルドコマンドを入力  
    `Makefile`  

## プロジェクトについて
  開発者 HaguraHikaru   
  ライセンス MIT LICENSE   
  アイコン icon.svg  
  名前の由来 Hikaruを逆順にし, 最後をlに変えただけ  
  バージョン履歴 0  
