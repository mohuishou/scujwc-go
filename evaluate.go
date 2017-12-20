package scujwc

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Evaluate 评教信息列表
type Evaluate struct {
	EvaluateID   string `json:"evaluate_id"`
	EvaluateType string `json:"evaluate_type"`
	TeacherName  string `json:"teacher_name"`
	TeacherID    string `json:"teacher_id"`
	TeacherType  string `json:"teacher_type"`
	CourseName   string `json:"course_name"`
	CourseID     string `json:"course_id"`
	Status       int    `json:"status"`
	Star         string `json:"star"`    //平均分
	Comment      string `json:"comment"` //评价内容
}

func (e *Evaluate) getParams() string {
	params := url.Values{}
	params.Set("wjbm", e.EvaluateID)
	params.Set("wjmc", Utf8ToGbk(e.EvaluateType))
	params.Set("bpr", e.TeacherID)
	params.Set("bprm", Utf8ToGbk(e.TeacherName))
	params.Set("pgnr", e.CourseID)
	params.Set("pgnrm", Utf8ToGbk(e.CourseName))
	params.Set("oper", "wjShow")
	params.Set("pageSize", "20")
	params.Set("page", "1")
	params.Set("currentPage", "1")
	return params.Encode()
}

//EvaluateListURL 列表页面链接
const EvaluateListURL = DOMAIN + "/jxpgXsAction.do?oper=listWj&pageSize=100"

//EvaluateURL 评教页面链接
const EvaluateURL = DOMAIN + "/jxpgXsAction.do"

func (j *Jwc) getEvaList() ([]Evaluate, error) {
	doc, err := j.get(EvaluateListURL, "")
	evaluateList := make([]Evaluate, 0)
	if err != nil {
		return nil, err
	}
	doc.Find("#user tbody td img").Each(func(i int, selection *goquery.Selection) {
		eva := &Evaluate{}
		val, exist := selection.Attr("name")
		if !exist {
			return
		}
		if data := strings.Split(val, "#@"); len(data) == 6 {
			eva.EvaluateID = data[0]
			eva.TeacherID = data[1]
			eva.TeacherName = data[2]
			eva.EvaluateType = data[3]
			eva.CourseName = data[4]
			eva.CourseID = data[5]
			eva.getEvaInfo(j)
			evaluateList = append(evaluateList, *eva)
		}
	})
	return evaluateList, nil
}

func (e *Evaluate) getEvaInfo(j *Jwc) error {
	params := e.getParams()
	log.Println(params)
	doc, err := j.jPost(EvaluateURL, params)
	if err != nil {
		return err
	}
	log.Println(doc.Html())
	return nil
}

// 获取分数，返回一个map 0000000038 => 2
func getEvaStar(doc *goquery.Document) (map[string]int, error) {
	stars := make(map[string]int)
	doc.Find("#tblView table tbody td input[type=\"radio\"]").Each(func(i int, s *goquery.Selection) {
		name, exist := s.Attr("name")
		if !exist {
			return
		}
		value, exist := s.Attr("value")
		if !exist {
			return
		}
		tmp := strings.Split(value, "_")
		if len(tmp) != 2 {
			return
		}

	})
	return stars, nil
}

// GetEvaluate 获取评教数据
func (j *Jwc) GetEvaluate() {

}
