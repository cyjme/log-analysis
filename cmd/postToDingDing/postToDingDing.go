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

func getNumByUrl(urlString string) string {
	u, err := url.Parse(urlString)
	q := u.Query()
	u.RawQuery = q.Encode()
	//pv
	resp, err := http.Get(u.String())
	checkErr("resp", err)
	num, err := ioutil.ReadAll(resp.Body)
	checkErr("ioutil", err)

	return string(num)
}

func PostYesterDay() {
	postUrl := "https://oapi.dingtalk.com/robot/send?access_token=2b4f269d94f7cc08b76dbdd2e156e23bf31a19b61d33a58301d77641ce77cfae"
	bodyType := "application/json;charset=utf-8"

	startDate := time.Now().AddDate(0, 0, -1).Format(timeLayout)
	endDate := time.Now().Format(timeLayout)

	fmt.Println("start", startDate)
	fmt.Println("end", endDate)

	//yesterDayPv
	pvUrl := "https://log.ideapar.com/countQuery?query=select count(*) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\""
	uvUrl := "https://log.ideapar.com/countQuery?query=select count(distinct userIp) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\""
	newIpUrl := "https://log.ideapar.com/countQuery?query=select count(distinct userIp) from pv where date > \"" + startDate + "\" and date < \"" + endDate + "\"" + "and userIp not in (select userIp from pv where date < \""+startDate+"\")"
	rqtUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-rqt"
	desingerUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-desinger"
	prooferUrl := "https://www.ideapar.com/fetchOperateData?queryType=yesterday-proofer"

	pvNum := getNumByUrl(pvUrl)
	uvNum := getNumByUrl(uvUrl)
	newIpNum := getNumByUrl(newIpUrl)
	rqtNum := getNumByUrl(rqtUrl)
	desingerNum := getNumByUrl(desingerUrl)
	prooferNum := getNumByUrl(prooferUrl)

	data := &ReqData{
		Msgtype: "markdown",
		Markdown: Markdown{
			"昨日数据",
			"### 昨日数据 \r\n" +
				"- **pv**:" + pvNum + " \r\n" +
				"- **uv**:" + uvNum + " \r\n" +
				"- **newIp**:" + newIpNum + " \r\n" +
				"- **新增需求**:" + rqtNum + " \r\n" +
				"- **新增设计师**:" + desingerNum + " \r\n" +
				"- **新增打样师**：" + prooferNum + " \r\n",
		},
	}

	jsonData, err := json.Marshal(data)
	fmt.Println("string", string(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	req := bytes.NewBuffer(jsonData)

	resp, err := http.Post(postUrl, bodyType, req)
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
