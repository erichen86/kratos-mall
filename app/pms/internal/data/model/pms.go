// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 相册表
type PmsAlbum struct {
	Id          int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name        string `gorm:"column:name"`
	CoverPic    string `gorm:"column:cover_pic"`
	PicCount    int    `gorm:"column:pic_count"`
	Sort        int    `gorm:"column:sort"`
	Description string `gorm:"column:description"`
}

func (m *PmsAlbum) TableName() string {
	return "pms_album"
}

// 画册图片表
type PmsAlbumPic struct {
	Id      int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	AlbumId int64  `gorm:"column:album_id"`
	Pic     string `gorm:"column:pic"`
}

func (m *PmsAlbumPic) TableName() string {
	return "pms_album_pic"
}

// 品牌表
type PmsBrand struct {
	Id                  int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name                string `gorm:"column:name"`
	FirstLetter         string `gorm:"column:first_letter"` // 首字母
	Sort                int    `gorm:"column:sort"`
	FactoryStatus       int    `gorm:"column:factory_status"` // 是否为品牌制造商：0->不是；1->是
	ShowStatus          int    `gorm:"column:show_status"`
	ProductCount        int    `gorm:"column:product_count"`         // 产品数量
	ProductCommentCount int    `gorm:"column:product_comment_count"` // 产品评论数量
	Logo                string `gorm:"column:logo"`                  // 品牌logo
	BigPic              string `gorm:"column:big_pic"`               // 专区大图
	BrandStory          string `gorm:"column:brand_story"`           // 品牌故事
}

func (m *PmsBrand) TableName() string {
	return "pms_brand"
}

// 商品评价表
type PmsComment struct {
	Id               int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId        int64     `gorm:"column:product_id"`
	MemberNickName   string    `gorm:"column:member_nick_name"`
	ProductName      string    `gorm:"column:product_name"`
	Star             int       `gorm:"column:star"`      // 评价星数：0->5
	MemberIp         string    `gorm:"column:member_ip"` // 评价的ip
	CreateTime       time.Time `gorm:"column:create_time"`
	ShowStatus       int       `gorm:"column:show_status"`
	ProductAttribute string    `gorm:"column:product_attribute"` // 购买时的商品属性
	CollectCouont    int       `gorm:"column:collect_couont"`
	ReadCount        int       `gorm:"column:read_count"`
	Content          string    `gorm:"column:content"`
	Pics             string    `gorm:"column:pics"`        // 上传图片地址，以逗号隔开
	MemberIcon       string    `gorm:"column:member_icon"` // 评论用户头像
	ReplayCount      int       `gorm:"column:replay_count"`
}

func (m *PmsComment) TableName() string {
	return "pms_comment"
}

// 产品评价回复表
type PmsCommentReplay struct {
	Id             int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	CommentId      int64     `gorm:"column:comment_id"`
	MemberNickName string    `gorm:"column:member_nick_name"`
	MemberIcon     string    `gorm:"column:member_icon"`
	Content        string    `gorm:"column:content"`
	CreateTime     time.Time `gorm:"column:create_time"`
	Type           int       `gorm:"column:type"` // 评论人员类型；0->会员；1->管理员
}

func (m *PmsCommentReplay) TableName() string {
	return "pms_comment_replay"
}

// 运费模版
type PmsFeightTemplate struct {
	Id             int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name           string `gorm:"column:name"`
	ChargeType     int    `gorm:"column:charge_type"`  // 计费类型:0->按重量；1->按件数
	FirstWeight    string `gorm:"column:first_weight"` // 首重kg
	FirstFee       string `gorm:"column:first_fee"`    // 首费（元）
	ContinueWeight string `gorm:"column:continue_weight"`
	ContinmeFee    string `gorm:"column:continme_fee"`
	Dest           string `gorm:"column:dest"` // 目的地（省、市）
}

func (m *PmsFeightTemplate) TableName() string {
	return "pms_feight_template"
}

// 商品会员价格表
type PmsMemberPrice struct {
	Id              int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId       int64  `gorm:"column:product_id"`
	MemberLevelId   int64  `gorm:"column:member_level_id"`
	MemberPrice     string `gorm:"column:member_price"` // 会员价格
	MemberLevelName string `gorm:"column:member_level_name"`
}

func (m *PmsMemberPrice) TableName() string {
	return "pms_member_price"
}

