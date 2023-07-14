package ethapi

import (
	"context"
	"fmt"
	"os"

	client "github.com/astriaorg/go-sequencer-client"
	sqproto "github.com/astriaorg/go-sequencer-client/proto"
	"github.com/ethereum/go-ethereum/log"
)

const (
	defaultChainID               = "ethereum"
	defaultTendermintRPCEndpoint = "http://localhost:26657"
)

func sendTransactionToSequencer(ctx context.Context, txBytes []byte) error {
	var (
		chainID, tendermintRPCEndpoint string
	)

	envChainId := os.Getenv("CHAIN_ID")
	if envChainId == "" {
		chainID = defaultChainID
	}

	envTendermintRPCEndpoint := os.Getenv("TENDERMINT_RPC_ENDPOINT")
	if envTendermintRPCEndpoint == "" {
		tendermintRPCEndpoint = defaultTendermintRPCEndpoint
	}

	signer, err := client.GenerateSigner()
	if err != nil {
		return err
	}

	// default tendermint RPC endpoint
	c, err := client.NewClient(tendermintRPCEndpoint)
	if err != nil {
		return err
	}

	tx := &sqproto.UnsignedTransaction{
		Nonce: 1,
		Actions: []*sqproto.Action{
			{
				Value: &sqproto.Action_SequenceAction{
					SequenceAction: &sqproto.SequenceAction{
						ChainId: []byte(chainID),
						Data:    txBytes,
					},
				},
			},
		},
	}

	signed, err := signer.SignTransaction(tx)
	if err != nil {
		return err
	}

	resp, err := c.BroadcastTxSync(ctx, signed)
	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("failed to broadcast tx (error code %d): %s", resp.Code, resp.Log)
	}

	log.Info("successfully broadcasted tx to sequencer", "hash", resp.Hash)
	return nil
}
