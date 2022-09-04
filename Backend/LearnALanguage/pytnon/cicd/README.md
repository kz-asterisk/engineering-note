# cicd

## Memo

- とりあえずやった方がいいこと
- pyenvでバージョン管理
- poetryでパッケージ管理&vevn
- devのdependenciesには下記を入れる
  - poetry add -D pytest autoflake isort black
- あとは型アノテーションで安全にをお手軽開発するため下記も入れる
  - poetry add pydantic
