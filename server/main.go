package main

import (
	"encoding/csv"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"os"
)

func main() {
	err, issue := jira_client()

	write_csv(err, issue)

}

func write_csv(err error, issue []jira.Issue) {
	//csv 文件相关
	f, err := os.Create("test.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f) //创建一个新的写入文件流
	//创建表头
	data := [][]string{
		{"key","描述", "经办人", "优先级"},
	}
	w.WriteAll(data) //写入数据

	//创建内容写入信息
	w1 := csv.NewWriter(f) //创建一个新的写入文件流
	write_data := [][]string{}
	for _, v := range issue {

		issue_data := [][]string{
			{
				v.Key,
				v.Fields.Summary,
				v.Fields.Assignee.DisplayName,
				v.Fields.Priority.Name,
			},
		}
		write_data = append(write_data, issue_data...)
	}
	w1.WriteAll(write_data)
	w.Flush()
}

func jira_client() (error, []jira.Issue) {
	tp := jira.BearerAuthTransport{Token: "xxxxxxxxxxxx"}

	client, _ := jira.NewClient(tp.Client(), "https://jira.expmle.net")
	req, err := client.NewRequest("GET", "rest/api/2/search/", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)
	jql := "project = xxxxx AND issuetype = Bug AND text ~ 【一体机】 AND created >= 2023-09-10 AND created <= 2023-09-16  ORDER BY priority DESC, updated DESC"
	issue, resq, err := client.Issue.Search(jql, nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resq.Total)
	for _, v := range issue {
		fmt.Println(v.Key,v.Fields.Summary, v.Fields.Assignee.DisplayName, v.Fields.Priority.Name)

	}
	return err, issue
}
