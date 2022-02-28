// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: observer_test.go
// @Date: 2022/2/28 20:35
// @Desc: 观察者模式测试方法
package observer

func ExampleObserver() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
