package main

var defaultTemplate = `
[cmd.check]
  macros = ["varcheck","ineffassign","golint","errcheck","dupl","structcheck","aligncheck","vet","vetshadow","gocyclo","deadcode", "gofmt"]

[cmd.test]
  macros = ["test"]

[cmd.fix]
  macros = ["gofmt-fix"]
  run-next = ["check"]

[vars]
  duplthreshold = 75
  min_confidence = 0.8
  ignoreDirs = ["^vendor$", "^Godeps$", "^\\..+$"]
  default = "fix"
  buildfileName = "gobuild.toml"
  parallelBuildCount = 16
  stop_loading_parent = [".git"]
  gocyclo_over = 10

[macro.aligncheck]
  cmd="aligncheck"
  goget="github.com/opennota/check/cmd/aligncheck"
  args=["."]
  if-files=[".*\\.go"]

[macro.test]
  cmd="go"
  args=["test", "."]
  if-files=[".*\\.go"]

[macro.deadcode]
  cmd="deadcode"
  goget="github.com/remyoudompheng/go-misc/deadcode"
  args=["."]
  if-files=[".*\\.go"]

[macro.dupl]
  cmd="dupl"
  goget="github.com/mibk/dupl"
  args=["-plumbing", "-threshold", "{duplthreshold}"]
  if-files=[".*\\.go"]
  append-files=true
  cross-directory=true

[macro.errcheck]
  cmd="errcheck"
  goget="github.com/kisielk/errcheck"
  args=[]
  if-files=[".*\\.go"]
  append-files=true

[macro.gocyclo]
  cmd="gocyclo"
  goget="github.com/alecthomas/gocyclo"
  args=["-over", "{gocyclo_over}"]
  append-files=true
  if-files=[".*\\.go"]

[macro.golint]
  cmd="golint"
  goget="github.com/golang/lint/golint"
  args=["-min_confidence", "{min_confidence}"]
  if-files=[".*\\.go"]
  append-files=true

[macro.ineffassign]
  cmd="ineffassign"
  goget="github.com/gordonklaus/ineffassign"
  args=["-n", "."]
  if-files=[".*\\.go"]

[macro.structcheck]
  cmd="structcheck"
  goget="github.com/opennota/check/cmd/structcheck"
  args=["."]
  if-files=[".*\\.go"]

[macro.varcheck]
  cmd="varcheck"
  goget="github.com/opennota/check/cmd/varcheck"
  args=[]
  if-files=[".*\\.go"]
  append-files=true

[macro.vet]
  cmd="go"
  args=["tool", "vet"]
  if-files=[".*\\.go"]
  append-files=true

[macro.vetshadow]
  cmd="go"
  args=["tool", "vet", "--shadow"]
  if-files=[".*\\.go"]
  append-files=true

[macro.gofmt]
  cmd="gofmt"
  args=["-s", "-l"]
  if-files=[".*\\.go"]
  append-files=true

[macro.goimports]
  cmd="goimports"
  args=["-l"]
  if-files=[".*\\.go"]
  goget="golang.org/x/tools/cmd/goimports"
  append-files=true

[macro.gofmt-fix]
  cmd="gofmt"
  args=["-s", "-w"]
  if-files=[".*\\.go"]
  append-files=true
`
