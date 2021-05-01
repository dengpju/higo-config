package main

import (
	"fmt"
	"github.com/dengpju/higo-config/config"
)

func main() {

	config.Set("h0.hh1.hhh1.hhhh1", "1")
	config.Set("h0.hh2", "2")
	config.Set("h0.hh2.hhh2", "3")
	config.Set("h0.hh2", "4")
	config.Set("h0.hh2.hhh3.hhhh2", "5")
	config.Set("h0.hh2.hhh3.hhhh3", "6")
	config.Set("h0.hh2.hhh3.hhhh4.hhhhh1", "7")
	config.Set("h0.hh1", "8")
	config.Set("h0.hh3", 9)
	fmt.Println(config.All().String())

	conf := config.Get("h0").(*config.Configure)
	fmt.Println(conf)
	fmt.Println(conf.Get("h0.hh1"))
	fmt.Println(config.String("h0.hh2.hhh3.hhhh4.hhhhh1"))
	fmt.Println(config.Int("h0.hh3"))
	fmt.Println(config.Get("h0"))
	fmt.Println(config.Get("h0.hh2"))
	fmt.Println(config.All().String())
	//fmt.Println(config.Get("hh2")) //测试不存在
	conf.Set("hh3", "10")
	conf.Set("hh4.hhh4", "10")
	conf.Set("hh5", "11")

	fmt.Println(config.All().String())
	fmt.Println(conf.FirstString())
	fmt.Println(conf.IndexString(2))
	fmt.Println(conf.EndString())
	fmt.Println(conf.Remove("hh2.hhh3"))
	fmt.Println(config.All().String())



	/**

	gg := configure.New()
	gg.Set("gg", configure.New().Set("gg1", "ggg"))
	fmt.Println(gg)

	tt := configure.New()
	tt.Set("tt", gg)
	fmt.Println(tt)

	configure.Set("hh", tt)

	fmt.Println(configure.All())
	fmt.Println(configure.Get("hh"))
	fmt.Println(configure.String("hh.tt.gg.gg1"))

	 */
}
