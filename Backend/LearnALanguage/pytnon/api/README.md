# api

## Memo

- Pythonにはお手軽にWebAPI作れるフレームワークが幾つかある
- flask bottleとかが昔からあるやつで、超お手軽系
- djangoも昔からあって、こいつは割と色々出来るから、ガチWeb開発というイメージ
- fastapiというのが個人的には好きで、OpenAPI?swagger?的なノリでAPIドキュメント自動生成してくれる機能も持ってる
- どの子もAPIサーバではあるけど、Webサーバ的なものと通信する必要がある
- pythonにはWSGI/ASGIというものがあり、これがWebサーバ的な役割及びpythonプログラムとの接続部分を担当する
- uvicornとかが有名
- APIサーバを起動する場合は、当該WSGIを起動する感じになるイメージ
- なお、例えばApacheサーバのうえでPython動かしたいよーとなったら、ApacheにWSGIのモジュールがあるみたいなので、そのモジュールをApacheへ入れてあげる的な仕組みよう
  - WSGIだけでも動くのだから、Apacheとかの上で動かすモチベーションは何かについては確認したい
  - 恐らくプロセス管理だとか、ログ周りだとか、そういう非機能部分で採用した方が良い何かがあるのかなと予想

## Web Flamework

- [fastapi](https://fastapi.tiangolo.com/ja/)

## WSGI

- TBU

## ASGI

- [uvicorn](https://www.uvicorn.org/)
- インスト―ルすると下記コマンドでWebサーバーを起動してくれる
  - `uvicorn main:app` にて起動
  - `uvicorn main:app --reload` とするとソースコードを変更するたびにWebサーバを再起動してくれる
