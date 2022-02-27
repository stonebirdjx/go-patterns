// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: bridge_test.go
// @Date: 2022/2/27 10:05
// @Desc: 桥接模式测试方法
package bridge

import "testing"

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")
	if err != nil {
		t.Fatal()
	}
}
