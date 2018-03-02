package postToDingDing

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"net/url"
)

var timeLayout = "2006-01-02 00:00:00"

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ReqData struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}




func PostYesterDay() {
	postUrl := "https://oapi.dingtalk.com/robot/send?access_token=d4ec0b6bc77e4a0d9d509641aa1b9107cec0564c5c4713bb1d036ec6d7de035a"
	bodyType := "application/json;charset=utf-8"

	startDate := time.Now().AddDate(0, 0, -1).Format(timeLayout)
	endDate := time.Now().Format(timeLayout)

	fmt.Println("start", startDate)
	fmt.Println("end", endDate)

	//yesterDayPv
	pvUrl := "https://log.ideapar.com/countQuery?query=select count(*) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\""
	uvUrl := "https://log.ideapar.com/countQuery?query=select count(distinct userIp) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\""
	newIpUrl := "https://log.ideapar.com/countQuery?query=select count(distinct userIp) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\""
	rqtUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-rqt"
	desingerUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-desinger"
	prooferUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-proofer"

	u, err := url.Parse(pvUrl)
	q := u.Query()
	u.RawQuery = q.Encode()
	//pv
	resp, err := http.Get(u.String())
	checkErr("resp", err)
	pvNum, err := ioutil.ReadAll(resp.Body)
	checkErr("ioutil", err)

	fmt.Println("res", string(pvNum))
	return
	//uv
	//newIp
	//designer
	//proofer

	data := &ReqData{
		Msgtype: "markdown",
		Markdown: Markdown{
			"昨日数据",
			"### 昨日数据 \r\n" +
				"- **pv**:1223 \r\n" +
				"- **uv**:233 \r\n" +
				"- **newIp**:2 \r\n" +
				"- **新增设计师**:23 \r\n" +
				"- **新增打样师**：23 \r\n",
		},
	}

	jsonData, err := json.Marshal(data)
	fmt.Println("string", string(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	req := bytes.NewBuffer(jsonData)

	resp, err = http.Post(postUrl, bodyType, req)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func lastWeek() {

}

func checkErr(remark string, err error) {
	fmt.Println(remark, err)
}
