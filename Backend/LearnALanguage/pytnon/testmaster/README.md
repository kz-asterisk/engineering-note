# testmaster

## Tool

- [unittest](https://docs.python.org/ja/3/library/unittest.html)
- [pytest](https://docs.pytest.org/en/7.0.x/getting-started.html)
  - [mock](https://webbibouroku.com/Blog/Article/pytest-mock)

## pytest

- とりあえずこれつかえばよさそう
- `poetry add -D pytest`
- `test_XXX` という名前で `def test_xxx():` でテスト関数を定義する
- `pytest` で実行してくれる
- `poetry add pytest-cov` でカバレッジ計測できる
  - pytest --cov=. --cov-report=html
    - htmlcovというディレクトリが作成される
    - この中のindex.htmlにテスト結果がある
    - [Ref.](https://hawksnowlog.blogspot.com/2020/08/how-to-up-test-coverage-with-pytest-cov.html)

## Memo

- [Github Action×pytest-cov](https://github.com/marketplace/actions/pytest-coverage-comment)
- [ユニットテストいろは](https://webbibouroku.com/Blog/Article/pytest-mock)
