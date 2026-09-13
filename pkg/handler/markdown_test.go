package handler

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	slackGoUtil "github.com/takara2314/slack-go-util"
)

// TestUnitMarkdownKeepsSingleNewlines pins the behaviour the conversations_add_message
// tool description promises for content_type=text/markdown: a single newline inside a
// paragraph survives as a line break. slack-go-util < v0.4.1 dropped soft line breaks,
// so "line one\nline two" was posted as "line oneline two"; this fails on a downgrade.
func TestUnitMarkdownKeepsSingleNewlines(t *testing.T) {
	blocks, err := slackGoUtil.ConvertMarkdownTextToBlocks(
		"Hello <@U0123ABCD>\nWe received a reversal request.\n• Transaction ID: 985d\n• Amount: 990.10 GHS\n\nThanks")
	require.NoError(t, err)
	require.Len(t, blocks, 2)

	section, ok := blocks[0].(*slack.SectionBlock)
	require.Truef(t, ok, "first block should be a section, got %T", blocks[0])
	assert.Equal(t, slack.MarkdownType, section.Text.Type)
	assert.Equal(t, "Hello <@U0123ABCD>\nWe received a reversal request.\n• Transaction ID: 985d\n• Amount: 990.10 GHS", section.Text.Text)
}
