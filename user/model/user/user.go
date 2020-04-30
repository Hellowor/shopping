package user

type UserInfo struct {
	Id        int64  `json:"id"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string
	LastName  string
}

type User struct {
	Id         int64
	Name       string       `orm:"size(40);unique"` //用户名
	Pwd        string       `orm:"size(40)"`        //密码
	Phone      string       `orm:"size(11)"`        //联系电话
	Email      string       `orm:"null"`            //邮箱
	Active     bool         `orm:"default(false)"`  //是否激活
	Addresses  []*Address   `orm:"reverse(many)"`   //联系地址
	OrderInfos []*OrderInfo `orm:"reverse(many)"`   //订单信息
}

type Address struct {
	Id         int
	Receiver   string       `orm:"size(40)"`       //收件人
	Addr       string       `orm:"size(100)"`      //详细地址
	PostCode   string       `orm:"size(6)"`        //邮编
	Phone      string       `orm:"size(11)"`       //联系电话
	IsDefault  bool         `orm:"default(false)"` //是否时默认地址
	User       *User        `orm:"rel(fk)"`        //所属用户
	OrderInfos []*OrderInfo `orm:"reverse(many)"`  //订单信息
}
