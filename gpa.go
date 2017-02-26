package scujwc

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Grades 成绩
type Grades struct {
	CourseID   string `json:"course_id"`
	LessonID   string `json:"lesson_id"`
	CourseName string `json:"course_name"`
	Credit     string `json:"credit"`
	CourseType string `json:"course_type"`
	Grade      string `json:"grade"`
	Term       int    `json:"term"`
	TermName   string `json:"term_name"`
}

//GPA 获取本学期成绩
func (j *Jwc) GPA() ([]Grades, error) {

	//获取goquery.Document 对象，以便解析需要的数据
	url := DOMAIN + "/bxqcjcxAction.do"
	doc, err := j.jPost(url, "")
	if err != nil {
		return nil, err
	}

	//抓取数据
	grade := make([]Grades, 10)
	g := Grades{Term: 1}
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
func (j *Jwc) GPAAll() ([][]Grades, error) {

	//获取goquery.Document 对象，以便解析需要的数据
	url := DOMAIN + "/gradeLnAllAction.do"
	doc, err := j.jPost(url, "type=ln&oper=qbinfo&lnxndm")
	if err != nil {
		return nil, err
	}

	//获取学期标识
	var terms [20]string
	doc.Find("b").Each(func(i int, s *goquery.Selection) {
		terms[i] = s.Text()
	})

	grades := make([][]Grades, 10)
	doc.Find("table.displayTag").Each(func(i int, sel *goquery.Selection) {
		//获取每学期成绩
		grade := make([]Grades, 10)
		g := Grades{Term: i, TermName: terms[i]}
		count := 0
		sel.Find("#user tr").Each(func(i int, s *goquery.Selection) {
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
		grades[i] = grade
	})
	return grades, nil
}

//GPANotPass 不及格成绩
func (j *Jwc) GPANotPass() ([][]Grades, error) {

	url := DOMAIN + "/gradeLnAllAction.do"
	doc, err := j.jPost(url, "type=ln&oper=bjg")
	if err != nil {
		return nil, err
	}
	grades := make([][]Grades, 2)
	doc.Find(".displayTag").Each(func(i int, sel *goquery.Selection) {
		//获取每学期成绩
		grade := make([]Grades, 1)

		g := Grades{Term: 0}
		count := 0
		sel.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
			g.CourseID = strings.TrimSpace(s.Find("td").Eq(0).Text())
			if len(g.CourseID) == 0 {
				return
			}
			g.LessonID = strings.TrimSpace(s.Find("td").Eq(1).Text())
			g.CourseName = strings.TrimSpace(s.Find("td").Eq(2).Text())
			g.Credit = strings.TrimSpace(s.Find("td").Eq(4).Text())
			g.CourseType = strings.TrimSpace(s.Find("td").Eq(5).Text())
			g.Grade = strings.TrimSpace(s.Find("td").Eq(6).Text())
			if count > 0 {
				grade = append(grade, g)
			} else {
				grade[count] = g
				count++
			}
		})
		grades[i] = grade
	})
	return grades, nil
}
