package dbStructure

import "database/sql"

type UserBaseInfo struct {
	User_ID 			string
	User_Name 			string
	User_Avatar 		string
	User_Password		sql.NullString
	User_Phone			sql.NullString
	User_Wechat_ID		sql.NullString
	User_Wechat_Token	sql.NullString
	User_Password_Hash	sql.NullString
	User_Token			sql.NullString
	User_Status			int
}