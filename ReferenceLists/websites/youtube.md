# youtube

========================================

youtube動画で得た知見のメモ

## メモ
- [SPA CMSツールVolto for Ploneの紹介とPWAへの展望 2020-4-25 A-4](https://www.youtube.com/watch?v=jB73SPyHhLU&t=2391s)
  - リッチなUX(GmailとかGoogleカレンダー)を作るためには、既存のWeb技術では大変
    - ページの一部を書き換えることができない
    - 状態を持つことができない
  - これを解決する方法としてSPA(Single Page Application)
  - SPAとは何か
    - JavaScriptを用いたWebアプリの実装
    - サーバから1ぺージのHTMLを返し、ルーティング機能を持つ
    - JavaScriptで作られたアプリを読み込む
    - REST APIでJsonをやりとり
    - スマホアプリのような動き
      ```
      通常のWebサイトの場合、次のぺージを閲覧する際に、ハイパーリンクを用いて、
      サーバに別のページのリクエストを送信するが、SPAではアプリ内でルーティングし、
      APIをコールなどして別の情報が表示される。
      ```
- [AIから不公平なバイアスを取り除く AI Fairness 360 Open Source Toolkit 2020-4-24 B-4](https://www.youtube.com/watch?v=NxdlTZ2ZiFA)
  - DeepLerningの学習過程でバイアスが発生する
  - AI Fairness という OSSでバイアスを取り除く事が出来る
  - IBMにはこの分野の商用製品がある
  - AI White Box (AIが出力した結果についてなぜそのような結果になったのか説明)という考え方がある

