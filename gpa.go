package scujwc

import (
	"strings"

	"reflect"

	"fmt"

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
	g := Grades{Term: 1}
	grade := g.get(j, doc.Selection)

	// //抓取数据
	// grade := make([]Grades, 0)
	// g := &Grades{Term: 1}
	// v := reflect.ValueOf(g)
	// elem := v.Elem()
	// eq := []int{0, 1, 2, 4, 5, 6}
	// doc.Find("#user tr").Each(func(i int, s *goquery.Selection) {
	// 	if i == 0 {
	// 		return
	// 	}
	// 	for k := 0; k < elem.NumField(); k++ {
	// 		if k == len(eq) {
	// 			break
	// 		}
	// 		elem.Field(k).SetString(strings.TrimSpace(s.Find("td").Eq(eq[k]).Text()))
	// 	}
	// 	grade = append(grade, *g)
	// })
	fmt.Println(grade)
	return grade, nil
}

func (g Grades) get(j *Jwc, doc *goquery.Selection) []Grades {
	//抓取数据
	grade := make([]Grades, 0)
	// g := &Grades{Term: 1}
	v := reflect.ValueOf(&g)
	elem := v.Elem()
	eq := []int{0, 1, 2, 4, 5, 6}
	doc.Find("#user tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		for k := 0; k < elem.NumField(); k++ {
			if k == len(eq) {
				break
			}
			elem.Field(k).SetString(strings.TrimSpace(s.Find("td").Eq(eq[k]).Text()))
		}
		grade = append(grade, g)
	})
	return grade
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

	grades := make([][]Grades, 0)
	doc.Find("table.displayTag").Each(func(i int, sel *goquery.Selection) {
		//获取每学期成绩
		// grade := make([]Grades, 10)
		g := Grades{Term: i, TermName: terms[i]}
		grade := g.get(j, sel)
		grades = append(grades, grade)
	})
	fmt.Println(grades)
	return grades, nil
}

//GPANotPass 不及格成绩
func (j *Jwc) GPANotPass() ([][]Grades, error) {
	url := DOMAIN + "/gradeLnAllAction.do"
	doc, err := j.jPost(url, "type=ln&oper=bjg")
	if err != nil {
		return nil, err
	}
	grades := make([][]Grades, 0)
	doc.Find(".displayTag").Each(func(i int, sel *goquery.Selection) {
		g := Grades{Term: 0}
		grade := g.get(j, sel)
		grades = append(grades, grade)
	})
	fmt.Println(grades)
	return grades, nil
}
