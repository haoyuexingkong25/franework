package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// ServiceRegister 服务注册
func ServiceRegister(serviceName string, port uint64) error {
	namingClient, err := clients.NewNamingClient(Client())
	if err != nil {
		return err
	}

	_, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        port,
		ServiceName: serviceName,
		Enable:      true,
		Healthy:     true,
	})

	return nil
}

// ServiceDiscoveries 服务发现
func ServiceDiscoveries(serviceName string) (string, error) {
	namingClient, err := clients.NewNamingClient(Client())
	if err != nil {
		return "", err
	}
	services, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
	})
	sprintf := fmt.Sprintf("%v:%v", services.Hosts[3].Ip, services.Hosts[3].Port)
	return sprintf, nil
}
