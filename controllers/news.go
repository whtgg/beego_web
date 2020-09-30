package controllers

import (
	"official/models"
	"strings"
	"fmt"
)

type NewsController struct {
	BaseController
}

func (self *NewsController) List() {
	self.Data["pageTitle"] = "新闻资讯"
	self.display()
}

func (self *NewsController) Add() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "新闻添加"
	self.display()
}

func (self *NewsController) Save() {
	News := new(models.News)
	News.Title = strings.TrimSpace(self.GetString("title"))
	News.SubTitle = strings.TrimSpace(self.GetString("subtitle"))
	News.Content = strings.TrimSpace(self.GetString("content"))
	News.Status, _ = self.GetInt8("status")
	News.Weight, _ = self.GetInt("weight")
	//categoryId, _ := self.GetInt("category")
	//News.Category = &models.Category{Id: categoryId}
	News.Genre, _ = self.GetInt8("genre")
	News.Tags,_ = self.GetInt("tags")
	News.Publish = self.userName
	News.Cover = self.GetString("cover")

	keyId, _ := self.GetInt("id")
	if keyId == 0 {
		if _, err := models.NewsAdd(News); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		res, err := models.NewsGetById(keyId)
		if res != nil {
			res.Title = News.Title
			res.SubTitle = News.SubTitle
			res.Content = News.Content
			res.Status = News.Status
			res.Weight = News.Weight
			res.Publish = News.Publish
			res.Genre = News.Genre
			res.Cover = News.Cover
			res.Tags = News.Tags
			if err := res.NewsUpdate(); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *NewsController) Edit() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "新闻编辑"

	keyId, err := self.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}

	news, _ := models.NewsGetById(keyId)
	list := make(map[string]interface{})
	list["id"] = news.Id
	list["title"] = news.Title
	list["subTitle"] = news.SubTitle
	list["status"] = news.Status
	list["content"] = news.Content
	list["weight"] = news.Weight
	list["genre"] = news.Genre
	list["cover"] = news.Cover
	list["tags"] = news.Tags
	self.Data["news"] = list
	self.display()
}

func (self *NewsController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = self.pageSize
	}

	title := strings.TrimSpace(self.GetString("title"))

	filters := make([]interface{}, 0)
	filters = append(filters, "status__lt", 9)

	if title != "" {
		filters = append(filters, "title__contains", title)
	}

	result, count := models.NewsGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["status"] = v.Status
		row["publish"] = v.Publish
		row["createTime"] = v.CreateTime.String()[:19]
		row["genre"] = v.Genre
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *NewsController) Delete() {
	keyId, _ := self.GetInt("id")
	res, _ := models.NewsGetById(keyId)
	res.Status = 9
	err := res.NewsUpdate()
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}
