// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: facade_test.go
// @Date: 2022/2/27 12:28
// @Desc: 外观模式测试方法
package facade

import (
	"reflect"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	service := UserService{}
	user, err := service.Login(13001010101, 1234)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(user, &User{Name: "test login"}) {
		t.Fatal()
	}

}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(13001010101, 1234)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(user, &User{Name: "test login"}) {
		t.Fatal()
	}
}

func TestUserService_Register(t *testing.T) {
	service := UserService{}
	user, err := service.Register(13001010101, 1234)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(user, &User{Name: "test register"}) {
		t.Fatal()
	}
}
