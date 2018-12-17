package models

import "github.com/astaxie/beego/orm"

/* 用户表- hr_user */
type User struct {
	Fid      int32  `orm:"auto"`
	Ygbh     string `orm:"size(50)"`
	User     string `orm:"size(50)"`
	Password string `orm:"size(50)"`
	Lxdh     string `orm:"size(50)"`
	Zhzt     string `orm:"size(50)"`
}

//组织机构表
type Organization struct {
	Fid  int32  `orm:"auto"`
	Zzmc string `orm:"size(100)"`
	Zzbm string `orm:"size(50)"`
	Jb   string `orm:"size(50)"`
	Jdid string `orm:"size(100)"`
	Fjd  string `orm:"size(100)"`
}

/* 基础信息表 */
type Eemployerr_information struct {
	Fid     int32  `orm:"auto"`
	Ygbh    string `orm:"size(50)"`
	Xm      string `orm:"size(50)"`
	Hf      string `orm:"size(10)"`
	Jg      string `orm:"size(10)"`
	Mz      string `orm:"size(100)"`
	Zzmm    string `orm:"size(50)"`
	Sfzh    string `orm:"size(10)"`
	Dh      string `orm:"size(50)"`
	Hkszd   string `orm:"size(100)"`
	Hkxz    string `orm:"size(50)"`
	Xzz     string `orm:"size(100)"`
	Gl      int32
	Jjlxr   string `orm:"size(50)"`
	Jjlxrdh string `orm:"size(50)"`
	Rylb    string `orm:"size(50)"`
	Bgqy    string `orm:"size(200)"`
	Rzsj    string
	Sxsj    string
	Zzsj    string
	Lgsj    string
	Lzsj    string
	Htdqsj  string
	Sfffgz  string `orm:"size(10)"`
	Jbgz    float64
	Gwgz    float64
	Jxjs    float64
	Lwry    string `orm:"size(50)"`
	Txr     string `orm:"size(50)"`
	Txrq    string
}

func InitDB() {
	orm.RegisterModelWithPrefix("hr_",
		new(User),
		new(Organization),
		new(Eemployerr_information),
	)
}
