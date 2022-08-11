package playground

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()

	pflag.Parse()
	opts.Paths = pflag.Args()

	os.Exit(m.Run())
}

func TestFeatures(t *testing.T) {
	tsi := func(ctx *godog.TestSuiteContext) {
	}
	si := func(ctx *godog.ScenarioContext) {
		ctx.Step(`^I eat (\d+)$`, func() error { return godog.ErrPending })
		ctx.Step(`^there are (\d+) godogs$`, func() error { return godog.ErrPending })
		ctx.Step(`^there should be (\d+) remaining$`, func() error { return godog.ErrPending })
	}

	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                 "godogs",
		Options:              &o,
		TestSuiteInitializer: tsi,
		ScenarioInitializer:  si,
	}.Run()

	assert.Equal(t, 0, status)
}
