package main

import (
	"github.com/pulingfu/tblschema"
)

func main() {
	//简单用法
	simple := tblschema.NewTblToStructHandler()
	simple.
		SetDsn("root:qsgct0791@tcp(home.xzyxzm.top:3307)/zhifa?charset=utf8").
		SetTableName("sft_law_mainbody").
		//默认路径为当前目录
		SetSavePath("model/SftLawMainbody.go").GenerateTblStruct()
	// SetPackageInfo("plf_test_package", "", "").s
}
