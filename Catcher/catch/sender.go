package catch

import (
	"FindGhost/Catcher/models"
	"FindGhost/Catcher/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// 发送httpInfo到审计系统
func SendHttpInfo(httpInfo *models.HttpInfo) error {
	httpInfoJson, err := json.Marshal(httpInfo)
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	secureKey := util.MakeSign(timeStamp, ApiKey)
	apiUrl := fmt.Sprintf("%v%v", ApiUrl, "/api/http")
	_, err = http.PostForm(apiUrl, url.Values{"timeStamp": {timeStamp}, "secureKey": {secureKey}, "data": {string(httpInfoJson)}})
	return err
}
