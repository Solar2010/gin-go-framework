package todo

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"net/http"
)

type Todo struct {
	Title    string `form:"title" json:"title" binding:"required"` // 注意：gin框架数字的存储形式都是 float64
	Content  string  `form:"content" json:"content" binding:"required"`
	Mark string  `form:"mark" json:"mark" binding:"required"`
	Deadline    string  `form:"deadline" json:"deadline" binding:"required"`
	IsPrivate   string  `form:"is_private" json:"is_private" binding:"required"`
	ShareList string `form:"share_list" json:"share_list"`
	InviteList string `form:"invite_list" json:"invite_list"`
}

func (t Todo) CheckParams(context *gin.Context)  {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&t); err != nil {
		errs := gin.H{
			"tips": "todo，参数校验失败",
			"err":  err.Error(),
		}
		response.ReturnJson(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(t, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnJson(context, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+",UserUpdate表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Todo{}).Submit(extraAddBindDataContext)
	}
}