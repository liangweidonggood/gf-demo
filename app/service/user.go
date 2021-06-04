package service

import (
	"errors"
	"fmt"
	"gf-demo/app/dao"
	"gf-demo/app/model"
)

type userService struct{}

var User = userService{}

//注册服务
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	// 账号唯一性数据检查
	if !s.CheckPassport(r.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	// 昵称唯一性数据检查
	if !s.CheckNickName(r.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(passport string) bool {
	if i, err := dao.User.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckNickName(nickname string) bool {
	if i, err := dao.User.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}
