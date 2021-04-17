package controller

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var officialAccount *officialaccount.OfficialAccount

//在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "wx0b534af1ea7cbb1e",
		AppSecret: "66b0d0eab7f2c6799743a79018c91f99",
		Token:     "1234",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	officialAccount = wc.GetOfficialAccount(cfg)

	return officialAccount
}
