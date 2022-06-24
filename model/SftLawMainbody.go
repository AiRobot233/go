package model

import (
	"github.com/astaxie/beego/orm"
	"project/utils"
	"strconv"
	"time"
)

const deleteTime = ""

func init() {
	setOrm()
}

//执法主体表
type SftLawMainBody struct {
	Id               int       `gorm:"column:id" `                //是否可空:NO
	Pid              int       `gorm:"column:pid" `               //是否可空:NO 父ID
	Name             string    `gorm:"column:name" `              //是否可空:NO 执法主体名称
	LawCategoryCode  string    `gorm:"column:law_category_code" ` //是否可空:YES 执法主体所属领域编号
	BodyLevelCode    string    `gorm:"column:body_level_code" `   //是否可空:YES 执法主体层级编号
	BodyType         string    `gorm:"column:body_type" `         //是否可空:YES 执法主体类别
	Status           int       `gorm:"column:status" `            //是否可空:YES 是否公开显示 0公开 1不公开
	IsDrus           int       `gorm:"column:is_drus" `           //是否可空:YES 是否直属单位 0是 1不是
	ProvinceId       string    `gorm:"column:province_id" `       //是否可空:YES 所属省
	CityId           string    `gorm:"column:city_id" `           //是否可空:YES 所属市
	CountyId         string    `gorm:"column:county_id" `         //是否可空:YES 所属县区
	PreparationCount int       `gorm:"column:preparation_count" ` //是否可空:YES 编制人数
	Desc             string    `gorm:"column:desc" `              //是否可空:YES 设定依据
	Createtime       time.Time `gorm:"column:createtime" `        //是否可空:YES
	Updatetime       time.Time `gorm:"column:updatetime" `        //是否可空:YES
	Deletetime       time.Time `gorm:"column:deletetime" `        //是否可空:YES 软删除时间
	CreatorId        int       `gorm:"column:creator_id" `        //是否可空:YES 创建人id
	Creator          string    `gorm:"column:creator" `           //是否可空:YES 创建人姓名
	Qrcode           string    `gorm:"column:qrcode" `            //是否可空:YES 二维码地址
	SbCode           string    `gorm:"column:sb_code" `           //是否可空:NO
	LawId            string    `gorm:"column:law_id" `            //是否可空:NO 执法类别
	MinCode          string    `gorm:"column:min_code" `          //是否可空:NO
	MaxCode          string    `gorm:"column:max_code" `          //是否可空:NO
	LawIds           string    `gorm:"column:law_ids" `           //是否可空:NO 执法类别多选id
	Type             string    `gorm:"column:type" `              //是否可空:YES
	ExpressAddress   string    `gorm:"column:express_address" `   //是否可空:NO 快递地址
}

func (*SftLawMainBody) TableName() string {
	return "sft_law_mainbody"
}

func getTableName() string {
	body := SftLawMainBody{}
	name := body.TableName()
	return name
}

//查询所有数据
func (*SftLawMainBody) Get(where [][]string, field string) []orm.Params {
	str := buildSql(getTableName(), where, field, "", deleteTime)
	var data []orm.Params
	_, _ = orm.NewOrm().Raw(str).Values(&data)
	return data
}

//查询单条数据
func (*SftLawMainBody) First(where [][]string, field string) orm.Params {
	str := buildSql(getTableName(), where, field, "1", deleteTime)
	var data []orm.Params
	num, err := orm.NewOrm().Raw(str).Values(&data)
	if err == nil && num > 0 {
		return data[0]
	}
	var arr orm.Params
	return arr
}

//查询分页数据
func (*SftLawMainBody) Paginate(where [][]string, field string, page int, size int) interface{} {
	limit := (page - 1) * size
	limitStr := strconv.Itoa(limit) + "," + strconv.Itoa(size)
	str := buildSql(getTableName(), where, field, limitStr, deleteTime)
	//查询数据
	var data []orm.Params
	_, _ = orm.NewOrm().Raw(str).Values(&data)
	//查询总数
	var count orm.ParamsList
	countStr := buildSql(getTableName(), where, "count(*)", "", deleteTime)
	_, _ = orm.NewOrm().Raw(countStr).ValuesFlat(&count)
	slice := make(map[string]interface{})
	slice["list"] = data
	slice["total"] = count[0]
	return slice
}

//修改/新增数据
func (*SftLawMainBody) Save(data SftLawMainBody) string {
	if data.Id == 0 {
		m, err := orm.NewOrm().Insert(&data)
		if err != nil {
			return err.Error()
		}
		return utils.Int64ToStr(m)
	} else {
		m, err := orm.NewOrm().Update(&data)
		if err != nil {
			return err.Error()
		}
		return utils.Int64ToStr(m)
	}
}

//真实删除数据
func (*SftLawMainBody) Delete(data SftLawMainBody) string {
	num, err := orm.NewOrm().Delete(&data)
	if err != nil {
		return err.Error()
	}
	return utils.Int64ToStr(num)
}

//软删除
func (*SftLawMainBody) Destroy(id int) string {
	if deleteTime == "" {
		m := SftLawMainBody{}
		m.Id = id
		return m.Delete(m)
	} else {
		num, _ := orm.NewOrm().QueryTable(getTableName()).Filter("id", id).Update(orm.Params{
			deleteTime: utils.NowTime(),
		})
		return utils.Int64ToStr(num)
	}
}
