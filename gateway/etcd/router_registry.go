package etcd

import (
	"context"
	"encoding/json"
	"gateway/router"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"log"
)

func NewEtcdRoutingMapper(etcdClient *clientv3.Client) *router.RoutingMapper {
	mapper := router.NewRoutingMapper()
	readEtcdData(mapper, etcdClient)
	go watchEtcdChange(mapper, etcdClient)
	return mapper
}

/**
从 Etcd 中全量读取数据路由数据
*/
func readEtcdData(mapper *router.RoutingMapper, etcdClient *clientv3.Client) {
	kvClient := clientv3.NewKV(etcdClient)
	resp, err := kvClient.Get(context.TODO(), "/gateway/routers/", clientv3.WithPrefix())
	if err != nil {
		zap.Error(err)
		return
	}
	for _, kv := range resp.Kvs {
		var route router.Router
		err := json.Unmarshal(kv.Value, &route)
		if err != nil {
			log.Fatal(err)
		}
		mapper.AddRouter(&route)
	}
}

/**
监听 Etcd 中数据变化
*/
func watchEtcdChange(mapper *router.RoutingMapper, etcdClient *clientv3.Client) {
	watchClient := clientv3.NewWatcher(etcdClient)
	defer func(watchClient clientv3.Watcher) {
		err := watchClient.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(watchClient)

	watchRespChan := watchClient.Watch(context.TODO(), "/gateway/routers/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				route := router.Router{}
				err := json.Unmarshal(event.Kv.Value, &route)
				if err != nil {
					log.Fatal(err)
				}
				mapper.AddRouter(&route)
			case mvccpb.DELETE:
				route := router.Router{}
				err := json.Unmarshal(event.PrevKv.Value, &route)
				if err != nil {
					log.Fatal(err)
				}
				mapper.DelRouter(&route)
			}
		}
	}
}
