// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.searchPosts

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// FeedSearchPosts_Output is the output of a app.bsky.feed.searchPosts call.
type FeedSearchPosts_Output struct {
	Cursor *string `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	// hitsTotal: Count of search hits. Optional, may be rounded/truncated, and may not be possible to paginate through all hits.
	HitsTotal *int64               `json:"hitsTotal,omitempty" cborgen:"hitsTotal,omitempty"`
	Posts     []*FeedDefs_PostView `json:"posts" cborgen:"posts"`
}

// FeedSearchPosts calls the XRPC method "app.bsky.feed.searchPosts".
//
// author: Filter to posts by the given account. Handles are resolved to DID before query-time.
// cursor: Optional pagination mechanism; may not necessarily allow scrolling through entire result set.
// domain: Filter to posts with URLs (facet links or embeds) linking to the given domain (hostname). Server may apply hostname normalization.
// lang: Filter to posts in the given language. Expected to be based on post language field, though server may override language detection.
// mentions: Filter to posts which mention the given account. Handles are resolved to DID before query-time. Only matches rich-text facet mentions.
// q: Search query string; syntax, phrase, boolean, and faceting is unspecified, but Lucene query syntax is recommended.
// since: Filter results for posts after the indicated datetime (inclusive). Expected to use 'sortAt' timestamp, which may not match 'createdAt'. Can be a datetime, or just an ISO date (YYYY-MM-DD).
// sort: Specifies the ranking order of results.
// tag: Filter to posts with the given tag (hashtag), based on rich-text facet or tag field. Do not include the hash (#) prefix. Multiple tags can be specified, with 'AND' matching.
// until: Filter results for posts before the indicated datetime (not inclusive). Expected to use 'sortAt' timestamp, which may not match 'createdAt'. Can be a datetime, or just an ISO date (YYY-MM-DD).
// url: Filter to posts with links (facet links or embeds) pointing to this URL. Server may apply URL normalization or fuzzy matching.
func FeedSearchPosts(ctx context.Context, c *xrpc.Client, author string, cursor string, domain string, lang string, limit int64, mentions string, q string, since string, sort string, tag []string, until string, url string) (*FeedSearchPosts_Output, error) {
	var out FeedSearchPosts_Output

	params := map[string]interface{}{
		"author":   author,
		"cursor":   cursor,
		"domain":   domain,
		"lang":     lang,
		"limit":    limit,
		"mentions": mentions,
		"q":        q,
		"since":    since,
		"sort":     sort,
		"tag":      tag,
		"until":    until,
		"url":      url,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.searchPosts", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
