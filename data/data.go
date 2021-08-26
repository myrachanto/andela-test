package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	// "strconv"

	"github.com/myrachanto/testgo/httperrors"
)


type Post struct {
	UserId int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}
type Pos struct {
	PostId int    `json:"postId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Body   string `json:"body,omitempty"`
}
type Results struct {
	UserId   int    `json:"userId,omitempty"`
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	Comments []string `json:"comments,omitempty"`

}
func Getblog() *httperrors.Httperror {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return httperrors.BadNotfound("not found")
	}
	respo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httperrors.Badrequest("Seomething went wrong")
	}
	// fmt.Println(">>>>>>>>>>>>",string(respo))
	// var resppost Post
	// json.Unmarshal(respo, &resppost)
	// fmt.Println(">>>>>>>", resppost)
	// // 

	var posti []map[string]interface{}
	err = json.Unmarshal([]byte(respo), &posti)
	if err != nil {
		return httperrors.Badrequest("Seomething went wrong UNmasharling!!!")
	}
	ps := Post{}
	pss := []Post{}
	for _, val := range posti {
		ps.Id  = int(val["id"].(float64))
		ps.UserId  = int(val["userId"].(float64))
		ps.Title = fmt.Sprintf("%s", val["title"])
		ps.Body = fmt.Sprintf("%s", val["body"])
		pss = append(pss, ps)
	}
	fmt.Println(">>>>>", pss)

	postcomment, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		return httperrors.BadNotfound("not found")
	}
	
	posts, err := ioutil.ReadAll(postcomment.Body)
	if err != nil {
		return httperrors.Badrequest("Seomething went wrong")
	}
	// fmt.Println(">>>>>>>>>>>post>", string(posts))
	var posti1 []map[string]interface{}
	err = json.Unmarshal([]byte(posts), &posti1)
	if err != nil {
		return httperrors.Badrequest("Seomething went wrong UNmasharling!!!")
	}
	
	// fmt.Println(">>>>>vvvvv", posti1)
	ps1 := Pos{}
	pss1 := []Pos{}
	for _, val := range posti1 {
		ps1.PostId  = int(val["postId"].(float64))
		ps1.Id  = int(val["id"].(float64))
		ps1.Name = fmt.Sprintf("%s", val["name"])
		ps1.Email = fmt.Sprintf("%s", val["email"])
		ps1.Body = fmt.Sprintf("%s", val["body"])
		pss1 = append(pss1, ps1)
	}
	// fmt.Println(">>>>>", pss1)
	res := mappost(pss, pss1)
	
	// fmt.Println(">>>>>", res)
    tocsv(res)

	return nil
}
func mappost( owner []Post, comments []Pos) []Results{
	results := []Results{}
	var el []string
	res := Results{}
	 for _, v := range owner {
		 for _, l := range comments {
			 if v.Id == l.PostId {
				res.Comments = append(el, l.Body)
			 }			 
		 }		 
		 res.UserId = v.UserId
		 res.Id = v.Id
		 res.Title = v.Title
		 res.Body = v.Body
		 results = append(results, res)

	 }
	 return results
}
func tocsv(data []Results){
	csvfile, err := os.Create("test.csv")
	if err != nil{
		log.Fatal("error creating csv file")
	}
	// m := make(map[int]Results)
	defer csvfile.Close()
	csvfilewriter := csv.NewWriter(csvfile)
	 defer csvfilewriter.Flush()
	 headers := []string {
		 "userid", "id", "title","body","comments",
	 }
	//  fmt.Println("data", data)
	 csvfilewriter.Write(headers)
	 for _, v := range data {
		 r := make([]string , 0, 1+len(headers))
		 userid := strconv.Itoa(v.UserId)
		 id := strconv.Itoa(v.Id)
		 title := v.Title
		 body := v.Body
		 commetns := v.Comments
		 comes := ""
		 for _,v := range commetns{
			 comes += v +"|"
		 }
		//   fmt.Println("rrrrrrrrrrrrrrrrrrr", id,title,body,commetns)
		 r = append(r, userid, id, title, body, comes)
		 fmt.Println("rrrrrrrrrrrrrrrrrrr", r)
		 csvfilewriter.Write(r)
	 }
}