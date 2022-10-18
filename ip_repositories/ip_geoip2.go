package main

import (
	"fmt"
	"log"
	"net"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	//db, err := geoip2.Open("GeoIP2-City.mmdb")
	//db, err := geoip2.Open("./ip_repositories/GeoLite2-City.mmdb")
	db, err := geoip2.Open("/Users/baizhiwen/Downloads/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	//ip := net.ParseIP("81.2.69.142")//英格兰
	//ip := net.ParseIP("124.202.190.252")//null
	//ip := net.ParseIP("61.148.243.236") //北京
	//ip := net.ParseIP("1.202.243.134")//北京
	//ip := net.ParseIP("218.82.164.202")//上海
	//ip := net.ParseIP("58.33.135.43")//上海 ，浦东新区
	//ip := net.ParseIP("17.218.239.99")//加利福尼亚州，city null
	//ip := net.ParseIP("1.202.243.134")//北京
	//ip := net.ParseIP("27.151.61.143")//null
	//ip := net.ParseIP("58.135.83.81")//null
	ip := net.ParseIP("59.174.98.189")//湖北省 武汉
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["zh-CN"])//区
	if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["zh-CN"])
	}
	fmt.Printf("Russian country name: %v\n", record.Country.Names["zh-CN"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