// 商品信息
type PmsProduct struct {
	Id                         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	BrandId                    int64     `gorm:"column:brand_id"`
	ProductCategoryId          int64     `gorm:"column:product_category_id"`
	FeightTemplateId           int64     `gorm:"column:feight_template_id"`
	ProductAttributeCategoryId int64     `gorm:"column:product_attribute_category_id"`
	Name                       string    `gorm:"column:name;NOT NULL"`
	Pic                        string    `gorm:"column:pic"`
	ProductSn                  string    `gorm:"column:product_sn;NOT NULL"` // 货号
	DeleteStatus               int       `gorm:"column:delete_status"`       // 删除状态：0->未删除；1->已删除
	PublishStatus              int       `gorm:"column:publish_status"`      // 上架状态：0->下架；1->上架
	NewStatus                  int       `gorm:"column:new_status"`          // 新品状态:0->不是新品；1->新品
	RecommandStatus            int       `gorm:"column:recommand_status"`    // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int       `gorm:"column:verify_status"`       // 审核状态：0->未审核；1->审核通过
	Sort                       int       `gorm:"column:sort"`                // 排序
	Sale                       int       `gorm:"column:sale"`                // 销量
	Price                      string    `gorm:"column:price"`
	PromotionPrice             string    `gorm:"column:promotion_price"`       // 促销价格
	GiftGrowth                 int       `gorm:"column:gift_growth;default:0"` // 赠送的成长值
	GiftPoint                  int       `gorm:"column:gift_point;default:0"`  // 赠送的积分
	UsePointLimit              int       `gorm:"column:use_point_limit"`       // 限制使用的积分数
	SubTitle                   string    `gorm:"column:sub_title"`             // 副标题
	Description                string    `gorm:"column:description"`           // 商品描述
	OriginalPrice              string    `gorm:"column:original_price"`        // 市场价
	Stock                      int       `gorm:"column:stock"`                 // 库存
	LowStock                   int       `gorm:"column:low_stock"`             // 库存预警值
	Unit                       string    `gorm:"column:unit"`                  // 单位
	Weight                     string    `gorm:"column:weight"`                // 商品重量，默认为克
	PreviewStatus              int       `gorm:"column:preview_status"`        // 是否为预告商品：0->不是；1->是
	ServiceIds                 string    `gorm:"column:service_ids"`           // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string    `gorm:"column:keywords"`
	Note                       string    `gorm:"column:note"`
	AlbumPics                  string    `gorm:"column:album_pics"` // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string    `gorm:"column:detail_title"`
	DetailDesc                 string    `gorm:"column:detail_desc"`
	DetailHtml                 string    `gorm:"column:detail_html"`           // 产品详情网页内容
	DetailMobileHtml           string    `gorm:"column:detail_mobile_html"`    // 移动端网页详情
	PromotionStartTime         time.Time `gorm:"column:promotion_start_time"`  // 促销开始时间
	PromotionEndTime           time.Time `gorm:"column:promotion_end_time"`    // 促销结束时间
	PromotionPerLimit          int       `gorm:"column:promotion_per_limit"`   // 活动限购数量
	PromotionType              int       `gorm:"column:promotion_type"`        // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string    `gorm:"column:brand_name"`            // 品牌名称
	ProductCategoryName        string    `gorm:"column:product_category_name"` // 商品分类名称
}

func (m *PmsProduct) TableName() string {
	return "pms_product"
}

// 商品属性参数表
type PmsProductAttribute struct {
	Id                         int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductAttributeCategoryId int64  `gorm:"column:product_attribute_category_id"`
	Name                       string `gorm:"column:name"`
	SelectType                 int    `gorm:"column:select_type"`     // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int    `gorm:"column:input_type"`      // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string `gorm:"column:input_list"`      // 可选值列表，以逗号隔开
	Sort                       int    `gorm:"column:sort"`            // 排序字段：最高的可以单独上传图片
	FilterType                 int    `gorm:"column:filter_type"`     // 分类筛选样式：1->普通；1->颜色
	SearchType                 int    `gorm:"column:search_type"`     // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int    `gorm:"column:related_status"`  // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              int    `gorm:"column:hand_add_status"` // 是否支持手动新增；0->不支持；1->支持
	Type                       int    `gorm:"column:type"`            // 属性的类型；0->规格；1->参数
}

func (m *PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}

// 产品属性分类表
type PmsProductAttributeCategory struct {
	Id             int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name           string `gorm:"column:name"`
	AttributeCount int    `gorm:"column:attribute_count;default:0"` // 属性数量
	ParamCount     int    `gorm:"column:param_count;default:0"`     // 参数数量
}

