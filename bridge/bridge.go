// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: bridge.go
// @Date: 2022/2/27 10:02
// @Desc: 桥接模式：是用于把抽象化与实现化解耦。
package bridge

// IMsgSender IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件
// 可能还有 电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

// Send Send
func (s *EmailMsgSender) Send(msg string) error {
	// 这里去发送消息
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
