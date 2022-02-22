// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: singleton.go
// @Date: 2022/2/22 22:30
// @Desc: 单例模式
package singleton

import "sync"

type singleton struct{}

var hungerSingleton *singleton

// 饿汉式：可以将问题及早暴露
func init() {
	hungerSingleton = &singleton{}
}

func GetHungerInstance() *singleton {
	return hungerSingleton
}

// 懒汉式：将问题延时暴露
var (
	lazySingleton *singleton
	once          sync.Once
)

func GetLazyInstance() *singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton{}
		})
	}
	return lazySingleton
}
