# 目的
echo+psqlの組み合わせでapiを作成する

# API

## todo作成
POST /todos 
```
{
   "content":"aa" 
}
```
example
```
curl -XPOST localhost:1323/todos -d '{"content":"aa"}' -H 'Content-Type: application/json'
```

GET /todos/:id
```
{
  "ID": 1234,
  "Content": "contents for xxx",
  "Status": true,
  "UpdatedAt": "2004-10-20T10:23:54Z",
  "CreatedAt": "2004-10-19T10:23:54Z"
}

```
example
```
curl -XGET localhost:1323/todos/1234
```


# 備忘：未実装/Todo
- repositoryのテストを実サーバーではなくtestcontainerに変更
- postgresqlへの接続先を環境変数から渡すようにする。
- DELETE/SEARCHが未実装
- GETのidが無かった時のエラーハンドリングが未実施
- IDの型が数値/文字列で一貫性がない