package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 + iota //管理员
	PermissionUser                        //普通用户
	PermissionVisitor                     //游客
	PermissionDisableUser                 //被禁言用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言用户"
	default:
		str = "其他"
	}
	return str
}
