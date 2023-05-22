package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	bodyInfo := getBodyInfo()
	header := getHeaders()
	url := getUrl()
	var postHeader string
	var err error
	var response Response
	var workInfos []WorkInfo
	var pageNum = 1
	// 获取数据
	for true {
		bodyInfo["pageNum"] = pageNum
		bytesData, _ := json.Marshal(bodyInfo)
		postHeader, err = PostHeader(url, bytesData, header)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		fmt.Printf("result:%s\n", postHeader)
		err = json.Unmarshal([]byte(postHeader), &response)
		if err != nil {
			fmt.Printf("error:%s\n", err)
		}
		if response.Data.List != nil && len(response.Data.List) != 0 {
			workInfos = append(workInfos, response.Data.List...)
			if !response.Data.HasNextPage || workInfos == nil || len(workInfos) == 0 {
				log.Printf("处理完成")
				break
			}
			var workIds []int
			for _, data := range workInfos {
				infos, _ := strconv.Atoi(data.ID)
				workIds = append(workIds, infos)
			}
			fmt.Printf("workIds:%v\n", workIds)
		}
		pageNum++
	}
	WriteToFile(workInfos)

}

func WriteToFile(dataList []WorkInfo) {
	headLine := []string{"工作地点", "岗位分类", "招聘部门", "岗位名称", "实习/正式", "工作经验", "工作职责", "岗位id", "公司内部联系人", "工作地址"}
	file, err := WriteFile(headLine)
	if err != nil {
		log.Printf("error:%s\n", err)
	}
	for _, data := range dataList {
		dataLine := []string{
			data.WorkCity,
			data.PostCodeValue,
			data.DeptName,
			data.PositionName,
			data.PositionTypeName,
			data.WorkExperience,
			data.PositionDescriptions,
			data.ID,
			strings.Join(data.LeaderList, ","),
			data.WorkLocation,
		}
		file.WriteLine(dataLine)
	}
	err = file.Save("工作职位.xlsx")
	if err != nil {
		log.Printf("error:%s\n", err)
	}
}

func getUrl() string {
	url := "https://rts.biliapi.net/api/rts/recommend/summary/list"
	return url
}

func getBodyJson() []byte {
	song := getBodyInfo()
	bytesData, _ := json.Marshal(song)
	return bytesData
}

func getBodyInfo() map[string]interface{} {
	song := make(map[string]interface{})
	song["pageSize"] = 10
	song["pageNum"] = 1
	song["postCode"] = []string{"1", "11"}
	song["deptCodeList"] = []string{}
	song["workLocationList"] = []string{"上海"}
	song["recruitType"] = 0
	song["positionTypeList"] = []string{"3"}
	return song
}

func getHeaders() map[string]string {
	header := make(map[string]string)
	header["authority"] = "rts.biliapi.net"
	header["accept"] = "application/json, text/plain, */*"
	header["accept-language"] = "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"
	header["content-type"] = "application/json;charset=UTF-8"
	header["cookie"] = "username=tianyu06; _AJSESSIONID=f05773e41ca134398434d8f5f9582741"
	header["origin"] = "https://rts.biliapi.net"
	header["referer"] = "https://rts.biliapi.net/internal-recommend/social/positions?code=1,11,7&location=%E4%B8%8A%E6%B5%B7"
	header["sec-ch-ua"] = "\"Microsoft Edge\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\""
	header["sec-ch-ua-mobile"] = "?0"
	header["sec-ch-ua-platform"] = "\"macOS\""
	header["sec-fetch-dest"] = "empty"
	header["sec-fetch-mode"] = "cors"
	header["sec-fetch-site"] = "same-origin"
	header["user-agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.28"
	header["x-appkey"] = "ops.ehr-api.ats"
	header["x-csrf"] = "2022-11-02T03:29:59.274Z"
	return header
}

func PostHeader(url string, msg []byte, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(msg)))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func WriteFile(headList []string) (*Excel, error) {
	var (
		err   error
		sheet *xlsx.Sheet
		row   *xlsx.Row
	)
	file := xlsx.NewFile()
	if sheet, err = file.AddSheet("职位信息"); err != nil {
		fmt.Println(err)
		return nil, err
	}
	//设置表头
	row = sheet.AddRow()
	row.SetHeightCM(1)
	row.WriteSlice(&headList, len(headList))
	return &Excel{File: file, Sheet: sheet, LenHead: len(headList)}, nil
}

