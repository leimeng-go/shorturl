package main

import (
	"flag"
	"fmt"

	"shorturl/api/internal/config"
	"shorturl/api/internal/handler"
	"shorturl/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//解析配置文件，封装yaml,json等格式配置解析实例
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	//新建 rest api htt服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
