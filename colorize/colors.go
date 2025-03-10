package colorize

import c "github.com/fatih/color"

var InfoFgBold = c.New(c.FgCyan, c.Bold).SprintFunc()
var WarningFgBold = c.New(c.FgYellow, c.Bold).SprintFunc()
var ErrorFgBold = c.New(c.FgRed, c.Bold).SprintFunc()
var SuccessFgBold = c.New(c.FgGreen, c.Bold).SprintFunc()

var InfoFg = c.New(c.FgCyan).SprintFunc()
var WarningFg = c.New(c.FgYellow).SprintFunc()
var ErrorFg = c.New(c.FgRed).SprintFunc()
var SuccessFg = c.New(c.FgGreen).SprintFunc()
