# CharactorEncording

\========================================

文字コードについて

## 文字コード is 何？

- TBU

## おさえておくべき肯綮

- TBU

## よくあるシチュエーション

- Windowsで作成したファイルはSJIS
- Windows環境と密なDBの文字コード設定等でバグる

## 秘法集

- iconvコマンドのオプション

  - `-c`
    - 変換できなかった文字を出力しない
  - `-f --from-code`
    - 入力のエンコーディング
  - `-t --to--code`
    - 出力のエンコーディング

- 文字化けしたときに何の文字コードが当たりをつけたい

  - [化けて出た文字幽霊の正体を探る](http://labs.timedia.co.jp/2010/12/identifying-electronic-ghosts.html)

  - 文字化け状況を調査

    ```sh
    echo '文字化けした文字' |
    iconv -c \
      -f 恐らく文字化け前のエンコーディング \
      -t 文字化けで出力されているエンコーディング
    ```
