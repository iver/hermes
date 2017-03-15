package lib

import (
	"bitbucket.com/ivan-iver/config"
)

type Provider interface {
	Init(*config.Config)
	Send(Email) error
}
