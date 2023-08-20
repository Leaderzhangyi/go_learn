package downloader

import "net/http"

type InfoRequest struct {
	Bvids []string
}

type VideoInfo struct {
	Code int `json:"code"`
	Data struct {
		Bvid  string `json:"bvid"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"data"`
}

type InfoResponse struct {
	Infos []VideoInfo
}

func BatchDownloadVideoInfo(request InfoRequest) (InfoResponse, error) {
	var response InfoResponse
	for _, bvid := range request.Bvids {
		var videoInfo VideoInfo
		resp, err := http.Get("https://api.bilibili.com/x/web-interface/view?bvid=" + bvid)

	}
}