func (m *PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

// 存储产品参数信息的表
type PmsProductAttributeValue struct {
	Id                 int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId          int64  `gorm:"column:product_id"`
	ProductAttributeId int64  `gorm:"column:product_attribute_id"`
	Value              string `gorm:"column:value"` // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}

func (m *PmsProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}

// 产品分类
type PmsProductCategory struct {
	Id           int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ParentId     int64  `gorm:"column:parent_id"` // 上机分类的编号：0表示一级分类
	Name         string `gorm:"column:name"`
	Level        int    `gorm:"column:level"` // 分类级别：0->1级；1->2级
	ProductCount int    `gorm:"column:product_count"`
	ProductUnit  string `gorm:"column:product_unit"`
	NavStatus    int    `gorm:"column:nav_status"`  // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int    `gorm:"column:show_status"` // 显示状态：0->不显示；1->显示
	Sort         int    `gorm:"column:sort"`
	Icon         string `gorm:"column:icon"` // 图标
	Keywords     string `gorm:"column:keywords"`
	Description  string `gorm:"column:description"` // 描述
}

func (m *PmsProductCategory) TableName() string {
	return "pms_product_category"
}

// 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelation struct {
	Id                 int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductCategoryId  int64 `gorm:"column:product_category_id"`
	ProductAttributeId int64 `gorm:"column:product_attribute_id"`
}

func (m *PmsProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}

// 产品满减表(只针对同商品)
type PmsProductFullReduction struct {
	Id          int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId   int64  `gorm:"column:product_id"`
	FullPrice   string `gorm:"column:full_price"`
	ReducePrice string `gorm:"column:reduce_price"`
}

func (m *PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}

// 产品阶梯价格表(只针对同商品)
type PmsProductLadder struct {
	Id        int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId int64  `gorm:"column:product_id"`
	Count     int    `gorm:"column:count"`    // 满足的商品数量
	Discount  string `gorm:"column:discount"` // 折扣
	Price     string `gorm:"column:price"`    // 折后价格
}

func (m *PmsProductLadder) TableName() string {
	return "pms_product_ladder"
}

type PmsProductOperateLog struct {
	Id               int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId        int64     `gorm:"column:product_id"`
	PriceOld         string    `gorm:"column:price_old"`
	PriceNew         string    `gorm:"column:price_new"`
	SalePriceOld     string    `gorm:"column:sale_price_old"`
	SalePriceNew     string    `gorm:"column:sale_price_new"`
	GiftPointOld     int       `gorm:"column:gift_point_old"` // 赠送的积分
	GiftPointNew     int       `gorm:"column:gift_point_new"`
	UsePointLimitOld int       `gorm:"column:use_point_limit_old"`
	UsePointLimitNew int       `gorm:"column:use_point_limit_new"`
	OperateMan       string    `gorm:"column:operate_man"` // 操作人
	CreateTime       time.Time `gorm:"column:create_time"`
}

func (m *PmsProductOperateLog) TableName() string {
	return "pms_product_operate_log"
}

// 商品审核记录
type PmsProductVertifyRecord struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId  int64     `gorm:"column:product_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	VertifyMan string    `gorm:"column:vertify_man"` // 审核人
	Status     int       `gorm:"column:status"`
	Detail     string    `gorm:"column:detail"` // 反馈详情
}

func (m *PmsProductVertifyRecord) TableName() string {
	return "pms_product_vertify_record"
}

// sku的库存
type PmsSkuStock struct {
	Id             int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId      int64  `gorm:"column:product_id"`
	SkuCode        string `gorm:"column:sku_code;NOT NULL"` // sku编码
	Price          string `gorm:"column:price"`
	Stock          int    `gorm:"column:stock;default:0"`      // 库存
	LowStock       int    `gorm:"column:low_stock"`            // 预警库存
	Pic            string `gorm:"column:pic"`                  // 展示图片
	Sale           int    `gorm:"column:sale"`                 // 销量
	PromotionPrice string `gorm:"column:promotion_price"`      // 单品促销价格
	LockStock      int    `gorm:"column:lock_stock;default:0"` // 锁定库存
	SpData         string `gorm:"column:sp_data"`              // 商品销售属性，json格式
}

func (m *PmsSkuStock) TableName() string {
	return "pms_sku_stock"
}
