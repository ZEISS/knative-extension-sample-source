// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.repo.uploadBlob

import (
	"context"
	"io"

	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// RepoUploadBlob_Output is the output of a com.atproto.repo.uploadBlob call.
type RepoUploadBlob_Output struct {
	Blob *util.LexBlob `json:"blob" cborgen:"blob"`
}

// RepoUploadBlob calls the XRPC method "com.atproto.repo.uploadBlob".
func RepoUploadBlob(ctx context.Context, c *xrpc.Client, input io.Reader) (*RepoUploadBlob_Output, error) {
	var out RepoUploadBlob_Output
	if err := c.Do(ctx, xrpc.Procedure, "*/*", "com.atproto.repo.uploadBlob", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
