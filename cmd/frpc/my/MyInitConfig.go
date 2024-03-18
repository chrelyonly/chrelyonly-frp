package my

import (
	"fmt"
	"github.com/fatedier/frp/cmd/frpc/user"
	"github.com/fatedier/frp/cmd/frpc/util"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"math/rand"
	"net"
	"os"
	"strconv"
)

// MyInitConfig	 初始化我的div配置
func InitConfig() (*v1.ClientCommonConfig,
	[]v1.ProxyConfigurer,
	[]v1.VisitorConfigurer,
	error) {
	var (
		cliCfg      *v1.ClientCommonConfig
		proxyCfgs   = make([]v1.ProxyConfigurer, 0)
		visitorCfgs = make([]v1.VisitorConfigurer, 0)
	)
	myConfig := util.DivConfigObject{}
	//是否定制版本
	if true {
		user.Luobo()
		myConfig.ServerName = "frp.imbhj.com"
		err := net.ErrClosed
		//选择协议
		cliCfg, proxyCfgs, visitorCfgs, err = BuildProtocol(myConfig, false)
		if err != nil {
			fmt.Println(err)
			fmt.Println("出现错误,程序退出")
			os.Exit(1)
		}
		return cliCfg, proxyCfgs, visitorCfgs, err
	} else {
		color.Yellow("选择启动配置模式：1快速映射模式,2高级模式,输入其他默认1")
		//	分快速模式和高级配置模式
		var modelTYpe int
		_, err := fmt.Scanln(&modelTYpe)
		if err != nil {
			modelTYpe = 1
		}
		if modelTYpe == 1 {
			color.Magenta("快速映射模式")
			//服务器节点 默认 阿里云
			myConfig.ServerName = "frp.chrelyonly.cn"
			//选择协议
			cliCfg, proxyCfgs, visitorCfgs, err = BuildProtocol(myConfig, true)
			if err != nil {
				fmt.Println(err)
				fmt.Println("出现错误,程序退出")
				os.Exit(1)
			}
		} else {
			color.Magenta("当前运行模式(手动生成配置文件)")

			//判断哪个服务器
			color.Green("选择服务器节点(默认上海节点): ")
			color.Yellow("1.上海阿里云")
			color.Cyan("2.四川-腾讯")
			color.Blue("3.美国")
			// 选择的服务名称
			var serverNameTemp = ""
			// 选择的服务编号节点
			var ServerNode int
			_, err := fmt.Scanln(&ServerNode)
			if err != nil {
				ServerNode = 1
			}
			if ServerNode == 1 {
				serverNameTemp = "上海阿里云"
				myConfig.ServerName = "frp.chrelyonly.cn"
			} else if ServerNode == 2 {
				serverNameTemp = "四川-腾讯"
				myConfig.ServerName = "frp.tx.chrelyonly.cn"
			} else if ServerNode == 3 {
				serverNameTemp = "美国"
				myConfig.ServerName = "frp.chrelyonly.cf"
			} else {
				serverNameTemp = "上海阿里云(限速10-30/mbps)"
				myConfig.ServerName = "frp.chrelyonly.cn"
			}
			color.Red("当前服务器节点: " + serverNameTemp + " 服务器地址: " + myConfig.ServerName)

			//选择协议
			cliCfg, proxyCfgs, visitorCfgs, err = BuildProtocol(myConfig, false)
			if err != nil {
				fmt.Println(err)
				fmt.Println("出现错误,程序退出")
				os.Exit(1)
			}
		}
		return cliCfg, proxyCfgs, visitorCfgs, err
	}
}

