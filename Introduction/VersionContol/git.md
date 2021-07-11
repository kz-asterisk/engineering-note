# KnowHow

## まず知ったりやってみた方がよさそうだなと思ったこと

### ざっくり概念と用語を知る

- まずは公式docを読む
  - <https://git-scm.com/book/ja/v2/%E4%BD%BF%E3%81%84%E5%A7%8B%E3%82%81%E3%82%8B-Git%E3%81%AE%E5%9F%BA%E6%9C%AC>
- gitはバージョン管理ツールの一つ
- 分散して同一コードへの変更を複数の人で行う際に、データの`完全性`を担保してくれる(ぶっこわれなようにしてくれる)
- git には`リモートリポジトリ`と呼ばれるコードの大本のマスター情報を補完する場所がある
- `リモートリポジトリ`を作れるところ(リポジトリホスティングサービス)は以下の通り

  ```text
  世間に一般公開されるもの系
  ・GitHub
  　※プライべート設定にすると公開されない

  自身/会社でサーバを立てて世の中に公開しない系
  ・gitlab
  ・BitBucket
  ・GitHubEnterprise

  ※BitBucketはクラウドサービス等もあるよう
  ```

- gitには`ブランチ`という概念が存在する
- 一般的に大本のマスター情報を管理する`マスターブランチ(master)`が存在する
- `リモートブランチ`のマスターブランチをベースに修正作業を行う
- 「ローカル」「ステージング」「リモート」という3つのエリアを変更内容やメタデータ等の情報が行き来することで変更管理を行っている

### 実際にGitHubでリポジトリを作成して使ってみる

- 自身の端末にgitをインストールする
- まずは準備でgitで名前とメアドを設定

  ```bash
  git config --global user.name "XXX"
  git config --global user.email XXX@XXX.com
  ```

- GitHubでサインアップしてアカウント作成する
  - <https://github.com/>
  - 二要素認証忘れずに
- Web画面でリポジトリを作成する
  - これが`リモートリポジトリ`になる
- `リモートリポジトリ`の`マスターブランチ(master)のコードをローカルコピーする

  ```bash
  git clone URL※
  ```

  ※URLは自身の`リモートリポジトリ`をWebで確認する際に`↓Code`というボタンを押すとURLが表示される(※httpsのものを選択すれば良い)

- ローカルへ作業用ブランチ※を作成する

  ```bash
  git checkout -b 作業用ブランチ名
  ```

  ※一般的には`feature/作業チケット名` 等が多い

  ```text
  お仕事では修正や開発タスクがチケット管理ツール等で一意のチケット番号を発行が付与されて担当者にアサインされるケースが多い。

  後でこの「作業は何のためにやったのか？」等の変更の文脈を知りたいときに便利(のちほど触れるがgitでは変更する際に、コミットメッセージ等文脈を残す事が容易で且つ、参照性も高いという印象がある)
  ```

- 作業用ブランチで編集を行う
- 編集した内容を`gitに認知させる`
  - `ステージングエリア`とうところに変更内容を登録する

    ```bash
    git status
    git add -A
    git status
    ```

    ※変更を確定させる前提として`ステージングエリア`への登録が必須

    ```text
    「-A」はすべての変更・追加・削除を反映するというオプション。他にも「特定のディレクトリへの変更だけ」や「追加だけ」とか「変更だけ」というオプションがある。
    ```

  - 変更を確定する

    ```bash
    git commit -m "XXX"
    ```

    ※XXXに変更の理由について文章を書く(Fix XXXなど)

    ※これでローカルの作業用ブランチの範囲のみへ変更内容が認知されたことになる(リモートリポジトリや他のローカルブランチへは何も変更はなされていない)

- `gitへ認知`させた編集内容を`リモートリポジトリの作業用ブランチ`へ反映させる(※反映させる際に`リモートリポジトリ`へ`作業用ブランチ`が作成される)
  - `ローカルの作業用ブランチ`を`リモートリポジトリの作業用ブランチ`へアップロードする

  ```bash
  git push origin feature/xxx
  ```

  ※`origin`は`リモートリポジトリ`表すもの(デフォルト名がorigin)

- `リモートリポジトリの作業用ブランチ`をマスターブランチ(master)へ反映させる`申請`を行う
  - `申請`は`プルリクエスト`と呼ばれる
  - Web画面上でGUI操作しても良いし、git push時に表示されるプルリクエスト用URLからアクセスしてGUI操作しても良い
  - 変更内容を反映する先のブランチを`宛先ブランチ`へ指定する
    - このチュートリアル的には`master`を指定
  - この別のブランチへ変更内容を統合することを`マージ`という

## git flow と github flowについてさらっと学ぶ

- branch名の由来とかgit運用の一般的な文化に触れて
  - <https://qiita.com/Yu-kiFujiwara/items/40b503683d6525c8d274>

## 実際のお仕事で使ったこと

### .gitconfig

- plogというaliasオプションを作成してlogを見やすくする

  ```bash
  [alias]
          plog = log --pretty=
          'format:%C(yellow)%h %C(green)%cd %C(reset)%s 
          %C(red)%d %C(cyan)[%an]' --date=iso
  ```

### upstream-branch(追跡ブランチ)

- `git fetch`や`git pull`を行う際、ローカルのブランチに対応するデフォルトのリモートのブランチを指定するもの

- これにより`git fetch`や`git pull`を行う際にリモートブランチの指定を省略可能

- `git branch -vv` にて確認可能

  ```bash
  master 1d2acad [origin/master] update README.md
  ```

  ※この例では`master`がローカルブランチで対応するリモートブランチが`origin/master`

- なおリモートリポジトリの情報を得るには`git remote -v`

  ```bash
  origin  https://github.com/kz-kz/engineering-note.git (fetch)
  origin  https://github.com/kz-kz/engineering-note.git (push)
  ```

### 操作取り消し系

- stageされていないファイルの変更を全て取り消し

  ```bash
    git checkout .
  ```

- stage状態 → add前状態

  ```bash
  git reset HEAD file_name
  ```

- commit → stage状態

  ```bash
  git rest --soft HEAD~
  ```

- 直前のコミット修正

  ```bash
  git commit --amend
  ```

### cherry-pick

- 特定のブランチの特定のコミットを任意のブランチのコミットとして反映させる
