package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

// PostBotDM posts a DM as the cloud bot user.
func (p *Plugin) PostBotDM(userID string, message string) {
}

// PostBotCustomDM posts a DM as the cloud bot user using custom post with action buttons.
func (p *Plugin) PostBotCustomDM(userID string, message string, todo string, issueID string) {
}

func (p *Plugin) createBotPostDM(post *model.Post, userID string) {
}

// ReplyPostBot post a message and a todo in the same thread as the post postID
func (p *Plugin) ReplyPostBot(postID, message, todo string) error {
	return nil
}
