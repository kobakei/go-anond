# go-anond

Goで作る匿名ダイアリーみたいなアプリです。Goの練習用に作っています。

## 機能

- ユーザー登録不要で、匿名で日記を書けます。
- 他人が書いた日記に匿名でコメントできます。

## 各種ツール

- パッケージ管理: [Glide](https://glide.sh/)
- フレームワーク: [Echo](https://echo.labstack.com/)
- データベース: SQLite3
- ORM: [gorm](http://jinzhu.me/gorm/)

## コマンド

```sh
$ go run server.go

$ go test $(glide nv)
```

## Herokuへのデプロイ

```sh
$ heroku create
$ git push heroku master
```

## ライセンス

Apache License 2.0です。
