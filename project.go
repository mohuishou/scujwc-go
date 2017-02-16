package scujwc

import (
	"fmt"
	"regexp"
	"strings"
)

//ProjectCourse 培养方案课程
type ProjectCourse struct {
	CourseID   string `json:"course_id"`
	CourseName string `json:"course_name"`
	Credit     string `json:"credit"`
	CourseType string `json:"course_type"`
	Grade      string `json:"grade"`
	Time       string `json:"time"`
	ChooseType string `json:"choose_type"` // 0:未选课程 1:已选及格课程 2:已选不及格课程
}

//Project 培养方案
type Project struct {
	Name         string `json:"name"`
	Credit       string `json:"credit"`
	CreditPass   string `json:"credit_pass"`
	CourseChoose string `json:"course_choose"`
	CoursePass   string `json:"course_pass"`
	CourseFail   string `json:"course_fail"`
}

//Project 获取培养方案
func (j *Jwc) Project() error {
	fmt.Println("**************************匹配开始******************************")

	//获取goquery.Document 对象，以便解析需要的数据
	url := DOMAIN + "/gradeLnAllAction.do"
	doc, err := j.jPost(url, "type=ln&oper=lnfaqk&flag=zx")
	if err != nil {
		return err
	}

	//由于该网页由js动态渲染得到,所以通过html字符串，正则匹配得到
	html, err := doc.Html()
	if err != nil {
		return err
	}

	//首先去除所有的换行符
	re, err := regexp.Compile(`\s`)
	if err != nil {
		return err
	}

	html = re.ReplaceAllString(html, "")

	//匹配到调用js函数当中所需的参数值
	re, err = regexp.Compile(`add\((.*?)\);`)
	if err != nil {
		return err
	}
	s := re.FindAllStringSubmatch(html, -1)

	//去除所有的引号
	re, err = regexp.Compile(`"|'`)
	if err != nil {
		return err
	}

	/*params
	用来保存抓取到的参数值

	参数值说明：

	*/
	params := make([][]string, 10)
	for i := range s {
		//去除引号
		tmp := re.ReplaceAllString(s[i][1], "")

		//分割字符串保存参数值
		if i > 9 {
			params = append(params, strings.Split(tmp, ","))

		} else {
			params[i] = strings.Split(tmp, ",")
		}

	}

	//TODO: 将参数值排序获得所需的数据
	fmt.Println(params)

	return nil
}
