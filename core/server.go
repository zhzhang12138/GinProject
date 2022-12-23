package core

import (
	"fmt"
	"time"

	"gin-project/global"
	"gin-project/initialize"
	"gin-project/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 初始化redis服务
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Println(`                       _oo0oo_`)
	fmt.Println(`                      o8888888o`)
	fmt.Println(`                      88" . "88`)
	fmt.Println(`                      (| -_- |)`)
	fmt.Println(`                      0\  =  /0`)
	fmt.Println("                    ___/`---'\\___")
	fmt.Println(`                  .' \\|     |// '.`)
	fmt.Println(`                 / \\|||  :  |||// \`)
	fmt.Println(`                / _||||| -:- |||||- \`)
	fmt.Println(`               |   | \\\  - /// |   |`)
	fmt.Println(`               | \_|  ''\---/''  |_/ |`)
	fmt.Println(`               \  .-\__  '-'  ___/-. /`)
	fmt.Println("             ___'. .'  /--.--\\  `. .'___\\")
	fmt.Println("          .\"\" '<  `.___\\_<|>_/___.' >' \"\".")
	fmt.Println("         | | :  `- \\`.;`\\ _ /`;.`/ - ` : | |")
	fmt.Println("         \\  \\ `_.   \\_ __\\ /__ _/   .-` /  /")
	fmt.Println("     =====`-.____`.___ \\_____/___.-`___.-'=====")
	fmt.Println("                       `=---='")
	fmt.Println(``)
	fmt.Println(`     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`)
	fmt.Println(``)
	fmt.Println(`        佛祖保佑       永不宕机     永无BUG`)
	fmt.Println(``)

	fmt.Printf(`
	欢迎使用 gin-project
	默认请求地址:http://127.0.0.1%s
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	`, address, address)

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
