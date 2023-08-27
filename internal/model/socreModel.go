package model

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"rankBoardAns/pkg/goredis"
	"time"
)

var (
	//go:embed incrScoreRank.lua
	luaScript string
)

const (
	scoreRankKey = "scoreKey"
	timeMax      = int64(1e13)
)

type IncrScoreReq struct {
	Uid   string `json:"uid"`
	Score int    `json:"score"`
}

type RankInfo struct {
	Rank  int    `json:"rank"`
	Uid   string `json:"uid"`
	Score int    `json:"score"`
}

type RankListResp struct {
	Status   string      `json:"status"`
	RankList []*RankInfo `json:"rankList"`
}

func AddScoreRank(uid string, addScore float64) error {
	ctx := context.Background()
	timeRank := float64(timeMax-time.Now().UnixMilli()) * 1e-13
	if _, err := goredis.Rdb.Eval(ctx, luaScript, []string{scoreRankKey}, uid, addScore, timeRank).Result(); err != nil {
		return err
	}
	return nil
}

func GetRankByMember(uid string) (int, error) {
	ctx := context.Background()
	rank, err := goredis.Rdb.ZRevRank(ctx, scoreRankKey, uid).Result()
	if err != nil {
		return 0, err
	}
	return int(rank), nil
}

func GetRankListByRankRange(start, end int64) ([]redis.Z, error) {
	ctx := context.Background()
	list, err := goredis.Rdb.ZRevRangeWithScores(ctx, scoreRankKey, start, end).Result()
	if err != nil {
		return nil, err
	}
	return list, nil
}
