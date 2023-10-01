package etcdplug

import (
	"context"

	"github.com/boxcolli/pepperlink/plugs"
)

func NewEtcdPlug(ctx context.Context, opt EtcdOption) (plugs.Plug, error) {
	clu := &etcdPlug{}
	{
		var err error
		clu.etcd, clu.cli, err = initEtcd(ctx, opt)
		if err != nil {
			return nil, err
		}

	}
	return clu, nil
}