func BuildProtocol(divConfig util.DivConfigObject, flag bool) (*v1.ClientCommonConfig,
	[]v1.ProxyConfigurer,
	[]v1.VisitorConfigurer,
	error) {

	//初始化一些配置餐宿
	divConfig.Token = "1172576293"
	divConfig.ServerAddr = divConfig.ServerName
	divConfig.ServerPort = 7000
	//泛域名后缀
	domain := "." + divConfig.ServerName
	//需要一个集合
	var list []util.ListMap
	//创建一个对象 存储数据
	var listMap util.ListMap
	//代理名称
	proxyName := "frp" + strconv.Itoa(rand.Intn(99999))
	if flag {
		color.Green("默认以http协议进行代理")
		fmt.Println("请输入代理ip地址(默认使用127.0.0.1),可以用作反代地址(公网ip或局域网ip或域名)")
		var localhost string
		_, err := fmt.Scanln(&localhost)
		if err != nil {
			localhost = "127.0.0.1"
		}
		listMap.LocalIp = localhost
		fmt.Println("请输入代理端口(默认80)")
		var LocalPort string
		_, err = fmt.Scanln(&LocalPort)
		if err != nil {
			LocalPort = "80"
		}
		listMap.LocalPort = LocalPort
		listMap.Type = "http"
		listMap.ExeType = "2"
		listMap.CustomDomains = proxyName + domain
	} else {

		//进行用户输入判断
		color.Green("支持代理协议: ")
		color.Magenta("1.以tcp协议模式")
		color.Green("2.以http协议模式")
		color.Yellow("3.web资源服务器模式")
		color.Blue("4.反代目标域名模式(已http2协议代理,目标必须带有ssl证书)")
		//fmt.Println("# 5(查看版本更新说明)")
		color.Yellow("# 请选择代理协议(输入1234),默认2:")
		var proxyType string
		_, err := fmt.Scanln(&proxyType)
		if err != nil {
			proxyType = "2"
		}
		//如果是资源服务器
		if proxyType == "3" {
			listMap.ExeType = "3"
			listMap.Type = proxyType
			fmt.Println("是否需要配置域名(true/false,默认true)")
			var isDom bool
			_, err = fmt.Scanln(&isDom)
			if err != nil {
				isDom = true
			}
			listMap.IsDom = isDom
			if isDom {
				listMap.Type = "http"
				fmt.Println("请输入代理转发域名,(前缀)." + divConfig.ServerName + ",只用输入需要的前缀,默认随机一个字符串")
				var CustomDomains string
				_, err = fmt.Scanln(&CustomDomains)
				if err != nil {
					CustomDomains = proxyName
				}
				listMap.CustomDomains = CustomDomains + domain
			} else {
				listMap.Type = "tcp"
				fmt.Println("请输入代理转发端口,默认随机50000-60000")
				var RemotePort string
				_, err = fmt.Scanln(&RemotePort)
				if err != nil {
					RemotePort = strconv.Itoa(50000 + rand.Intn(9999))
				}
				listMap.RemotePort = RemotePort
			}
			listMap.Plugin = "static_file"
			//配置路径
			fmt.Println("请配置路径 (E:\\dev\\Git   /www/temp/) ,默认 (./)当前目录")
			var path string
			_, err = fmt.Scanln(&path)
			if err != nil {
				path = "./"
			}
			listMap.PluginLocalPath = path
			//是否需要访问前缀
			fmt.Println("请配置路径前缀 /img   (默认不需要)")
			var staticPath string
			_, err = fmt.Scanln(&staticPath)
			if err != nil {
				staticPath = ""
			}
			listMap.PluginStripPrefix = staticPath

			fmt.Println("请配置访问用户名  (默认不需要)")
			var username string
			_, err = fmt.Scanln(&username)
			if err != nil {
				listMap.PluginHttpUser = ""
			} else {
				listMap.PluginHttpUser = username
				fmt.Println("请配置访问密码")
				var password string
				_, err = fmt.Scanln(&password)
				if err != nil {
					log.Infof("请配置密码,默认123456789")
					password = "123456789"
				}
				listMap.PluginHttpPasswd = password
			}
		}
		//如果域名转发
		if proxyType == "4" {
			listMap.ExeType = "4"
			listMap.Type = "http"
			fmt.Println("请输入代理转发域名,(前缀)." + divConfig.ServerName + ",只用输入需要的前缀,默认随机一个字符串")
			var CustomDomains string
			_, err = fmt.Scanln(&CustomDomains)
			if err != nil {
				CustomDomains = proxyName
			}
			listMap.CustomDomains = CustomDomains + domain
			listMap.Plugin = "http2https"
			fmt.Println("请输入反代目标域名,(www.baidu.com,127.0.0.1:80/443,默认www.baidu.com)")
			var PluginLocalAddr string
			_, err = fmt.Scanln(&PluginLocalAddr)
			if err != nil {
				PluginLocalAddr = "www.baidu.com"
			}
			listMap.PluginLocalAddr = PluginLocalAddr
			fmt.Println("请填写域名请求头,(www.baidu.com,127.0.0.1(可不填,填错403,默认填写的反代目标域名))")
			var PluginHostHeaderRewrite string
			_, err = fmt.Scanln(&PluginHostHeaderRewrite)
			if err != nil {
				PluginHostHeaderRewrite = PluginLocalAddr
			}
			listMap.PluginHostHeaderRewrite = PluginHostHeaderRewrite
			//fmt.Println("请填写域名来源请求头(可不填,填错403,默认填写的反代目标域名)")
			//var PluginHeaderX string
			//_, err = fmt.Scanln(&PluginHeaderX)
			//if err != nil {
			//	PluginHeaderX = PluginLocalAddr
			//}
			//listMap.PluginHeaderX = PluginHeaderX
		}
		//如果是正常穿透 内网映射使用流量则进行填写
		if proxyType == "1" || proxyType == "2" {
			fmt.Println("请输入代理ip地址(默认使用127.0.0.1),可以用作反代地址(公网ip或局域网ip或域名)")
			var localhost string
			_, err = fmt.Scanln(&localhost)
			if err != nil {
				localhost = "127.0.0.1"
			}
			listMap.LocalIp = localhost
			fmt.Println("请输入代理端口(默认80)")
			var LocalPort string
			_, err = fmt.Scanln(&LocalPort)
			if err != nil {
				LocalPort = "80"
			}
			listMap.LocalPort = LocalPort
			//[判断是tcp 还是 http
			if proxyType == "1" {
				listMap.Type = "tcp"
				listMap.ExeType = "1"
				fmt.Println("请输入代理转发端口(50000-60000),默认随机50000-60000")
				var RemotePort string
				_, err = fmt.Scanln(&RemotePort)
				if err != nil {
					RemotePort = strconv.Itoa(50000 + rand.Intn(9999))
				}
				listMap.RemotePort = RemotePort
			}
			if proxyType == "2" {
				listMap.Type = "http"
				listMap.ExeType = "2"
				fmt.Println("请输入代理转发域名,(前缀)." + divConfig.ServerName + ",只用输入需要的前缀,默认随机一个字符串(提供免费域名与多个端口)")
				var CustomDomains string
				_, err = fmt.Scanln(&CustomDomains)
				if err != nil {
					CustomDomains = proxyName
				}
				listMap.CustomDomains = CustomDomains + domain

				fmt.Println("请配置访问用户名  (默认不需要)")
				var username string
				_, err = fmt.Scanln(&username)
				if err != nil {
					listMap.PluginHttpUser = ""
				} else {
					listMap.PluginHttpUser = username
					fmt.Println("请配置访问密码")
					var password string
					_, err = fmt.Scanln(&password)
					if err != nil {
						log.Infof("请配置密码,默认123456789")
						password = "123456789"
					}
					listMap.PluginHttpPasswd = password
				}
			}
		}
		fmt.Println("是否加密数据传输(true/false,默认true),参考tls")
		var UseEncryption bool
		_, err = fmt.Scanln(&UseEncryption)
		if err != nil {
			UseEncryption = true
		}
		listMap.UseEncryption = UseEncryption

		fmt.Println("是否压缩数据传输(true/false,默认true),参考gzip")
		var UseCompression bool
		_, err = fmt.Scanln(&UseCompression)
		if err != nil {
			UseCompression = true
		}
		listMap.UseCompression = UseCompression
	}

	listMap.Comment = proxyName
	//保存代理配置
	util.GlobalListMap = listMap
	util.GlobalDivConfigObject = divConfig
	list = append(list, listMap)
	return MyDivInputExeJsonList(list, divConfig)
}

