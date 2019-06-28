package model

type Profile struct {
	Name string				//房源名称
	Rent int 				//租金
	RentalMethod string 	//租赁方式
	PropertyType string 	//房屋类型
	//Description string 		//房源描述
	ImageUrl string			//图片地址
	Floor string 			//楼层
	Oriented string 		//朝向
	Address string 			//地址
	Community string 		//小区
	PropertyCompany string 	//物业公司
	PropertyCosts string 	//物业费
}