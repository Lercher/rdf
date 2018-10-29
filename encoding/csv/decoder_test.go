package csv

import (
	"time"
	"compress/gzip"
	"os"
	"runtime"
	"testing"

	"github.com/lercher/rdf/graph"
)

// Download from https://get-information-schools.service.gov.uk/Downloads and gzip it:
const path = `D:\Profiles\mlercher\Downloads\edubasealldata20181029.csv.gz`

// comes with git reop
const smallpath = `edubasealldata20181029_20.csv`

const (
	nsEst    = `http://education.data.gov.uk/def/school/establishment/`
	nsSchool = `http://education.data.gov.uk/def/school/`
)

func small() (*graph.Graph, error) {
	f, err := os.Open(smallpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := NewDecoder(f, nsEst, nsSchool)
	return dec.Decode()
}

func TestLoadSmallCSVFile(t *testing.T) {
	g, err := small()
	if err != nil {
		t.Fatal(err)
	}
	if g.CountValues() != 667 {
		t.Errorf("want %d values, got %d", 667, g.CountValues())
	}
	if g.CountTriples() != 2640 {
		t.Errorf("want %d triples, got %d", 2640, g.CountTriples())
	}

	ds := g.DistinctS()
	if len(ds) != 20 {
		for _, s := range ds {
			t.Log("S:", s.String(g))
		}
		t.Errorf("want %d distinct S, got %d", 20, len(ds))
	}

	dp := g.DistinctP()
	if len(dp) != 132 {
		for _, p := range dp {
			t.Log("P:", p.String(g))
		}
		t.Errorf("want %d distinct P, got %d", 132, len(dp))
	}

	do := g.DistinctO()
	if len(do) != 515 {
		t.Errorf("want %d distinct O, got %d", 515, len(do))
	}
}

func TestSmallSPattern(t *testing.T) {
	g, err := small()
	if err != nil {
		t.Fatal(err)
	}

	vv := g.VirtualValue(nsEst + `100012`)
	if !vv.Known {
		t.Errorf("%q is not known in the loaded graph", vv.Value)
	}

	tp := &graph.TriplePattern{S: vv.Pattern()}
	ms := g.Match(tp)
	if len(ms) != 132 {
		t.Log(tp.String(g))
		for _, tr := range ms {
			t.Log(tr.String(g))
		}
		t.Errorf("want %d S-matches, got %d", 132, len(ms))
	}

	found := false
	want := nsSchool + `EstablishmentName`
	for i, tr := range ms {
		got := tr.P.Value(g)
		if got == want {
			found = true
			break
		}
		t.Log(i, "got P:", got)
	}
	if !found {
		t.Errorf("expected to find P %s", want)
	}
}

func TestSmallSPPattern(t *testing.T) {
	g, err := small()
	if err != nil {
		t.Fatal(err)
	}

	vv := g.VirtualValue(nsEst + `100012`)
	if !vv.Known {
		t.Errorf("%q is not known in the loaded graph", vv.Value)
	}

	vv2 := g.VirtualValue(nsSchool + `EstablishmentName`)
	if !vv2.Known {
		t.Errorf("%q is not known in the loaded graph", vv2.Value)
	}

	tp := &graph.TriplePattern{S: vv.Pattern(), P: vv2.Pattern()}
	ms := g.Match(tp)
	if len(ms) != 1 {
		t.Log(tp.String(g))
		for _, tr := range ms {
			t.Log(tr.String(g))
		}
		t.Errorf("want %d SP-match, got %d", 1, len(ms))
	} else {
		tr0 := ms[0].String(g)
		want0 := `[(string:http://education.data.gov.uk/def/school/establishment/100012) (string:http://education.data.gov.uk/def/school/EstablishmentName) (string:Carlton Primary School)]`
		if tr0 != want0 {
			t.Errorf("want %s, got %s", want0, tr0)
		}
	}
}

func TestLoadLargeCSVFile(t *testing.T) {
	ti := time.Now()

	if testing.Short() {
		t.Skip("Not testing in short mode")
	}
	f, err := os.Open(path)
	if err != nil {
		t.Skip(err)
	}
	defer f.Close()

	zr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer zr.Close()

	dec := NewDecoder(zr, nsEst, nsSchool)
	dec.Reindex = false
	g, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	if g.CountValues() != 433796 {
		t.Errorf("want %d values, got %d", 433796, g.CountValues())
	}
	if g.CountTriples() != 6249408 {
		t.Errorf("want %d triples, got %d", 6249408, g.CountTriples())
	}
	PrintMemUsage(t, "dataset only")
	t.Log(time.Now().Sub(ti), "dataset only")

	g.RebuildIndex()
	PrintMemUsage(t, "reindexed")
	t.Log(time.Now().Sub(ti),"reindexed")

	ds, val, idx := g.ByteSizes()
	t.Logf("Dataset %v MiB", bToMb(uint64(ds)))
	t.Logf("Values  %v MiB", bToMb(uint64(val)))
	t.Logf("Indices %v MiB", bToMb(uint64(idx)))



	vv := g.VirtualValue(nsEst + `129216`)
	if !vv.Known {
		t.Errorf("%q is not known in the loaded graph", vv.Value)
	}

	vv2 := g.VirtualValue(nsSchool + `EstablishmentName`)
	if !vv2.Known {
		t.Errorf("%q is not known in the loaded graph", vv2.Value)
	}

	tp := &graph.TriplePattern{S: vv.Pattern(), P: vv2.Pattern()}
	ms := g.Match(tp)
	if len(ms) != 1 {
		t.Log(tp.String(g))
		for _, tr := range ms {
			t.Log(tr.String(g))
		}
		t.Errorf("want %d SP-match, got %d", 1, len(ms))
	} else {
		tr0 := ms[0].String(g)
		want0 := `[(string:http://education.data.gov.uk/def/school/establishment/129216) (string:http://education.data.gov.uk/def/school/EstablishmentName) (string:Yarborough Middle School)]`
		if tr0 != want0 {
			t.Errorf("want %s, got %s", want0, tr0)
		}
	}
	t.Log(time.Now().Sub(ti), "SP pattern")
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// See https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage(t *testing.T, kind string) {
	t.Helper()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	t.Logf("Alloc %v MiB, %s", bToMb(m.Alloc), kind)
	t.Logf("\tTotalAlloc %v MiB", bToMb(m.TotalAlloc))
	t.Logf("\tSys        %v MiB", bToMb(m.Sys))
	t.Logf("\tNumGC      %v", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
