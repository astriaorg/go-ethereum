// Package execution provides the gRPC server for the execution layer.
//
// Its procedures will be called from the conductor. It is responsible
// for immediately executing lists of ordered transactions that come from the shared sequencer.
package execution

import (
	"context"

	executionv1 "github.com/ethereum/go-ethereum/grpc/gen/proto/execution/v1"
	"google.golang.org/grpc"
)

// executionServiceServer is the implementation of the ExecutionServiceServer interface.
type executionServiceServer struct {
	// NOTE - from the generated code:
	// All implementations must embed UnimplementedExecutionServiceServer
	// for forward compatibility
	executionv1.UnimplementedExecutionServiceServer

	// TODO - will need access to the consensus api to call functions for building a block
	// e.g. getPayload, newPayload, forkchoiceUpdated

	// TODO - will need access to forkchoice on first run.
	// this will probably be passed in when calling NewServer
}

// FIXME - how do we know which hash to start with? will probably need another api function like
// GetHeadHash() to get the head hash of the forkchoice

func (s *executionServiceServer) DoBlock(ctx context.Context, req *executionv1.DoBlockRequest) (*executionv1.DoBlockResponse, error) {
	// Request.Header.ParentHash needs to match forkchoice head hash
	// ParentHash should be the forkchoice head of the last block

	// TODO - need to call consensus api to build a block
	return nil, nil
}

func NewServer() *grpc.Server {
	server := grpc.NewServer()
	executionv1.RegisterExecutionServiceServer(server, &executionServiceServer{})
	return server
}
