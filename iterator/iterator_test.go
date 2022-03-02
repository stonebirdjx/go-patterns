// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: iterator_test.go
// @Date: 2022/3/2 8:25
// @Desc: 迭代器模式测试方法
package iterator

import "testing"

func TestArrayInt_Iterator(t *testing.T) {
	data := ArrayInt{1, 3, 5, 7, 8}
	iterator := data.Iterator()
	// i 用于测试
	i := 0
	for iterator.HasNext() {
		if data[i] != iterator.CurrentItem() {
			t.Fatal()
		}
		iterator.Next()
		i++
	}
}
