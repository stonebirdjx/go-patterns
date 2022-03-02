// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: command_test.go
// @Date: 2022/3/2 11:44
// @Desc: 命令模式测试方法
package command

func ExampleCommand() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}
