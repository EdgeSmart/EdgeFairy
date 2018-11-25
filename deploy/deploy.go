package deploy

import (
	"github.com/EdgeSmart/EdgeFairy/user"
)

func Run() error {
	user.Login()
	return nil
}
