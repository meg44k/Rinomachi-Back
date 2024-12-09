#!/bin/bash

REQUESTS=(
    "GET http://localhost:8080/users"
    "POST http://localhost:8080/users"
    "GET http://localhost:8080/users/id"
    "DELETE http://localhost:8080/users/id"
    "GET http://localhost:8080/buildings"
    "POST http://localhost:8080/buildings"
    "GET http://localhost:8080/buildings/id"
    "GET http://localhost:8080/users/id/favorites"
    "DELETE http://localhost:8080/users/id/favorites/id"
    "GET http://localhost:8080/users/id/histories"
    "DELETE http://localhost:8080/users/id/histories/id"
)

# リクエストヘッダー
HEADERS=(
  "-H" "Content-Type: plain/text"
)

# 各ペアを処理するループ
for REQUEST in "${REQUESTS[@]}"; do
  # URLとメソッドを分割
  METHOD=$(echo "$REQUEST" | awk '{print $1}')
  URL=$(echo "$REQUEST" | awk '{print $2}')

  echo "--Sending $METHOD request to $URL--"

  # cURLコマンドを実行
  response=$(curl -s -o /dev/null -w "%{http_code}" -X "$METHOD" "${HEADERS[@]}" "$URL")

  # ステータスコードを確認
  if [ "$response" -eq 200 ]; then
    echo "O $METHOD request to $URL successful! Status code: $response"
  else
    echo "X $METHOD request to $URL failed. Status code: $response"
  fi
done
