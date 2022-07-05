package model

import (
	"github.com/astaxie/beego/orm"
	"project/utils"
	"strconv"
	"time"
)

const deleteTime = ""

//执法主体表
type SftLawMainBody struct {
	Id               int             `orm:"column:id" `                                    //是否可空:NO
	Pid              int             `orm:"column:pid" `                                   //是否可空:NO 父ID
	Name             string          `orm:"column:name" `                                  //是否可空:NO 执法主体名称
	LawCategoryCode  string          `orm:"column:law_category_code" `                     //是否可空:YES 执法主体所属领域编号
	BodyLevelCode    string          `orm:"column:body_level_code" `                       //是否可空:YES 执法主体层级编号
	BodyType         string          `orm:"column:body_type" `                             //是否可空:YES 执法主体类别
	Status           int             `orm:"column:status" `                                //是否可空:YES 是否公开显示 0公开 1不公开
	IsDrus           int             `orm:"column:is_drus" `                               //是否可空:YES 是否直属单位 0是 1不是
	ProvinceId       string          `orm:"column:province_id" `                           //是否可空:YES 所属省
	CityId           string          `orm:"column:city_id" `                               //是否可空:YES 所属市
	CountyId         string          `orm:"column:county_id" `                             //是否可空:YES 所属县区
	PreparationCount int             `orm:"column:preparation_count" `                     //是否可空:YES 编制人数
	Desc             string          `orm:"column:desc" `                                  //是否可空:YES 设定依据
	Createtime       time.Time       `orm:"column:createtime;auto_now_add;type(datetime)"` //是否可空:YES
	Updatetime       time.Time       `orm:"column:updatetime;auto_now;type(datetime)" `    //是否可空:YES
	Deletetime       time.Time       `orm:"column:deletetime" `                            //是否可空:YES 软删除时间
	CreatorId        int             `orm:"column:creator_id" `                            //是否可空:YES 创建人id
	Creator          string          `orm:"column:creator" `                               //是否可空:YES 创建人姓名
	Qrcode           string          `orm:"column:qrcode" `                                //是否可空:YES 二维码地址
	SbCode           string          `orm:"column:sb_code" `                               //是否可空:NO
	LawId            string          `orm:"column:law_id" `                                //是否可空:NO 执法类别
	MinCode          string          `orm:"column:min_code" `                              //是否可空:NO
	MaxCode          string          `orm:"column:max_code" `                              //是否可空:NO
	LawIds           string          `orm:"column:law_ids" `                               //是否可空:NO 执法类别多选id
	Type             string          `orm:"column:type" `                                  //是否可空:YES
	ExpressAddress   string          `orm:"column:express_address" `                       //是否可空:NO 快递地址
	LawCategory      *SftLawCategory `orm:"rel(fk);on_delete(set_null);null"`              //设置一对多关系
}

func (*SftLawMainBody) TableName() string {
	return "sft_law_mainbody"
}

func getTableName() string {
	body := SftLawMainBody{}
	return body.TableName()
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
