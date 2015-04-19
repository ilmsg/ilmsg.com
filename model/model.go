package model

import(
	"time"
	"html/template"
	"gopkg.in/mgo.v2/bson"
)

type Contact struct {
    Id          int `form:"Id"`
    Email       string `form:"email"`
    Content     string `form:"content"`
    Date        string `form:"date"`
}

type Article struct {
    Id          int				`form:"Id"`
    Title       string			`form:"title"`
    Content     template.HTML	`form:"content"`
    Date        string			`form:"date"`
    Category    string			`form:"category"`
}

type Project struct {
    Id				bson.ObjectId	`form:"id" json:"id" bson:"_id"`
    Title			string			`form:"title"`
    Description		string			`form:"description"`
	Image			string			`form:"image"`
	Order			string			`form:"order"`
	Show			string			`form:"show"`
	Content			template.HTML	`form:"content"`
	Tag				string 			`form:"tag"`
	Demo			string 			`form:"demo"`
    Date			time.Time		`form:"date"`
    Category		string			`form:"category"`
}

type Og struct {
	Title			string		`form:"title"`
	Description     string		`form:"description"`
	Url				string		`form:"url"`
	Image			string		`form:"image"`
	PublishedTime	time.Time	`form:"published_time"`
	ModifiedTime	time.Time	`form:"modified_time"`
}
