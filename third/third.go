package third

import "github.com/michaeldimchuk/golang-monorepo/second"

func Third() string {
	return "third version 0: " + second.Second()
}
