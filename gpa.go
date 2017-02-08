package scujwc

import "errors"
import "github.com/PuerkitoBio/goquery"
import "strings"

//Grades 成绩
type Grades struct {
	CourseID   string
	LessonID   string
	CourseName string
	Credit     string
	CourseType string
	Grade      string
}

//GPA 获取本学期成绩
func (j *Jwc) GPA() ([]Grades, error) {
	//登录检测
	if j.isLogin == 0 {
		err := errors.New("尚未登录，请先登录！")
		return nil, err
	}

	//获取goquery.Document 对象，以便解析需要的数据
	url := DOMAIN + "/bxqcjcxAction.do"
	doc, err := j.post(url, "")
	if err != nil {
		return nil, err
	}

	//抓取数据
	grade := make([]Grades, 10)
	var g Grades
	count := 0
	doc.Find("#user tr").Each(func(i int, s *goquery.Selection) {
		g.CourseID = strings.TrimSpace(s.Find("td").Eq(0).Text())
		if len(g.CourseID) == 0 {
			return
		}
		g.LessonID = strings.TrimSpace(s.Find("td").Eq(1).Text())
		g.CourseName = strings.TrimSpace(s.Find("td").Eq(2).Text())
		g.Credit = strings.TrimSpace(s.Find("td").Eq(4).Text())
		g.CourseType = strings.TrimSpace(s.Find("td").Eq(5).Text())
		g.Grade = strings.TrimSpace(s.Find("td").Eq(6).Text())
		if count > 9 {
			grade = append(grade, g)
		} else {
			grade[count] = g
			count++
		}
	})
	return grade, nil
}

//GPAAll 获取所有成绩
func (j *Jwc) GPAAll() ([]Grades, error) {

}
