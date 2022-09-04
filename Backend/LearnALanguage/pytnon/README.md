# python is ニシキヘビ

## api

- fastapiとか使ってみたい([/api]([./api/)にまとめる)

## cicd

- 最低限のテストとデプロイについて([/cicd]([./cicd/)にまとめる)

## Syntax的な部分等、雑多はこちら

- pass

```pyhon
class MyEmptyClass:
     pass
```

- 関数の引数が\*\*はdictや\*はタプル
  - \*\*は最後じゃなきゃだめ

```python
def cheeseshop(kind, *arguments, **keywords):
    print("-- Do you have any", kind, "?")
    print("-- I'm sorry, we're all out of", kind)
    for arg in arguments:
        print(arg)
    print("-" * 40)
    for kw in keywords:
        print(kw, ":", keywords[kw])

cheeseshop("Limburger", "It's very runny, sir.",
           "It's really very, VERY runny, sir.",
           shopkeeper="Michael Palin",
           client="John Cleese",
           sketch="Cheese Shop Sketch")
```

- リスト内包表記

  - 無駄に変数作らない事が出来る

  ```python
  squares = list(map(lambda x: x**2, range(10)))
  squares = [x**2 for x in range(10)]
  ```

- from XXX import $$$

  - モジュール内の名前(XXX)を、import を実行しているモジュールのシンボルテーブル内に直接取り込む
  - $$$という名前の関数をそのまま実行できる
  - from XXX import *　で全部
  - import XXX as alias_XXX で名前変えて呼び出せる
  - form XXX import $$$ as alias\_\* で名前変えて呼び出せる

- モジュール検索パス

  - インタープリターはまずその名前のビルトインモジュールを探します
  - 見つからなかった場合は、 spam.py という名前のファイルを sys.path にあるディレクトリのリストから探します
  - sys.path
    - 入力されたスクリプトのあるディレクトリ (あるいはファイルが指定されなかったときはカレントディレクトリ)。
    - PYTHONPATH (ディレクトリ名のリスト。シェル変数の PATH と同じ構文)。
    - The installation-dependent default (by convention including a site-packages directory, handled by the site module).
      - /usr/local/lib/pythonX.Y/site-packages/bar
        -/usr/local/lib/pythonX.Y/site-packages/foo

- 組込み関数 dir() は、あるモジュールがどんな名前を定義しているか調べるために使われます

- パッケージ (package) は、Python のモジュール名前空間を "ドット付きモジュール名" を使って構造化する手段

  - モジュールを利用すると、別々のモジュールの著者が互いのグローバル変数名について心配しなくても済むようになる

  ```dir
  sound/                        Top-level package
      __init__.py               Initialize the sound package
      formats/                  Subpackage for file format conversions
              __init__.py
              wavread.py
              wavwrite.py
              aiffread.py
              aiffwrite.py
              auread.py
              auwrite.py
              ...
      effects/                  Subpackage for sound effects
              __init__.py
              echo.py
              surround.py
              reverse.py
              ...
      filters/                  Subpackage for filters
              __init__.py
              equalizer.py
              vocoder.py
              karaoke.py
              ...
  ```

  - パッケージを import する際、 Python は sys.path 上のディレクトリを検索して、トップレベルのパッケージの入ったサブディレクトリを探します。
  - ファイルを含むディレクトリをパッケージとしてPython に扱わせるには、ファイル __init__.py が必要
  - __init__.py ではパッケージのための初期化コードを実行したり、後述の __all__ 変数を設定してもかまいません。
    - sound/effects/__init__.py
      - `__all__ = ["echo", "surround", "reverse"]`

- 警告 f.write() を with キーワードや f.close() を使わずに呼び出した場合、プログラムが正常に終了した場合でも、 f.write() の実引数がディスクに完全に 書き込まれないことがあります 。

- エラーには (少なくとも) 二つのはっきり異なる種類が

  - 構文エラー (syntax error)
  - 例外 (exception)

- try: except

  - 通常except発生したらプログラムが異常終了するが、そうならないようするとかできる
  - finally: でクリーンナップ動作を定義できる

- raise 文を使って、特定の例外を発生させる

  - raise の唯一の引数は例外インスタンスか例外クラス (Exception を継承したクラス)
  - 例外が発生したかどうかを判定したいだけで、その例外を処理するつもりがなければ、単純な形式の raise 文を使って例外を再送出させることができます:

- クラスを理解するために

  - 名前空間 (namespace) とは、名前からオブジェクトへの対応付け (mapping)
    - ほとんどの名前空間は、現状では Python の辞書として実装されています
    - 二つの別々のモジュールの両方で関数 maximize という関数を定義することができ、定義自体は混同されることはありません
  - 属性 という言葉は、ドットに続く名前すべてに対して使っています
    - 例えば式 z.real で、 real はオブジェクト z の属性
    - 厳密にいえば、モジュール内の名前に対する参照は属性の参照
    - __main__ という名前のモジュールの一部分であるとみなされるので、独自の名前空間を持つことになります。
  - スコープ (scope) とは、ある名前空間が直接アクセスできるような、 Python プログラムのテキスト上の領域

  ```python
  def scope_test():
    def do_local():
        spam = "local spam"

    def do_nonlocal():
        nonlocal spam
        spam = "nonlocal spam"

    def do_global():
        global spam
        spam = "global spam"

    spam = "test spam"
    do_local()
    print("After local assignment:", spam)
    do_nonlocal()
    print("After nonlocal assignment:", spam)
    do_global()
    print("After global assignment:", spam)

    scope_test()
    print("In global scope:", spam)
    このコード例の出力は:

    After local assignment: test spam
    After nonlocal assignment: nonlocal spam
    After global assignment: nonlocal spam
    In global scope: global spam
  ```

  - クラスが __init__() メソッドを定義している場合、クラスのインスタンスを生成すると、新しく生成されたクラスインスタンスに対して自動的に __init__() を呼び出します

    - __init__() メソッドに複数の引数をもたせることができます

    ```python
    >>> class Complex:
    ...     def __init__(self, realpart, imagpart):
    ...         self.r = realpart
    ...         self.i = imagpart
    ...
    >>> x = Complex(3.0, -4.5)
    >>> x.r, x.i
    (3.0, -4.5)
    ```

### Tips

- CLIからpythonファイルの関数を直接実行

  ```python
  python -c "import main; main.chometaro()";
  ```

  - import XXXのXXXはファイル名(上述は同一ディレクトリ上のmain.pyの例)

- venv

  - 普段poetryにて暗黙的に作っているので敢えてやってみる
  - 仮想環境作成

    ```bash
    python -m venv .venv
    ```

    - `.venv` が仮想環境名
    - poetryがこの名前でディレクトリ作るしgitignore的にもこれがよさげ
    - -m の意味

      ```text
      -m mod : run library module as a script (terminates option list)
      ```

      - 標準でインストールされているモジュールは[ここ](https://docs.python.org/3/py-modindex.html)にある
        - venvもある
    - 仮想環境起動

      ```bash
      source .venv/bin/activate
      ```

    - 仮想環境停止

      ```bash
      deactivate
      ```

- Pipenv

  - 仮想環境作成とライブラリ管理両方を行うツール
  - `pip install pipenv` でインストール
  - `pipenv install ${pythonバージョン}` で仮想環境作成
    - こんな感じで仮想環境出来上がる

      ```bash
      /home/kz/.local/share/virtualenvs/case_pipenv-GKwsGtJV
      ```

  - `pipenv shell` で起動
  - `pipenv install pandas` のように外部モジュールをインストール
    - `Pipfile` と `Pipfile.lock` が生成される
    - Pipfile
      - pythonバージョンやPipenv経由でインストールしたモジュールのバージョン情報がtomlで記載されている
    - Pipfile.lock
      - 依存モジュールについて記載されているもの
    - pipenv sync を使うと意図しないアップデートがなされない見込み
      - [参考になる資料](https://dev.classmethod.jp/articles/pipenv-sync-is-useful/)

- poetry
  - 仮想環境作成とライブラリ管理両方を行うツール
  - Pipenvとの違いはPEP518に従ったpyproject.tomlを利用するところ
    - モジュール以外にメタデータやビルド設定などをまとめて書ける
    - [Ref.](https://peps.python.org/pep-0518/#specification)
  - `poetry new XXX` でプロジェクト作成
    - testディレクトリ等々含めて自動で生成してくれる
  - `poetry init` で既存プロジェクトをpoetry管理できる
  - `poetry add` でモジュール追加
    - pyproject.tomlにモジュール名が書かれる
    - poetry.lock に依存関係のパッケージが記載される
  - `poetry shell` で起動
  - `poetry run pyton XXX` で実行
  - `poetry run pyton pytest` とかも実行可能

- クラスにメソッド作る時はselfを引数に入れる
  - 入れない場合、外部から呼び出した際に引数エラー(takes 0 positional arguments but 1 was given)が発生
  - 外部からメソッドを呼ぶときに暗黙的にselfを渡しているため
  - [Ref.](https://office54.net/python/error/python-typeerror-argument)

- formatは[PEP8](https://peps.python.org/pep-0008/)
