package command

import (
	"fmt"
	"github.com/ProtonMail/gluon/imap/parser"
)

type LogoutCommand struct{}

func (l LogoutCommand) String() string {
	return fmt.Sprintf("LOGOUT")
}

func (l LogoutCommand) SanitizedString() string {
	return l.String()
}

type LogoutCommandParser struct{}

func (LogoutCommandParser) FromParser(p *parser.Parser) (Payload, error) {
	return &LogoutCommand{}, nil
}
