package command

import (
	"bytes"
	"github.com/ProtonMail/gluon/imap/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_DoneCommand(t *testing.T) {
	input := toIMAPLine(`DONE`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "", Payload: &DoneCommand{}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "done", p.LastParsedCommand())
	require.Empty(t, p.LastParsedTag())
}

func TestParser_DoneCommandAfterTagIsError(t *testing.T) {
	input := toIMAPLine(`tag DONE`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)
	_, err := p.Parse()
	require.Error(t, err)
	require.Equal(t, "tag", p.LastParsedTag())
}
