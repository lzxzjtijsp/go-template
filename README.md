# 起動
```
go run main.go
```

# Dockerビルド
```
docker build -t my-go-app .
```

# Dockerコンテナの起動
```
docker run -p 8080:8080 my-go-app
```

# Docker Composeビルド
```
docker-compose build
```

# Docker Compose起動
```
docker-compose up -d --build --remove-orphans
```

# Docker Compose mysql接続
```
docker-compose exec mysql mysql -h 127.0.0.1 -P 3306 -u root -p
```

# Docker Compose停止
```
docker-compose down -v 
docker-compose rm -v
```

# リクエスト
```
curl -H "Authorization: Bearer Your-Expected-Value" http://localhost:8080/
```
