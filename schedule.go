package scujwc

import (
	"bytes"
	"errors"
	"strings"

	"reflect"

	"fmt"

	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/jordic/goics"
)

//Schedule 课程数据
type Schedule struct {
	Project    string   `json:"project"`
	CourseID   string   `json:"course_id"`
	CourseName string   `json:"course_name"`
	LessonID   string   `json:"lesson_id"`
	Credit     string   `json:"credit"`
	CourseType string   `json:"course_type"`
	ExamType   string   `json:"exam_type"`
	Teachers   []string `json:"teacher"`
	StudyWay   string   `json:"study_way"`
	ChooseType string   `json:"choose_type"`
	AllWeek    string   `json:"all_week"`
	Day        string   `json:"day"`
	Session    string   `json:"session"`
	Campus     string   `json:"campus"`
	Building   string   `json:"building"`
	Classroom  string   `json:"classroom"`
}

const (
	JA = 1
	WJ = 2
)

//Schedule 课程表
func (j *Jwc) Schedule() (data []Schedule, err error) {
	data, err = getSchedule(*j)
	return data, err
}

func getSchedule(j Jwc) (data []Schedule, err error) {
	data = make([]Schedule, 0)
	url := DOMAIN + "/xkAction.do"
	doc, err := j.jPost(url, "actionType=6")
	if err != nil {
		return nil, err
	}

	//通过反射利用字段间的对应关系，来进行字段赋值
	schedule := &Schedule{}
	// t := reflect.TypeOf(schedule)
	v := reflect.ValueOf(schedule)
	elem := v.Elem()

	doc.Find(".displayTag").Eq(1).Find("tr").Each(func(i int, sel *goquery.Selection) {
		td := sel.Find("td")
		index := 0
		k := 0
		t := elem.Type()

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

			switch t.Field(k).Name {
			case "Teachers":
				teachers := teacherParse(s.Text())
				schedule.Teachers = teachers
			case "AllWeek":
				allWeek := weekParse(s.Text())
				schedule.AllWeek = allWeek
			default:
				elem.Field(k).SetString(strings.TrimSpace(s.Text()))
			}

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

//ScheduleIcal 生成日历
func (j *Jwc) ScheduleIcal() error {

	//基本信息设置
	c := goics.NewComponent()
	c.SetType("VCALENDAR")
	c.AddProperty("CALSCAL", "GREGORIAN")
	c.AddProperty("VERSION", "2.0")
	c.AddProperty("X-WR-CALNAME", "SCUPLUS-课表")
	c.AddProperty("X-WR-TIMEZONE", "Asia/Shanghai")
	c.AddProperty("VERSION", "2.0")
	c.AddProperty("VERSION", "2.0")
	c.AddProperty("VERSION", "2.0")
	c.AddProperty("PRODID", "-//Mohuishou//SCUPLUS//FYSCU")

	vtime := goics.NewComponent()
	vtime.SetType("VTIMEZONE")
	vtime.AddProperty("TZID", "Asia/Shanghai")
	vtime.AddProperty("X-LIC-LOCATION", "Asia/Shanghai")

	standard := goics.NewComponent()
	standard.SetType("STANDARD")
	standard.AddProperty("TZOFFSETFROM", "+0800")
	standard.AddProperty("TZOFFSETTO", "+0800")
	standard.AddProperty("TZNAME", "CST")
	standard.AddProperty("DTSTART", "19700101T000000")

	vtime.AddComponent(standard)
	c.AddComponent(vtime)

	//事件设置
	e := goics.NewComponent()
	e.SetType("VEVENT")

	ins := &EventTest{
		component: c,
	}

	w := &bytes.Buffer{}
	enc := goics.NewICalEncode(w)
	enc.Encode(ins)

	fmt.Println(w)

	return nil
}

//教师解析，返回包含每个教师名字的数组
func teacherParse(t string) (teachers []string) {
	t = strings.TrimSpace(t)
	teachers = strings.Split(t, " ")
	return teachers
}

//上课时间解析
func weekParse(w string) (allWeek string) {
	w = strings.TrimSpace(w)
	allWeek, err := dayParse(w)
	if err == nil {
		return allWeek
	}

	return "123"
}

func sessionParse(session string) (data string, err error) {
	sessions := strings.Split(session, "~")
	fmt.Println(sessions)
	if len(sessions) != 2 {
		//todo:解析
		return "", errors.New("错误")
	}
	fmt.Println(sessions)
	start, _ := strconv.Atoi(sessions[0])
	end, _ := strconv.Atoi(sessions[1])
	for i := start; i < end; i++ {
		s := strconv.Itoa(i)
		data = data + s + ","
	}
	data = data + sessions[1]
	return data, nil
}

//classTime 返回所在校区的上下课时间，以秒的形式返回
func classTime(session int, campus int) (data [2]int, err error) {

	//江安校区时刻表
	//上课时间 "0815","0910","1015","1110","1350","1445","1550","1645","1740","1920","2015","2110"
	//下课时间 "0900","0955","1100","1155","1435","1530","1635","1730","1825","2005","2100","2155"
	// classTimeJA := [2][12]int{
	// 	[12]int{},
	// 	[12]int{},
	// }

	// //望江校区时刻表
	// //上课时间 "0800","0855","1000","1055","1400","1455","1550","1655","1750","1930","2025","2120"
	// //下课时间 "0845","0940","1045","1140","1445","1540","1635","1740","1835","2015","2110","2205"
	// classTimeWJ := [2][12]int{
	// 	[12]int{},
	// 	[12]int{},
	// }

	return data, nil
}

type EventTest struct {
	component goics.Componenter
}

func (evt *EventTest) EmitICal() goics.Componenter {
	return evt.component
}
