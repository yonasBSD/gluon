package command

import (
	"bytes"
	"github.com/ProtonMail/gluon/rfcparser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_UnselectCommand(t *testing.T) {
	input := toIMAPLine(`tag UNSELECT`)
	s := rfcparser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &Unselect{}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "unselect", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}
