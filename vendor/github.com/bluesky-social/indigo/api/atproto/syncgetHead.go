// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.getHead

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// SyncGetHead_Output is the output of a com.atproto.sync.getHead call.
type SyncGetHead_Output struct {
	Root string `json:"root" cborgen:"root"`
}

// SyncGetHead calls the XRPC method "com.atproto.sync.getHead".
//
// did: The DID of the repo.
func SyncGetHead(ctx context.Context, c *xrpc.Client, did string) (*SyncGetHead_Output, error) {
	var out SyncGetHead_Output

	params := map[string]interface{}{
		"did": did,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.sync.getHead", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
