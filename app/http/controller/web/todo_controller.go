package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/service/todo"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {

}

func (t *Todo) Submit(context *gin.Context){
	token := context.GetHeader("Authorization")
	tokenData := strings.Split(token, " ")
	userInfo, isEffective := userstoken.CreateUserFactory().ParseToken(tokenData[1])
	if isEffective == consts.JwtTokenOK {
		userId := userInfo.UserId
		title := context.GetString(consts.ValidatorPrefix + "title")
		content := context.GetString(consts.ValidatorPrefix + "content")
		mark := context.GetString(consts.ValidatorPrefix + "mark")
		deadline := context.GetString(consts.ValidatorPrefix + "deadline")
		inviteList := context.GetString(consts.ValidatorPrefix + "invite_list")
		shareList := context.GetString(consts.ValidatorPrefix + "share_list")
		isPrivate := context.GetInt(consts.ValidatorPrefix + "is_private")
		fmt.Println(shareList)
		time, _ := strconv.ParseInt(deadline, 10, 64)
		if todo.CreateTodoFactory().Submit(title, content, mark, time, inviteList, shareList, isPrivate, userId) {
			response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
		} else {
			response.ReturnJson(context, http.StatusOK, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
		}
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdRefreshTokenFailCode, consts.CurdRefreshTokenFailMsg, "")
	}
}
