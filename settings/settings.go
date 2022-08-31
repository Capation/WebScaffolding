package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {

	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()           // 读取配置信息
	if err != nil {                      // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 当配置文件发生变化时调用一个回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed..")
	})

	//r := gin.Default()
	//// 访问/version的返回值会随配置文件的变化而变化
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//
	//if err := r.Run(
	//	fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
	//	panic(err)
	//}

	return
}
