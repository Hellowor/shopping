package model

import (
	"github.com/rex-ss/library/database/mysql"
	"time"
)

type TpshopCategory struct {
	Id         int
	CateName   string `orm:"default('')"` //商品类型名
	Pid        int    `orm:"default(0)"`  //父类型id
	IsShow     int    `orm:"default(1)"`  //是否显示
	CreateTime int    `orm:"null"`
	UpdateTime int    `orm:"null"`
	DeleteTime int    `orm:"null"`
}

type Goods struct {
	Id       int
	Name     string      `orm:"size(20)"`  //商品名称
	Detail   string      `orm:"size(200)"` //详细描述
	GoodsSKU []*GoodsSKU `orm:"reverse(many)"`
}

//商品类型表
type GoodsType struct {
	Id                   int
	Name                 string                  //种类名称
	Logo                 string                  //logo
	Image                string                  //图片
	GoodsSKU             []*GoodsSKU             `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

type GoodsSKU struct {
	Id                   int
	Goods                *Goods                  `orm:"rel(fk)"` //商品SPU
	GoodsType            *GoodsType              `orm:"rel(fk)"` //商品所属种类
	Name                 string                  //商品名称
	Desc                 string                  //商品简介
	Price                int                     //商品价格
	Unite                string                  //商品单位
	Image                string                  //商品图片
	Stock                int                     `orm:"default(1)"`   //商品库存
	Sales                int                     `orm:"default(0)"`   //商品销量
	Status               int                     `orm:"default(1)"`   //商品状态
	Time                 time.Time               `orm:"auto_now_add"` //添加时间
	GoodsImage           []*GoodsImage           `orm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
	OrderGoods           []*OrderGoods           `orm:"reverse(many)"`
}

//商品图片表
type GoodsImage struct {
	Id       int
	Image    string    //商品图片
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品SKU
}

//首页轮播商品展示表
type IndexGoodsBanner struct {
	Id       int
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品sku
	Image    string    //商品图片
	Index    int       `orm:"default(0)"` //展示顺序
}

//首页分类商品展示表
type IndexTypeGoodsBanner struct {
	Id          int
	GoodsType   *GoodsType `orm:"rel(fk)"`    //商品类型
	GoodsSKU    *GoodsSKU  `orm:"rel(fk)"`    //商品sku
	DisplayType int        `orm:"default(1)"` //展示类型 0代表文字，1代表图片
	Index       int        `orm:"default(0)"` //展示顺序
}

//首页促销商品展示表
type IndexPromotionBanner struct {
	Id    int
	Name  string `orm:"size(20)"` //活动名称
	Url   string `orm:"size(50)"` //活动链接
	Image string //活动图片
	Index int    `orm:"default(0)"` //展示顺序
}

//订单表
type OrderInfo struct {
	Id           int
	OrderId      string        `orm:"unique"` //订单号
	UserId       int64         `json:"userId"`
	AddressId    int64         `orm:"rel(fk)"` //地址
	PayMethod    int           //付款方式
	TotalCount   int           `orm:"default(1)"` //商品数量
	TotalPrice   int           //商品总价
	TransitPrice int           //运费
	Orderstatus  int           `orm:"default(1)"`   //订单状态
	TradeNo      string        `orm:"default('')"`  //支付编号
	Time         time.Time     `orm:"auto_now_add"` //订单时间
	OrderGoods   []*OrderGoods `orm:"reverse(many)"`
}

//订单商品关系表
type OrderGoods struct {
	Id        int
	OrderInfo *OrderInfo `orm:"rel(fk)"`    //订单
	GoodsSKU  *GoodsSKU  `orm:"rel(fk)"`    //商品
	Count     int        `orm:"default(1)"` //商品数量
	Price     int        //商品价格
	Comment   string     `orm:"default('')"` //评论
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
