package font

import (
	"testing"
	"Tiny-Go-Crawler/Crawler/fetcher"
)

func TestFont(t *testing.T) {
	file, _ := fetcher.Fetch("https://hf.58.com/hezu/38560664837252x.shtml?shangquan=datonglu")
	GetFontFile(file)
}