// MyDivInputExeJsonList 处理代理配置文件
func MyDivInputExeJsonList(list []util.ListMap, divConfig util.DivConfigObject) (*v1.ClientCommonConfig,
	[]v1.ProxyConfigurer,
	[]v1.VisitorConfigurer,
	error) {
	var (
		//cfg  :=  *v1.ClientCommonConfig
		cfg         = &v1.ClientCommonConfig{}
		proxyCfgs   = make([]v1.ProxyConfigurer, 0)
		visitorCfgs = make([]v1.VisitorConfigurer, 0)
	)
	for i := 0; i < len(list); i++ {
		if list[i].ExeType == "1" {
			//可能有多个tcp代理
			info := v1.TCPProxyConfig{}
			info.LocalIP = list[i].LocalIp
			info.Transport.BandwidthLimitMode = "client"
			info.LocalPort, _ = strconv.Atoi(list[i].LocalPort)
			info.Name = list[i].Comment
			info.RemotePort, _ = strconv.Atoi(list[i].RemotePort)
			info.Transport.UseEncryption = list[i].UseEncryption
			info.Transport.UseCompression = list[i].UseCompression
			proxyCfgs = append(proxyCfgs, &info)
			//tempFrpConFigList[list[i].Comment] = &info
		}
		if list[i].ExeType == "2" {
			//可能有多个tcp代理
			info := v1.HTTPProxyConfig{}
			info.LocalIP = list[i].LocalIp
			info.Type = "http"
			info.LocalPort, _ = strconv.Atoi(list[i].LocalPort)
			info.Name = list[i].Comment
			info.CustomDomains = []string{list[i].CustomDomains}
			info.HTTPUser = list[i].PluginHttpUser
			info.HTTPPassword = list[i].PluginHttpPasswd
			info.Transport.BandwidthLimitMode = "client"
			info.Transport.UseEncryption = list[i].UseEncryption
			info.Transport.UseCompression = list[i].UseCompression
			info.HostHeaderRewrite = list[i].CustomDomains
			//tempFrpConFigList[list[i].Comment] = &info
			proxyCfgs = append(proxyCfgs, &info)
		}
		if list[i].ExeType == "3" {
			if list[i].IsDom {
				info := v1.HTTPProxyConfig{}
				info.Name = list[i].Comment
				info.Type = "http"
				info.Plugin.ClientPluginOptions = &v1.StaticFilePluginOptions{
					LocalPath:    list[i].PluginLocalPath,
					StripPrefix:  list[i].PluginStripPrefix,
					HTTPUser:     list[i].PluginHttpUser,
					HTTPPassword: list[i].PluginHttpPasswd,
				}
				info.LocalIP = "127.0.0.1"
				info.Plugin.Type = "static_file"
				info.Transport.BandwidthLimitMode = "client"
				info.DomainConfig.CustomDomains = []string{list[i].CustomDomains}
				info.LocalPort = 0
				info.Transport.UseEncryption = list[i].UseEncryption
				info.Transport.UseCompression = list[i].UseCompression
				info.HostHeaderRewrite = list[i].CustomDomains
				//tempFrpConFigList[list[i].Comment] = &info
				proxyCfgs = append(proxyCfgs, &info)
			} else {
				info := v1.TCPProxyConfig{}
				info.Transport.BandwidthLimitMode = "client"
				info.Plugin.Type = "static_file"
				info.Name = list[i].Comment
				info.RemotePort, _ = strconv.Atoi(list[i].RemotePort)
				info.Plugin.ClientPluginOptions = &v1.StaticFilePluginOptions{
					LocalPath:    list[i].PluginLocalPath,
					StripPrefix:  list[i].PluginStripPrefix,
					HTTPUser:     list[i].PluginHttpUser,
					HTTPPassword: list[i].PluginHttpPasswd,
				}
				info.Transport.UseEncryption = list[i].UseEncryption
				info.Transport.UseCompression = list[i].UseCompression
				//tempFrpConFigList[list[i].Comment] = &info
				proxyCfgs = append(proxyCfgs, &info)
			}
		}
		if list[i].ExeType == "4" {
			//可能有多个tcp代理
			info := v1.HTTPProxyConfig{}
			info.Plugin.Type = "http2https"
			info.Type = "http"
			info.Transport.BandwidthLimitMode = "client"
			info.Name = list[i].Comment
			info.CustomDomains = []string{list[i].CustomDomains}
			info.Plugin.ClientPluginOptions = &v1.HTTP2HTTPSPluginOptions{
				LocalAddr:         list[i].PluginLocalAddr,
				HostHeaderRewrite: list[i].PluginHostHeaderRewrite,
				RequestHeaders: v1.HeaderOperations{
					Set: map[string]string{
						"plugin_header_X-From-Where": list[i].PluginHeaderX,
					},
				},
			}
			//tempFrpConFigList[list[i].Comment] = &info
			proxyCfgs = append(proxyCfgs, &info)
		}
	}
	buildCfg(cfg, divConfig)
	return cfg, proxyCfgs, visitorCfgs, nil

}
func buildCfg(cfg *v1.ClientCommonConfig, divConfig util.DivConfigObject) {

	//组装返回数据
	//tempFrpConFigList
	//服务器地址
	cfg.ServerAddr = divConfig.ServerAddr
	//服务器端口
	cfg.ServerPort = divConfig.ServerPort
	//服务器校验token
	cfg.Auth.Token = divConfig.Token
	//服务器认证方法
	cfg.Auth.Method = "token"
	//服务器认证超时时间
	cfg.Transport.DialServerTimeout = 10
	//保持长链接
	cfg.Transport.DialServerKeepAlive = 7200
	//# 如果使用tcp流多路复用，默认值为true，必须与FRP相同
	cfg.Transport.TCPMux = lo.ToPtr(true)
	//#指定tcp mux的保持活动间隔。
	//#仅当tcp_mux为真时有效。
	cfg.Transport.TCPMuxKeepaliveInterval = 60
	//日志输出模式
	cfg.Log.To = "console"
	cfg.Log.DisablePrintColor = false
	cfg.Log.Level = "info"
	//日志最大保存时间
	cfg.Log.MaxDays = 3
	//# 将提前建立连接，默认值为零
	cfg.Transport.PoolCount = 5
	//#决定是否在首次登录失败时退出程序，否则继续重新登录到frps
	//#默认值为true
	cfg.LoginFailExit = lo.ToPtr(false)
	//#用于连接到服务器的通信协议
	//#现在它支持tcp、kcp和websocket，默认为tcp
	//"tcp", "kcp", "quic", "websocket" and "wss"
	cfg.Transport.Protocol = "tcp"
	//#如果tls_enable为真，frpc将通过tls连接FRP
	cfg.Transport.TLS.Enable = lo.ToPtr(false)
	//#默认情况下，如果启用tls，frpc将连接FRP和第一个自定义字节。
	//#如果DisableCustomTLSFirstByte为true，frpc将不发送该自定义字节。
	cfg.Transport.TLS.DisableCustomTLSFirstByte = lo.ToPtr(false)
	//连接心跳
	cfg.Transport.HeartbeatInterval = 10
	//连接心跳超时
	cfg.Transport.HeartbeatTimeout = 60
	//传输数据包大小
	cfg.UDPPacketSize = 1500
}
