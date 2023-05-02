# urakil
[![Coverage Status](https://coveralls.io/repos/github/HaguraHikaru/urakil/badge.svg?branch=main)](https://coveralls.io/github/HaguraHikaru/urakil?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/HaguraHikaru/urakil)](https://goreportcard.com/report/github.com/HaguraHikaru/urakil) 
[![codebeat badge](https://codebeat.co/badges/85e23949-4905-4960-8ea9-cf87f2f2f708)](https://codebeat.co/projects/github-com-hagurahikaru-urakil-main)  

![](https://img.shields.io/github/license/HaguraHikaru/urakil)  

URLを短縮させるCLIソフトウェア

# 🚪概要
URLは他者とWebページを共有する際に便利な形式であるが,ページによってURLが非常に長くなるものがある.
Bit.lyなどURLを短縮させるWebサービスがあるが,WebブラウザからURLを入力することは少々手間である.
本ソフトウェアはCLI上で動作し,入力したURLをBit.lyのAPIを利用し,短縮させて出力する.


# 📖使い方
    USAGE:  
        urakil [OPTIONS] [URL...]
        urakil [OPTIONS] [FILE]  
                  
    OPTIONS:  
      -h --help  
          ヘルプの表示  
      -v --version  
          バージョン確認
      -f, --input-file  
          ファイルを指定し,変換したURLを一括で標準出力  

      
# ✈️インストール方法 
GitHubからコードをダウンロードする  
    `git clone git@github.com:HaguraHikaru/urakil.git`  
  
urakilに移動しビルドコマンドを入力  
    `Makefile`  

# 😄プロジェクトについて
  ## 開発者
  HaguraHikaru   
  ## ライセンス  
  MIT LICENSE   
  ![](https://img.shields.io/github/license/HaguraHikaru/urakil)
  ## アイコン 
  遺伝子組み換えのようなものをイメージ   
   <img src="icon.svg" width="20%" />  
  ## 名前の由来  
  Hikaruを逆順にし, 最後をlに変えただけ  
