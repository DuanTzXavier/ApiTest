package readingIN

import "database/sql"

type EssaysInfo struct {
	Essay_ID 			string
	Essay_Name 			string
	Essay_Words_Count 	string
	Essay_From 			string
	Essay_Author 		string
	Essay_Creat_Time 	string
	Essay_Status		int
	Essay_Tags 			sql.NullString
}
