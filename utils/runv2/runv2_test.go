package runv2

import (
	"io"
	"strings"
	"testing"
)

func TestParseMediaPlaylist_WithMapURI(t *testing.T) {
	playlist := `#EXTM3U
#EXT-X-VERSION:6
#EXT-X-TARGETDURATION:6
#EXT-X-MAP:URI="init.mp4",BYTERANGE="798@0"
#EXT-X-KEY:METHOD=SAMPLE-AES,URI="skd://itunes.apple.com/P000000000/s1/e1",KEYFORMAT="com.apple.streamingkeydelivery",KEYFORMATVERSIONS="1"
#EXTINF:5.9999,
#EXT-X-BYTERANGE:12345@798
segment.mp4
#EXT-X-ENDLIST
`

	segments, mapURI, err := parseMediaPlaylist(io.NopCloser(strings.NewReader(playlist)))
	if err != nil {
		t.Fatalf("parseMediaPlaylist returned error: %v", err)
	}
	if mapURI != "init.mp4" {
		t.Fatalf("expected map URI %q, got %q", "init.mp4", mapURI)
	}
	if len(segments) == 0 || segments[0] == nil {
		t.Fatal("expected at least one media segment")
	}
	if segments[0].URI != "segment.mp4" {
		t.Fatalf("expected first segment URI %q, got %q", "segment.mp4", segments[0].URI)
	}
}

func TestParseMediaPlaylist_WithoutMapURI(t *testing.T) {
	playlist := `#EXTM3U
#EXT-X-VERSION:6
#EXT-X-TARGETDURATION:6
#EXT-X-KEY:METHOD=SAMPLE-AES,URI="skd://itunes.apple.com/P000000000/s1/e1",KEYFORMAT="com.apple.streamingkeydelivery",KEYFORMATVERSIONS="1"
#EXTINF:5.9999,
segment.mp4
#EXT-X-ENDLIST
`

	segments, mapURI, err := parseMediaPlaylist(io.NopCloser(strings.NewReader(playlist)))
	if err != nil {
		t.Fatalf("parseMediaPlaylist returned error: %v", err)
	}
	if mapURI != "" {
		t.Fatalf("expected empty map URI, got %q", mapURI)
	}
	if len(segments) == 0 || segments[0] == nil {
		t.Fatal("expected at least one media segment")
	}
}
