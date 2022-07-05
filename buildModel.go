package main

import (
	"github.com/pulingfu/tblschema"
)

func main() {
	//简单用法
	simple := tblschema.NewTblToStructHandler()
	simple.
		SetDsn("root:root@tcp(127.0.0.1:3306)/zhifa?charset=utf8").
		SetTableName("sft_law_category").
		//默认路径为当前目录
		SetSavePath("model/SftLawCategory.go").GenerateTblStruct()
	// SetPackageInfo("plf_test_package", "", "").s
}
