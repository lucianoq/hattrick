package chpp

import "testing"

func TestDecodeXML_Success(t *testing.T) {
	buf := []byte(`<?xml version="1.0"?>
<HattrickData>
	<FileName>bookmarks.xml</FileName>
	<Version>1.0</Version>
	<User>123</User>
	<FetchedDate>2026-07-09 12:00:00</FetchedDate>
	<BookmarkList>
		<Bookmark>
			<BookmarkID>1</BookmarkID>
		</Bookmark>
	</BookmarkList>
</HattrickData>`)

	x, err := DecodeXML[BookmarksXML](buf, BookmarksAPIFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if x.FileName != "bookmarks.xml" {
		t.Errorf("FileName = %q, want %q", x.FileName, "bookmarks.xml")
	}
	if len(x.Bookmarks) != 1 || x.Bookmarks[0].BookmarkID != "1" {
		t.Errorf("Bookmarks = %+v, want a single bookmark with ID 1", x.Bookmarks)
	}
}

func TestDecodeXML_ChppError(t *testing.T) {
	buf := []byte(`<?xml version="1.0"?>
<HattrickData>
	<FileName>bookmarks.xml</FileName>
	<Version>1.0</Version>
	<Error>This action requires the logged in user to be a Supporter</Error>
	<ErrorCode>6</ErrorCode>
</HattrickData>`)

	_, err := DecodeXML[BookmarksXML](buf, BookmarksAPIFile)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if err.Error() != "This action requires the logged in user to be a Supporter" {
		t.Errorf("err = %q, want the CHPP <Error> text", err.Error())
	}
}

func TestDecodeXML_FileNameMismatch(t *testing.T) {
	buf := []byte(`<?xml version="1.0"?>
<HattrickData>
	<FileName>club.xml</FileName>
	<Version>1.0</Version>
</HattrickData>`)

	_, err := DecodeXML[BookmarksXML](buf, BookmarksAPIFile)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
