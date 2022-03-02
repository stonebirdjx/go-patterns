// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: mediator_test.go
// @Date: 2022/3/2 12:43
// @Desc: 中介者模式测试方法
package mediator

import "testing"

func TestDemo(t *testing.T) {
	usernameInput := Input("username input")
	passwordInput := Input("password input")
	repeatPwdInput := Input("repeat password input")

	selection := Selection("登录")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPwdInput,
	}
	d.HandleEvent(&selection)

	regSelection := Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
