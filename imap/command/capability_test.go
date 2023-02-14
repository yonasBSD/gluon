package command

import (
	"bytes"
	"github.com/ProtonMail/gluon/imap/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_CapabilityCommand(t *testing.T) {
	input := toIMAPLine(`tag CAPABILITY`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &CapabilityCommand{}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "capability", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}
