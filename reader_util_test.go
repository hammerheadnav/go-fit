package fit_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/cespare/xxhash"
	"github.com/kortschak/utter"
	"github.com/hammerheadnav/go-fit"
)

func fitFingerprint(fit *fit.File) uint64 {
	h := xxhash.New()
	utter.Fdump(h, fit)
	return h.Sum64()
}

func fitUtterDump(fit *fit.File, path string, compressed bool) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	var w io.WriteCloser
	if compressed {
		w = gzip.NewWriter(f)
	} else {
		w = f
	}

	utter.Fdump(w, fit)

	if !compressed {
		return f.Close()
	}

	err = w.Close()
	if err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

// testingDecodeOpts is a convenience type to simplify code generation for
// reader_files_test.go.
type testingDecodeOpts int

const (
	tdoNone testingDecodeOpts = iota
	tdoAllWithDiscardLogger
	tdoStderrLogger
)

func (tdo testingDecodeOpts) opts() []fit.DecodeOption {
	return tdoOpts[tdo]
}

var tdoOpts = [...][]fit.DecodeOption{
	nil,
	{
		fit.WithUnknownFields(),
		fit.WithUnknownMessages(),
		fit.WithLogger(discardLogger()), // For test coverage.
	},
	{
		fit.WithStdLogger(), // For debugging.
	},
}

func (tdo testingDecodeOpts) String() string {
	return tdoString[tdo]
}

var tdoString = [...]string{
	"tdoNone",
	"tdoAllWithDiscardLogger",
	"tdoStderrLogger",
}

func discardLogger() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}

func regenerateDecodeTestTable() error {
	const header = `// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries. 
	
package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
	skipEncode  bool
	encodeNote  string
}{
`

	g := new(gen)
	g.WriteString(header)
	for _, dto := range decodeTestFiles {
		g.openField()
		g.writeStringQField(dto.folder)
		g.writeStringQField(dto.name)
		g.writeBoolField(dto.wantErr)
		g.writeUintField(dto.fingerprint)
		g.writeBoolField(dto.compress)
		g.writeStringField(dto.dopts.String())
		g.writeBoolField(dto.skipEncode)
		g.writeStringQField(dto.encodeNote)
		g.closeField()
	}

	g.WriteString("}\n")

	preFormat := g.Bytes()
	src, err := format.Source(preFormat)
	if err != nil {
		return fmt.Errorf("format.Source: %v on\n%s", err, preFormat)
	}

	return ioutil.WriteFile("reader_files_test.go", src, 0644)
}

type gen struct {
	bytes.Buffer
}

func (g *gen) openField() {
	g.WriteString("{\n")
}

func (g *gen) closeField() {
	g.WriteString("},\n")
}

func (g *gen) writeStringField(s string) {
	g.WriteString(s)
	g.WriteString(",\n")
}

func (g *gen) writeStringQField(s string) {
	g.WriteString(strconv.Quote(s))
	g.WriteString(",\n")
}

func (g *gen) writeBoolField(b bool) {
	g.WriteString(strconv.FormatBool(b))
	g.WriteString(",\n")
}

func (g *gen) writeUintField(u uint64) {
	g.WriteString(strconv.FormatUint(u, 10))
	g.WriteString(",\n")
}
