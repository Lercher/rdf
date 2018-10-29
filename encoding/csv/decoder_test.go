package csv

import (
	"compress/gzip"
	"os"
	"runtime"
	"testing"
)

// Download from https://get-information-schools.service.gov.uk/Downloads and gzip it:
const path = `D:\Profiles\mlercher\Downloads\edubasealldata20181029.csv.gz`

// comes with git reop
const smallpath = `edubasealldata20181029_20.csv`

func TestLoadSmallCSVFile(t *testing.T) {
	f, err := os.Open(smallpath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dec := NewDecoder(f, "http://education.data.gov.uk/def/school/establishment/", "http://education.data.gov.uk/def/school/")
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if g.CountValues() != 667 {
		t.Errorf("want %d values, got %d", 667, g.CountValues())
	}
	if len(g.Dataset) != 2640 {
		t.Errorf("want %d triples, got %d", 2640, len(g.Dataset))
	}
	for i, tr := range g.Dataset {
		if i > 2600 {
			t.Logf("%2d %s", i+1, tr.String(g))
		}
	}

	got := g.Dataset[2613].String(g)
	want := `[(string:http://education.data.gov.uk/def/school/establishment/100019) (string:http://education.data.gov.uk/def/school/AdministrativeWard%20%28code%29) (string:E05000135)]`
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestLoadLargeCSVFile(t *testing.T) {
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	zr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer zr.Close()

	dec := NewDecoder(zr, "http://education.data.gov.uk/def/school/establishment/", "http://education.data.gov.uk/def/school/")
	dec.Reindex = false
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if g.CountValues() != 433796 {
		t.Errorf("want %d values, got %d", 433796, g.CountValues())
	}
	if len(g.Dataset) != 6249408 {
		t.Errorf("want %d triples, got %d", 6249408, len(g.Dataset))
	}
	runtime.GC()
	PrintMemUsage(t, "dataset only")

	g.RebuildIndex()
	runtime.GC()
	PrintMemUsage(t, "reindexed")
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// See https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage(t *testing.T, kind string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	t.Logf("Alloc %v MiB, %s", bToMb(m.Alloc), kind)
	t.Logf("\tTotalAlloc %v MiB", bToMb(m.TotalAlloc))
	t.Logf("\tSys        %v MiB", bToMb(m.Sys))
	t.Logf("\tNumGC      %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
