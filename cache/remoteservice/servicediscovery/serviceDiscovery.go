package servicediscovery

// DiscoveryHandle 服务发现接口 实现服务发现的接口
type DiscoveryHandle interface {
	ServiceDiscovery(serviceName string, groupName string, clusters []string) (urlStr string, err error)
}

//func GetHealthHost(serviceName, groupName string, clusters []string) (urlStr string, err error) {
//	defer func() {
//		if r := recover(); r != nil {
//			msg := fmt.Sprintf("panic: %v", r)
//			logger.ErrLogger().WithFields(logrus.Fields{
//				"service":   "remoteservice",
//				"func":      "getHealthHost",
//				"backtrace": utils.GetStack(),
//			}).Error(msg)
//			err = errors.New(msg)
//		}
//	}()
//	instance, err := client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
//		Clusters:    clusters,
//		ServiceName: serviceName,
//		GroupName:   groupName,
//	})
//	if err != nil {
//		return "", err
//	}
//	return fmt.Sprintf("http://%v:%v", instance.Ip, instance.Port), nil
//}
//
//func init() {
//	nacosConfig := loader.LoadNacosConfig()
//	sc := make([]constant.ServerConfig, 0)
//	for _, server := range nacosConfig.Server {
//		sc = append(sc, constant.ServerConfig{
//			Scheme:      server.Scheme,
//			IpAddr:      server.IP,
//			Port:        server.Port,
//			ContextPath: server.ContextPath,
//		})
//	}
//
//	logLevel := "error"
//	if loader.IsDebugMode() {
//		logLevel = "info"
//	}
//
//	utils.CheckFilePath(nacosConfig.LogDir)
//	utils.CheckFilePath(nacosConfig.CacheDir)
//
//	cc := constant.ClientConfig{
//		NamespaceId:         nacosConfig.Namespace, //namespace id
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              nacosConfig.LogDir,
//		CacheDir:            nacosConfig.CacheDir,
//		LogLevel:            logLevel,
//	}
//
//	dns := constant.ClientConfig{
//		NamespaceId:         nacosConfig.DomainNameSpace, //namespace id
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              nacosConfig.LogDir,
//		CacheDir:            nacosConfig.CacheDir,
//		LogLevel:            logLevel,
//	}
//
//	var err error
//	client, err = clients.NewNamingClient(
//		vo.NacosClientParam{
//			ClientConfig:  &cc,
//			ServerConfigs: sc,
//		},
//	)
//
//	if err != nil {
//		panic(fmt.Errorf("nacos init err. err: %v", err))
//	}
//
//	configClient, err = clients.NewConfigClient(vo.NacosClientParam{
//		ClientConfig:  &dns,
//		ServerConfigs: sc,
//	})
//	if err != nil {
//		panic(fmt.Errorf("nacos init err. err: %v", err))
//	}
//}
