# LearnALanguage

まずはじめにやること

## プログラミング言語 is 何

- 下記2つの概念がある
  - 仕様
    - ざっくりいうと設計
  - 処理系
    - ざっくりいうと実装
- e.g.) pythonという言語仕様をcpythonという処理系で実装
  - cpythonはC言語で書かれている
- コンパイル型、インタプリタ型

## 言語にはバージョンがある

- 言語のバージョンを変更すると依然問題なく動いたコードが動かなくなる等がある
- [anyenv](https://github.com/anyenv/anyenv)いれるとよさげ
  - 様々な言語のバージョン管理を統合的に行う事が出来るツール
  - pythonならpyenv、goならgoenv､rubyならrubyenvというバージョン管理ツールが元来存在していた。
  - これら全部を取りまとめ(anyenvを通してこれらのバージョン管理ツールを導入出来る)
