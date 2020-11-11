package model

import (
	"fmt"
	"goskeleton/app/utils/yml_config"
	"log"
)

type TodoModel struct {
	*BaseModel
	Id       int64  `json:"id"`
	Title string `json:"title"`
	Content     string `json:"content"`
	Mark    string `json:"mark"`
	Deadline   int64    `json:"deadline"`
	InviteList    string `json:"invite_list"`
	ShareList    string `json:"share_list"`
	IsPrivate    int `json:"is_private"`
	UserId    int64 `json:"user_id"`
}
func CreateTodoFactory(sqlType string) *TodoModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &TodoModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("usersModel工厂初始化失败")
	return nil
}

func (t *TodoModel) Add (title string, content string, mark string, deadline int64, visitList string, shareList string, isPrivate int, userId int64) bool {
	fmt.Println(title, content, mark,deadline)
	sql := "INSERT  INTO g_todo(title,content,mark, deadline, invite_list, share_list, is_private, user_id) SELECT ?,?,?,?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM g_todo WHERE  user_id=? and title=?)"
	if t.ExecuteSql(sql, title, content, mark, deadline, visitList, shareList, isPrivate, userId, userId, title) > 0 {
		return true
	}
	return false
}
