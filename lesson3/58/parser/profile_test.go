package parser

import (
	"testing"
	"io/ioutil"
	"Tiny-Go-Crawler/Crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "可短租 地铁3号线 整租精装一室开间 无中介 家电全新")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 elements; but was %v", len(result.Items))
	}

	profile := result.Items[0].(model.Profile)

	//fmt.Println(profile)
	expected := model.Profile{
		Name: "可短租 地铁3号线 整租精装一室开间 无中介 家电全新",
		Rent: 2339,
		RentalMethod: "整租",
		PropertyType: "1室1厅1卫46平精装修",
		ImageUrl: "https://pic4.58cdn.com.cn/anjuke_58/964fa985631deaf60435ae7015cc1792?w=696&h=522&crop=1",
		Floor: "低层 / 22层",
		Oriented: "南",
		Address: "开华道3号",
		Community: "海泰国际公寓",
		PropertyCompany: "天津新技术产业园区物业管理有限公司",
		PropertyCosts: "0.00元/平米/月",
	}

	if profile.Name != expected.Name {
		t.Errorf("expected: %s; but was %s,",  profile.Name, expected.Name)
	}
	if profile.Rent != expected.Rent {
		t.Errorf("expected: %s; but was %s,",  profile.Rent, expected.Rent)
	}
	if profile.RentalMethod != expected.RentalMethod {
		t.Errorf("expected: %s; but was %s,",  profile.RentalMethod, expected.RentalMethod)
	}
	if profile.PropertyType != expected.PropertyType {
		t.Errorf("expected: %s; but was %s,",  profile.PropertyType, expected.PropertyType)
	}
	if profile.ImageUrl != expected.ImageUrl {
		t.Errorf("expected: %s; but was %s,",  profile.ImageUrl, expected.ImageUrl)
	}
	if profile.Floor != expected.Floor {
		t.Errorf("expected: %s; but was %s,",  profile.Floor, expected.Floor)
	}
	if profile.Oriented != expected.Oriented {
		t.Errorf("expected: %s; but was %s,",  profile.Oriented, expected.Oriented)
	}
	if profile.Address != expected.Address {
		t.Errorf("expected: %s; but was %s,",  profile.Address, expected.Address)
	}
	if profile.Community != expected.Community {
		t.Errorf("expected: %s; but was %s,",  profile.Community, expected.Community)
	}
	if profile.PropertyCompany != expected.PropertyCompany {
		t.Errorf("expected: %s; but was %s,",  profile.PropertyCompany, expected.PropertyCompany)
	}
	if profile.PropertyCosts != expected.PropertyCosts {
		t.Errorf("expected: %s; but was %s,",  profile.PropertyCosts, expected.PropertyCosts)
	}
}