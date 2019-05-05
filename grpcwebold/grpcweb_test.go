package grpcweb

import (
	"testing"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func parseProto(t *testing.T, fname string) []*desc.FileDescriptor {
	t.Helper()

	p := &protoparse.Parser{
		ImportPaths: []string{"testdata"},
	}
	d, err := p.ParseFiles(fname)
	require.NoError(t, err)
	return d
}
