package core

import (
	"github.com/GoAsyncFunc/NeXT-Server/api/panel"
	"github.com/GoAsyncFunc/NeXT-Server/conf"
)

type AddUsersParams struct {
	Tag   string
	Users []panel.UserInfo
	*panel.NodeInfo
}

type Core interface {
	Start() error
	Close() error
	AddNode(tag string, info *panel.NodeInfo, config *conf.Options) error
	DelNode(tag string) error
	AddUsers(p *AddUsersParams) (added int, err error)
	GetUserTrafficSlice(tag string, reset bool) ([]panel.UserTraffic, error)
	DelUsers(users []panel.UserInfo, tag string, info *panel.NodeInfo) error
}
