package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rex-ss/library/database/mysql"
	"time"
)

type TpshopCategory struct {
	Id         int
	CateName   string `gorm:"default('')"` //商品类型名
	Pid        int    `gorm:"default(0)"`  //父类型id
	IsShow     int    `gorm:"default(1)"`  //是否显示
	CreateTime int    `gorm:"null"`
	UpdateTime int    `gorm:"null"`
	DeleteTime int    `gorm:"null"`
}

type Goods struct {
	gorm.Model
	Name     string      `gorm:"size(20)"`  //商品名称
	Detail   string      `gorm:"size(200)"` //详细描述
	GoodsSKU []*GoodsSKU `gorm:"reverse(many)"`
}

//商品类型表
type GoodsType struct {
	gorm.Model
	Name                 string                  //种类名称
	Logo                 string                  //logo
	Image                string                  //图片
	GoodsSKU             []GoodsSKU             `gorm:"foreignkey:GoodsTypeID"`
	IndexTypeGoodsBanner []IndexTypeGoodsBanner `gorm:"reverse(many)"`
}

type GoodsSKU struct {
	Id                   int
	Goods                *Goods                  `gorm:"rel(fk)"` //商品SPU
	GoodsType            *GoodsType              `gorm:"rel(fk)"` //商品所属种类
	Name                 string                  //商品名称
	Desc                 string                  //商品简介
	Price                int                     //商品价格
	Unite                string                  //商品单位
	Image                string                  //商品图片
	Stock                int                     `gorm:"default(1)"`   //商品库存
	Sales                int                     `gorm:"default(0)"`   //商品销量
	Status               int                     `gorm:"default(1)"`   //商品状态
	Time                 time.Time               `gorm:"auto_now_add"` //添加时间
	GoodsImage           []*GoodsImage           `gorm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     `gorm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `gorm:"reverse(many)"`
	OrderGoods           []*OrderGoods           `gorm:"reverse(many)"`
}

//商品图片表
type GoodsImage struct {
	Id       int
	Image    string    //商品图片
	GoodsSKU *GoodsSKU `gorm:"rel(fk)"` //商品SKU
}

//首页轮播商品展示表
type IndexGoodsBanner struct {
	Id       int
	GoodsSKU *GoodsSKU `gorm:"rel(fk)"` //商品sku
	Image    string    //商品图片
	Index    int       `gorm:"default(0)"` //展示顺序
}

//首页分类商品展示表
type IndexTypeGoodsBanner struct {
	Id          int
	GoodsType   *GoodsType `gorm:"rel(fk)"`    //商品类型
	GoodsSKU    *GoodsSKU  `gorm:"rel(fk)"`    //商品sku
	DisplayType int        `gorm:"default(1)"` //展示类型 0代表文字，1代表图片
	Index       int        `gorm:"default(0)"` //展示顺序
}

//首页促销商品展示表
type IndexPromotionBanner struct {
	Id    int
	Name  string `gorm:"size(20)"` //活动名称
	Url   string `gorm:"size(50)"` //活动链接
	Image string //活动图片
	Index int    `gorm:"default(0)"` //展示顺序
}

///
//订单表
type OrderInfo struct {
	Id           int
	OrderId      string        `gorm:"unique"` //订单号
	UserId       int64         `json:"userId"`
	AddressId    int64         `gorm:"rel(fk)"` //地址
	PayMethod    int           //付款方式
	TotalCount   int           `gorm:"default(1)"` //商品数量
	TotalPrice   int           //商品总价
	TransitPrice int           //运费
	Orderstatus  int           `gorm:"default(1)"`   //订单状态
	TradeNo      string        `gorm:"default('')"`  //支付编号
	Time         time.Time     `gorm:"auto_now_add"` //订单时间
	OrderGoods   []*OrderGoods `gorm:"reverse(many)"`
}

//订单商品关系表
type OrderGoods struct {
	Id        int
	OrderInfo *OrderInfo `gorm:"rel(fk)"`    //订单
	GoodsSKU  *GoodsSKU  `gorm:"rel(fk)"`    //商品
	Count     int        `gorm:"default(1)"` //商品数量
	Price     int        //商品价格
	Comment   string     `gorm:"default('')"` //评论
}

func init() {
	//注册数据库
	//注册表结构

	//运行生成表
}

type IProduct interface {
	Conn() error
	Insert(product *Product) (int64, error)
	Delete(int64) bool
	Update(product *Product) error
	SelectByKey(int64) (*Product, error)
	SelectAll()
}

type PD struct {
	mysql.Dao
}

func (p *PD) Conn() {
	return
}
