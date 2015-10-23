package main

var defaultTemplate = `
[vars]
  ignoreDirs = [".git", "Godeps", "vendor"]
  stopLoadingParent = [".git"]
  buildFlags = ["."]
  artifactsEnv = "CIRCLE_ARTIFACTS"
  testReportEnv = "CIRCLE_TEST_REPORTS"

[gotestcoverage]
  timeout = "10s"
  cpu = "4"
  parallel = 8
  race = true
  covermode = "atomic"

[install]
  [install.goget]
    gometalinter = "github.com/alecthomas/gometalinter"
    golint = "github.com/golang/lint/golint"
    go-junit-report = "github.com/jstemmer/go-junit-report"
    goimports = "golang.org/x/tools/cmd/goimports"
    gocyclo = "github.com/alecthomas/gocyclo"
    aligncheck = "github.com/opennota/check/cmd/aligncheck"
    varcheck = "github.com/opennota/check/cmd/varcheck"
    dupl = "github.com/mibk/dupl"
`
