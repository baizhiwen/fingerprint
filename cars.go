package main

import "fmt"

//大量运用高阶函数(回调函数)

type Any interface {
}

type Car struct {
	Model        string
	Manufacturer string //制造商
	BuildYear    int
}

type Cars []*Car

//定义一个通用函数Process()函数,它接受一个作用于每一辆car的f函数作参数
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

//在上面的基础上,实现了一个查找函数来获取子集合,并在Process()中传入一个闭包执行(这样就可以访问局部切片cars)?
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

//实现Map功能,产出除car对象以外的东西
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)

	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

//我们也可以根据参数返回不同的函数。也许我们想根据不同的厂商添加汽车到不同的集合，
//但是这（这种映射关系）可能会是会改变的。
//所以我们可以定义一个函数来产生特定的添加函数和 map 集：
func MakeSortedAppender(manufaturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufaturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

func main() {
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	//Query
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})

	fmt.Println(allCars)
	fmt.Println(allNewBMWs)

	manufaturers := []string{"Ford", "Aston Martin", "Land River", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufaturers)
	allCars.Process(sortedAppender)
	fmt.Println(sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
