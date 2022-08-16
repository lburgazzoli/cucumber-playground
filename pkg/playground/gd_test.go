package playground

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)

	flag.Func("godog.paths", "", func(val string) error {
		opts.Paths = strings.SplitAfter(val, ",")
		return nil
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestFeatures(t *testing.T) {
	si := func(ctx *godog.ScenarioContext) {
		ctx.Step(`^I eat (\d+)$`, func() error { return godog.ErrPending })
		ctx.Step(`^there are (\d+) godogs$`, func() error { return godog.ErrPending })
		ctx.Step(`^there should be (\d+) remaining$`, func() error { return godog.ErrPending })
	}

	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                "godogs",
		Options:             &o,
		ScenarioInitializer: si,
	}.Run()

	assert.Equal(t, 0, status)
}
