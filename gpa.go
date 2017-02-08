package scujwc

import "errors"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "fmt"

//Grades 成绩
type Grades struct {
	courseID   string
	lessonID   string
	courseName string
	credit     string
	courseType string
	grade      string
}

//GPA 1
func (j *Jwc) GPA() (Grades, error) {
	if j.isLogin == 0 {
		err := errors.New("尚未登录，请先登录！")
		return err
	}
	url := DOMAIN + "/bxqcjcxAction.do"
	doc, err := j.post(url, "")
	if err != nil {
		return nil, err
	}
	grade := make([]Grades, 10)
	var g Grades
	count := 0
	doc.Find("#user tr").Each(func(i int, s *goquery.Selection) {
		g.courseID = strings.TrimSpace(s.Find("td").Eq(0).Text())
		if len(g.courseID) == 0 {
			return
		}
		g.lessonID = strings.TrimSpace(s.Find("td").Eq(1).Text())
		g.courseName = strings.TrimSpace(s.Find("td").Eq(2).Text())
		g.credit = strings.TrimSpace(s.Find("td").Eq(4).Text())
		g.courseType = strings.TrimSpace(s.Find("td").Eq(5).Text())
		g.grade = strings.TrimSpace(s.Find("td").Eq(6).Text())
		if count > 9 {
			grade = append(grade, g)
		} else {
			grade[count] = g
			count++
		}
	})
	fmt.Println(grade)
	return grade, nil
}

//GPAAll duo
func GPAAll() {

}
