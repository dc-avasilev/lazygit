package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var CommitMultiline = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Commit with a multi-line commit message",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFile("myfile", "myfile content")
	},
	Run: func(shell *Shell, input *Input, assert *Assert, keys config.KeybindingConfig) {
		assert.CommitCount(0)

		input.PrimaryAction()
		input.PressKeys(keys.Files.CommitChanges)

		input.Type("first line")
		input.PressKeys(keys.Universal.AppendNewline)
		input.PressKeys(keys.Universal.AppendNewline)
		input.Type("third line")
		input.Confirm()

		assert.CommitCount(1)
		assert.MatchHeadCommitMessage(Equals("first line"))

		input.SwitchToCommitsWindow()
		assert.MatchMainViewContent(MatchesRegexp("first line\n\\s*\n\\s*third line"))
	},
})