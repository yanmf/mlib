package mservice

import (
	"lib/lib/zdef"
)

// Registry 服务注册与发现接口
type Registry interface {
	// 注册服务
	Register(instance *ServiceBase) error
	// 注销服务
	Deregister(instance *ServiceBase) error
	// 服务发现
	Discover(serviceName string) ([]*ServiceBase, error)
	// 心跳检测
	Heartbeat(instance *ServiceBase) <-chan struct{}
}

type ServiceBase struct {
	ServiceConf *zdef.SessionConf
}
