# module and package

ライブラリやらパッケージやら適当にやってしまっているので少しだけ調べてみる

## ライブラリ？

- 厳密な定義はない？
- モジュールやパッケージの事を指す

## モジュール？

- 所謂XXX.py
- 関数やクラスが纏まっている

## パッケージ？

- モジュールを集めたもの

## pip

- パッケージを管理するツール
- インストール
  - pip install
- インストールしたファイル群
  - pip freeze > requirements.txt
  - 依存モジュールも併せて記録される
    - バージョンは固定
- インストールの場所(anyanv & pyenvを利用しているケース)
  - pipを起動したpythonのsite-packages配下にインストールされるパターン
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/lib/python3.10/site-packages/pandas-1.4.4.dist-info/*
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/lib/python3.10/site-packages/pandas/*
  - pipを起動したpythonのsite-packages配下にインストールされるパターンとbin配下にインストールされるパターン
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/bin/py.test
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/bin/pytest
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/lib/python3.10/site-packages/_pytest/*
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/lib/python3.10/site-packages/pytest-7.1.3.dist-info/*
    /home/kz/.anyenv/envs/pyenv/versions/3.10.2/lib/python3.10/site-packages/pytest/*
    - bin配下にある場合は`python XXX` とせずに直接実行できる
