package service

import (
	"rankBoardAns/internal/model"
)

func IncrScoreRank(uid string, addScore int) error {
	return model.AddScoreRank(uid, float64(addScore))
}

func GetRankListByUid(uid string) ([]*model.RankInfo, error) {
	rank, err := model.GetRankByMember(uid)
	if err != nil {
		return nil, err
	}
	start, end := rank-10, rank+10
	if start < 0 {
		start = 0
	}
	list, err := model.GetRankListByRankRange(int64(start), int64(end))
	if err != nil {
		return nil, err
	}
	rankList := make([]*model.RankInfo, 0, len(list))
	for i, rankInfo := range list {
		rankList = append(rankList, &model.RankInfo{
			Rank:  start + i + 1,
			Uid:   rankInfo.Member.(string),
			Score: int(rankInfo.Score),
		})
	}
	return rankList, nil
}
