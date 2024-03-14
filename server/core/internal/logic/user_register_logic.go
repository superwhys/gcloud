package logic

import (
	"context"
	"gcloud/core/define"
	"gcloud/core/helper"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	resp = new(types.UserRegisterReply)
	if len(req.Name) < 6 {
		resp.Msg = "用户名长度不能小于6位"
		return
	}
	if len(req.Password) < 6 {
		resp.Msg = "密码长度不能小于6位"
		return
	}

	// 判断code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		logx.Errorf("query code from rdb error: %v", err)
		resp.Msg = "无效验证码"
		return
	}
	if code != req.Code {
		resp.Msg = "验证码错误"
		return
	}

	// 判断用户名是否已存在
	var count int64
	err = l.svcCtx.Engine.
		Table("user_basic").
		Where("name = ?", req.Name).
		Count(&count).Error
	if err != nil {
		logx.Errorf("check username from db error: %v", err)
		resp.Msg = "出错了"
		return
	}
	if count > 0 {
		resp.Msg = "用户名已存在"
		return
	}

	// 入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Email:    req.Email,
		Avatar:   define.AvatarBaseUrl + helper.Random() + ".png",
		Password: helper.Md5(req.Password),
		Capacity: define.UserRepositoryMinSize,
	}
	// fix: 需指定添加字段 (Select())，不推荐使用 Omit()
	err = l.svcCtx.Engine.
		Select("identity", "name", "email", "password", "capacity", "created_at", "updated_at").
		Create(user).Error
	if err != nil {
		logx.Errorf("add user to db error: %v", err)
		resp.Msg = "出错了"
		return
	}
	resp.Msg = "注册成功"
	resp.Code = 200
	return
}

func Random() {
	panic("unimplemented")
}