func (e *Excel) Save(fileName string) error {
	if err := e.File.Save("/Users/tianyu06/test/" + fileName); err != nil {
		return err
	}
	return nil
}

func (e *Excel) WriteLine(dataList []string) {
	row := e.Sheet.AddRow()
	row.SetHeightCM(1)
	row.WriteSlice(&dataList, len(dataList))
	e.Count++
}

type Excel struct {
	File    *xlsx.File
	Sheet   *xlsx.Sheet
	Count   int
	LenHead int
}

type Response struct {
	Code    int        `json:"code"`
	Data    Pagination `json:"data"`
	Message string     `json:"message"`
}

type Pagination struct {
	EndRow            int        `json:"endRow"`
	FirstPage         int        `json:"firstPage"`
	HasNextPage       bool       `json:"hasNextPage"`
	HasPreviousPage   bool       `json:"hasPreviousPage"`
	IsFirstPage       bool       `json:"isFirstPage"`
	IsLastPage        bool       `json:"isLastPage"`
	LastPage          int        `json:"lastPage"`
	NavigateFirstPage int        `json:"navigateFirstPage"`
	NavigateLastPage  int        `json:"navigateLastPage"`
	NavigatePages     int        `json:"navigatePages"`
	NavigatepageNums  []int      `json:"navigatepageNums"`
	NextPage          int        `json:"nextPage"`
	PageNum           int        `json:"pageNum"`
	PageSize          int        `json:"pageSize"`
	Pages             int        `json:"pages"`
	PrePage           int        `json:"prePage"`
	Size              int        `json:"size"`
	StartRow          int        `json:"startRow"`
	Total             int        `json:"total"`
	List              []WorkInfo `json:"list"`
}

type WorkInfo struct {
	AtsPositionBasicID            string      `json:"atsPositionBasicId"`
	Ctime                         string      `json:"ctime"`
	DeptCode                      string      `json:"deptCode"`
	DeptName                      string      `json:"deptName"`
	EducationRequirements         string      `json:"educationRequirements"`
	GraduationEndTime             interface{} `json:"graduationEndTime"`
	GraduationStartTime           interface{} `json:"graduationStartTime"`
	GraduationTime                string      `json:"graduationTime"`
	HotRecruit                    int         `json:"hotRecruit"`
	ID                            string      `json:"id"`
	InternshipGraduationEndTime   interface{} `json:"internshipGraduationEndTime"`
	InternshipGraduationStartTime interface{} `json:"internshipGraduationStartTime"`
	InternshipGraduationTime      string      `json:"internshipGraduationTime"`
	IsDeleted                     bool        `json:"isDeleted"`
	LeaderList                    []string    `json:"leaderList"`
	Mtime                         string      `json:"mtime"`
	PositionCloseTime             string      `json:"positionCloseTime"`
	PositionDescriptions          string      `json:"positionDescriptions"`
	PositionLeaderName            string      `json:"positionLeaderName"`
	PositionName                  string      `json:"positionName"`
	PositionSortOne               string      `json:"positionSortOne"`
	PositionSortThree             string      `json:"positionSortThree"`
	PositionSortTwo               string      `json:"positionSortTwo"`
	PositionStatus                bool        `json:"positionStatus"`
	PositionType                  string      `json:"positionType"`
	PositionTypeName              string      `json:"positionTypeName"`
	PostCode                      string      `json:"postCode"`
	PostCodeValue                 string      `json:"postCodeValue"`
	PushTime                      interface{} `json:"pushTime"`
	RecruitType                   int         `json:"recruitType"`
	WebApplyEndTime               interface{} `json:"webApplyEndTime"`
	WebApplyStartTime             interface{} `json:"webApplyStartTime"`
	WorkCity                      string      `json:"workCity"`
	WorkExperience                string      `json:"workExperience"`
	WorkLocation                  string      `json:"workLocation"`
}

func getSchoolType() []string {
	result := []string{"普通院校", "985院校", "211院校", "港澳台院校", "国外院校", "中学", "职业教育", "培训机构", "985_211_双一流院校", "985_211院校", "985_双一流院校", "211_双一流院"}
	return result
}

func getDegree() []string {
	result := []string{"高中及以下", "大专", "本科", "硕士", "博士"}
	return result
}
