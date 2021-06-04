package api

import (
	"github.com/gogf/gf/net/ghttp"
	"log"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	log.Println("/hello")

	/*
		//rpc调用
		client := zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{
				Hosts: []string{"192.168.160.128:2379"},
				Key:   "kde-iot-go-historydata.rpc",
			},
		})

		conn := client.Conn()
		hello := historydata.NewHistorydataClient(conn)
		reply, err := hello.RedisFindHistory(context.Background(), &historydata.SnReq{
			Sn: "2344",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(reply)*/
	r.Response.Writeln("Hello World!")
}
