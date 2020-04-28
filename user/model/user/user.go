package user

type UserInfo struct {
	Id        int64  `json:"id"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string
	LastName  string
}
