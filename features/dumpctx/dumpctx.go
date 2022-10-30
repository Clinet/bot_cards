package dumpctx

import (
	"github.com/Clinet/clinet/cmds"
)

var CmdRoot *cmds.Cmd

func init() {
	CmdRoot = cmds.NewCmd("dumpctx", "DEBUG: Dumps command context as built by Clinet", nil)
	CmdRoot.AddSubCmds(
		cmds.NewCmd("sub1", "Test subcommand 1", handleDumpCtx),
		cmds.NewCmd("sub2", "Test subcommand 2", handleDumpCtx),
	)
}

func handleDumpCtx(ctx *cmds.CmdCtx) *cmds.CmdResp {
	return cmds.NewCmdRespEmbed("Dump of ctx (*cmds.CmdCtx)", "```JSON\n" + ctx.String() + "```")
}