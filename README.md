## SetUp
### dep install
```
$ go get -u github.com/golang/dep/cmd/dep
```
外部パッケージをインストールした場合、以下のコマンドを打ってください。
```
$ dep ensure
```

### gorm install
```
go get github.com/jinzhu/gorm
```

### docker起動
```
docker-compose up
```

## migration
### goose install
```
$ go get bitbucket.org/liamstask/goose/cmd/goose
```
### migrationファイル作成
```
$ goose create 〇〇 sql
```

### migration実行
```
$goose up
```
