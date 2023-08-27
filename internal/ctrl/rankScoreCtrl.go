package ctrl

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"rankBoardAns/internal/model"
	"rankBoardAns/internal/service"
)

// IncrScore 增加分数
func IncrScore(writer http.ResponseWriter, request *http.Request) {
	user := new(model.IncrScoreReq)
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(user)
	if err != nil {
		log.Print("incrScore: request fail")
		return
	}
	respText := "success"
	err = service.IncrScoreRank(user.Uid, user.Score)
	if err != nil {
		log.Print(err.Error())
		respText = "fail"
	}
	writer.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(writer).Encode(respText); err != nil {
		log.Print(err.Error())
	}
}

// GetScoreRankByUid 通过uid获取排行榜
func GetScoreRankByUid(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	uid := query.Get("uid")
	resp := &model.RankListResp{
		Status: "success",
	}
	rankList, err := service.GetRankListByUid(uid)
	if err != nil {
		log.Print(err.Error())
		resp.Status = "fail"
	}
	resp.RankList = rankList
	writer.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(writer).Encode(resp); err != nil {
		log.Print(err.Error())
	}
}
