package dbStructure

import "database/sql"

type UserBaseInfox struct {
	Id 				string
	Name 			string
	Avator 			string
	Password		sql.NullString
	Phone			sql.NullString
	WechatID		sql.NullString
	WechatToken		sql.NullString
	PasswordHash	sql.NullString
	Token			sql.NullString
	Status			int
}