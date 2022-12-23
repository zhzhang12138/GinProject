package core

import (
	"flag"
	"fmt"
	"gin-project/core/internal"
	"gin-project/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		// 解析命令行以获取-c指定要使用的配置文件的标志
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		// 如果-c未设置该标志，该函数将检查常量指定的环境变量internal.ConfigEnv
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				// 如果未设置此环境变量，该函数会查看当前的 Gin 模式（由gin.Mode()函数指定）并config根据该模式设置为默认配置文件路径
				switch gin.Mode() {
				case gin.DebugMode:
					// 调试模式。在这种模式下，Gin 会打印许多调试信息，并且可以在浏览器中看到详细的错误信息。这对于开发和调试项目很有帮助，但是在生产环境中并不适用。
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					// 发布模式。在这种模式下，Gin 会最小化输出信息，并且会对错误信息进行压缩，以保护应用程序的安全。这种模式适用于生产环境，可以提高应用程序的性能。
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					// 测试模式。在这种模式下，Gin 不会输出任何信息，并且错误信息也不会被压缩。这种模式适用于测试环境，可以方便测试人员查看测试结果。
					fmt.Printf("Testing")
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	// 读取并解析配置文件。如果读取失败，会输出错误信息并终止程序。
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 启动了配置文件的监控。如果配置文件发生了改变，会触发 OnConfigChange 回调函数。
	v.WatchConfig()

	// 回调函数中，使用 v.Unmarshal 将配置文件中的数据反序列化到全局变量 global.GVA_CONFIG 中。如果反序列化失败，会输出错误信息。
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 使用 v.Unmarshal 将配置文件中的数据反序列化到全局变量 global.GVA_CONFIG 中。如果反序列化失败，会输出错误信息。
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
