package scujwc

import (
	"strings"

	"reflect"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

//ScheduleData 课程数据
type ScheduleData struct {
	Project    string `json:"project"`
	CourseID   string `json:"course_id"`
	CourseName string `json:"course_name"`
	LessonID   string `json:"lesson_id"`
	Credit     string `json:"credit"`
	CourseType string `json:"course_type"`
	ExamType   string `json:"exam_type"`
	Teacher    string `json:"teacher"`
	StudyWay   string `json:"study_way"`
	ChooseType string `json:"choose_type"`
	AllWeek    string `json:"all_week"`
	Day        string `json:"day"`
	Session    string `json:"session"`
	Campus     string `json:"campus"`
	Building   string `json:"building"`
	Classroom  string `json:"classroom"`
}

//Schedule 课程表
func (j *Jwc) Schedule() (data []ScheduleData, err error) {
	data = make([]ScheduleData, 0)
	url := DOMAIN + "/xkAction.do"
	doc, err := j.jPost(url, "actionType=6")
	if err != nil {
		return nil, err
	}

	//通过反射利用字段间的对应关系，来进行字段赋值
	schedule := &ScheduleData{}
	v := reflect.ValueOf(schedule)
	elem := v.Elem()

	doc.Find(".displayTag").Eq(1).Find("tr").Each(func(i int, sel *goquery.Selection) {
		td := sel.Find("td")
		index := 0
		k := 0

		//长度小于7说明，该课程为上一课程的不同时间段
		if td.Size() < 7 {
			k = 10
		}

		//获取数据
		for ; k < elem.NumField(); k++ {
			//跳过大纲日历
			if k == 8 {
				index++
			}
			s := td.Eq(index)
			elem.Field(k).SetString(strings.TrimSpace(s.Text()))
			index++
		}

		//只有长度大于1，才说明这一行不是标题行
		if td.Size() > 1 {
			data = append(data, *schedule)
		}
	})
	fmt.Println(data)
	return data, nil

}
