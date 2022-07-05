package model

import (
	"github.com/astaxie/beego/orm"
	"project/utils"
	"time"
)

//执法所属领域表
type SftLawCategory struct {
	Code        string            `orm:"pk;column:code" `    //是否可空:NO
	Name        string            `orm:"column:name" `       //是否可空:YES
	Deletetime  time.Time         `orm:"column:deletetime" ` //是否可空:YES 软删除时间
	LawMainBody []*SftLawMainBody `orm:"reverse(many)"`      // 设置一对多的反向关系
}

func (*SftLawCategory) TableName() string {
	return "sft_law_category"
}

//真实删除数据
func (*SftLawCategory) Delete(data SftLawCategory) string {
	num, err := orm.NewOrm().Delete(&data)
	if err != nil {
		return err.Error()
	}
	return utils.Int64ToStr(num)
}
