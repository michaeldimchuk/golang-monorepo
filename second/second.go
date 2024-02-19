package second

import "github.com/michaeldimchuk/golang-monorepo/first"

func Second() string {
	return "second module, version 0: " + first.FirstV2()
}
