// Package execution provides the gRPC server for the execution layer.
//
// Its procedures will be called from the conductor. It is responsible
// for immediately executing lists of ordered transactions that come from the shared sequencer.
package execution

import (
	"context"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/catalyst"
	executionv1 "github.com/ethereum/go-ethereum/grpc/gen/proto/execution/v1"
	"github.com/ethereum/go-ethereum/log"
)

// executionServiceServer is the implementation of the ExecutionServiceServer interface.
type ExecutionServiceServer struct {
	// NOTE - from the generated code:
	// All implementations must embed UnimplementedExecutionServiceServer
	// for forward compatibility
	executionv1.UnimplementedExecutionServiceServer

	consensus *catalyst.ConsensusAPI
	eth       *eth.Ethereum
}

func NewExecutionServiceServer(eth *eth.Ethereum) *ExecutionServiceServer {
	consensus := catalyst.NewConsensusAPI(eth)

	return &ExecutionServiceServer{
		eth:       eth,
		consensus: consensus,
	}
}

func (s *ExecutionServiceServer) DoBlock(ctx context.Context, req *executionv1.DoBlockRequest) (*executionv1.DoBlockResponse, error) {
	log.Info("DoBlock called request", "request", req)
	prevHeadHash := common.BytesToHash(req.PrevStateRoot)

	// The Engine API has been modified to use transactions from this mempool and abide by it's ordering.
	s.eth.TxPool().SetAstriaOrdered(req.Transactions)

	// Do the whole Engine API in a single loop
	startForkChoice := &engine.ForkchoiceStateV1{
		HeadBlockHash:      prevHeadHash,
		SafeBlockHash:      prevHeadHash,
		FinalizedBlockHash: prevHeadHash,
	}
	payloadAttributes := &engine.PayloadAttributes{
		Timestamp:             uint64(req.GetTimestamp().GetSeconds()),
		Random:                common.Hash{},
		SuggestedFeeRecipient: common.Address{},
	}
	fcStartResp, err := s.consensus.ForkchoiceUpdatedV1(*startForkChoice, payloadAttributes)
	if err != nil {
		return nil, err
	}
	payloadResp, err := s.consensus.GetPayloadV1(*fcStartResp.PayloadID)
	if err != nil {
		return nil, err
	}
	payloadStatus, err := s.consensus.NewPayloadV1(*payloadResp)
	if err != nil {
		return nil, err
	}
	newForkChoice := &engine.ForkchoiceStateV1{
		HeadBlockHash:      *payloadStatus.LatestValidHash,
		SafeBlockHash:      *payloadStatus.LatestValidHash,
		FinalizedBlockHash: *payloadStatus.LatestValidHash,
	}
	fcEndResp, err := s.consensus.ForkchoiceUpdatedV1(*newForkChoice, nil)
	if err != nil {
		return nil, err
	}

	res := &executionv1.DoBlockResponse{
		StateRoot: fcEndResp.PayloadStatus.LatestValidHash.Bytes(),
	}
	return res, nil
}

func (s *ExecutionServiceServer) InitState(ctx context.Context, req *executionv1.InitStateRequest) (*executionv1.InitStateResponse, error) {
	currHead := s.eth.BlockChain().CurrentHeader()
	res := &executionv1.InitStateResponse{
		StateRoot: currHead.Hash().Bytes(),
	}

	return res, nil
}
