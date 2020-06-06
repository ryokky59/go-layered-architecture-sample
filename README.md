# layered-architecture-go-sample

## 概要
このリポジトリは[DDDを意識しながらレイヤードアーキテクチャとGoでAPIサーバーを構築する - Qiita](https://qiita.com/ryokky59/items/6c2b35169fb6acafce15)の解説用です。

dockerがインストールされていれば動かすことができます。

## Setup

1. `$ docker-compose build`

2. `$ docker-compose up`

## Example

- `$ curl -X POST "http://localhost:8080/task" -H "Content-Type: application/json" -d '{"title": "ほげ", "content": "ふが"}'`

- `$ curl -X GET "http://localhost:8080/task/1"`

- `$ curl -X PUT "http://localhost:8080/task/1" -H "Content-Type: application/json" -d '{"title": "hoge", "content": "huga"}'`

- `$ curl -X DELETE "http://localhost:8080/task/1"`
