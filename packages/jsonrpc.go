package packages

import "github.com/ybbus/jsonrpc"

func init() {
	Packages["jsonrpc"] = map[string]interface{}{
		"Params":            jsonrpc.Params,
		"NewClient":         jsonrpc.NewClient,
		"NewClientWithOpts": jsonrpc.NewClientWithOpts,
		"RPCRequests":       RPCRequests,
		"NewRequest":        jsonrpc.NewRequest,
	}
	PackageTypes["jsonrpc"] = map[string]interface{}{
		"HTTPError":     jsonrpc.HTTPError{},
		"RPCClientOpts": jsonrpc.RPCClientOpts{},
		"RPCError":      jsonrpc.RPCError{},
		"RPCRequest":    jsonrpc.RPCRequest{},
		"RPCResponse":   jsonrpc.RPCResponse{},
	}
}

func RPCRequests(r []*jsonrpc.RPCRequest) {
	return jsonrpc.RPCRequests(r)
}

func RPCResponses(r []*jsonrpc.RPCResponse) jsonrpc.RPCResponses [
	return jsonrpc.RPCResponses(r)
]