package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func Client() vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

func GetConfig() (string, error) {
	configClient, err := clients.NewConfigClient(Client())
	if err != nil {
		return "", err
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.dataId"),
		Group:  viper.GetString("nacos.group"),
	})
	return content, nil
}
