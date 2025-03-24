#!/bin/bash

# .env の読み込み
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# 引数チェック
if [ -z "$1" ]; then
  echo "使い方: ./summarize.sh <URL>"
  exit 1
fi

# Goコードを実行
go run ./cmd/summarize "$1"
