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