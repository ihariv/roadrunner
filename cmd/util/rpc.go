package util

import (
	"errors"
	"github.com/ihariv/roadrunner/service"
	rrpc "github.com/ihariv/roadrunner/service/rpc"
	"net/rpc"
)

// RPCClient returns RPC client associated with given rr service container.
func RPCClient(container service.Container) (*rpc.Client, error) {
	svc, st := container.Get(rrpc.ID)
	if st < service.StatusOK {
		return nil, errors.New("RPC service is not configured")
	}

	return svc.(*rrpc.Service).Client()
}
