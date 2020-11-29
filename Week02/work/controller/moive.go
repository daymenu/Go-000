package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daymenu/Go-000/Week02/work/biz"
	"github.com/daymenu/Go-000/Week02/work/model"
)

// SearchMoive 搜索喜欢的影片
func SearchMoive(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if name == "" {
		log.Println("name:" + name)
		return
	}
	searchMoive, err := biz.FindMoive(name)
	if err != nil {
		log.Println(err)
		return
	}

	response := model.SearchMoiveResponse{Code: 200, SearchMoive: searchMoive}
	var responseJSON []byte
	if responseJSON, err = json.Marshal(response); err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(responseJSON))

	log.Println("result:" + string(responseJSON))
}
