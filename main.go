package main

import (
	_ "embed"
	"net/http"
	"rankBoardAns/internal/ctrl"
	"rankBoardAns/pkg/goredis"
)

func main() {
	goredis.Setup()

	http.HandleFunc("/incrScore", ctrl.IncrScore)       // post 增加分数
	http.HandleFunc("/getRank", ctrl.GetScoreRankByUid) // get: 获取排行榜列表

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
