package service

import (
	"getaway/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var OfficialAccount *officialaccount.OfficialAccount

//在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat(gCfg *config.Config) *officialaccount.OfficialAccount {

	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     gCfg.OfficialAccountConfig.AppID,
		AppSecret: gCfg.OfficialAccountConfig.AppSecret,
		Token:     gCfg.OfficialAccountConfig.Token,
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	OfficialAccount = wc.GetOfficialAccount(cfg)

	return OfficialAccount
}
