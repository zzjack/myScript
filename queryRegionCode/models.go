package main

type Database struct {
	User string	`json:"user"`
	Port string	`json:"port"`
	Host string	`json:"host"`
	Password string	`json:"password"`
	Name string	`json:"name"`
}

type GeoCode struct {
	DistrictFullCode string `gorm:"column:district_full_code"`
	CountryCode string `gorm:"column:country_code"`
	CountryName string `gorm:"column:country_name"`
	ProvinceCode string `gorm:"column:province_code"`
	ProvinceName string `gorm:"column:province_name"`
	CityCode string `gorm:"column:city_code"`
	CityFullCode string `gorm:"column:city_full_code"`
	CityName string `gorm:"column:city_name"`
	DistrictCode string `gorm:"district_code"`
	DistrictName string `gorm:"district_name"`
}

func (g GeoCode)TableName() string{
	return "cfd_ref_geocode"
}

type BsdHighDangerArea struct {
	ProvinceCode string  `gorm:"column:province_code"`
	ProvinceName string `gorm:"column:province_name"`
	CityFullCode string `gorm:"column:city_full_code"`
	CityName string `gorm:"column:city_name"`
	DistrictCode string `gorm:"column:district_code"`
	DistrictName string `gorm:"column:district_name"`
	Type string `gorm:"column:type"`
}

func (b BsdHighDangerArea)TableName()string{
	return "bsd_high_danger_area"
}


