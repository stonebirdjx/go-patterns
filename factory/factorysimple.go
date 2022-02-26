// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: factorysimple.go
// @Date: 2022/2/26 10:26
// @Desc: 简单工厂模式
package factory

import "fmt"

type selector int

var hi selector = 1
var hello selector = 2

type api interface {
	say(name string) string
}

type hiApi struct{}

func (*hiApi) say(name string) string {
	return fmt.Sprintf("hi,%s", name)
}

type helloApi struct{}

func (*helloApi) say(name string) string {
	return fmt.Sprintf("hello,%s", name)
}

func NewApi(s selector) api {
	switch s {
	case hi:
		return &hiApi{}
	case hello:
		return &helloApi{}
	default:
	}
	return nil
}
