package initialize

import (
	"gin-project/global"
	"gin-project/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	// 创建一个新的缓存
	global.BlackCache = local_cache.NewCache(
		// local_cache.SetDefaultExpire 函数设置缓存项的默认过期时间
		local_cache.SetDefaultExpire(dr),
	)
	// 调用 local_cache.NewCache 函数后，程序就可以使用该缓存来存储和获取数据。
	// 例如，可以使用 cache.Set 函数向缓存中添加新的数据项，或使用 cache.Get 函数从缓存中获取数据项。
}
