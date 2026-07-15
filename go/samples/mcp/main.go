package main

import (
	"context"
	"fmt"

	copilot "github.com/github/copilot-sdk/go"
)

// Demonstrate using an MCP server (filesystem) with the Copilot SDK.
// The @modelcontextprotocol/server-filesystem package must be available via npx.
//
// Set the COPILOT_CLI_PATH environment variable if the Copilot CLI is not in PATH.

func main() {
	ctx := context.Background()
	client := copilot.NewClient(nil)
	if err := client.Start(ctx); err != nil {
		panic(err)
	}
	defer client.Stop()

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
		MCPServers: map[string]copilot.MCPServerConfig{
			"filesystem": copilot.MCPStdioServerConfig{
				Command: "npx",
				Args:    []string{"-y", "@modelcontextprotocol/server-filesystem", "/tmp"},
				Tools:   []string{"*"},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer session.Disconnect()

	fmt.Printf("Session created: %s\n", session.SessionID)

	reply, err := session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: "List the files in the allowed directory",
	})
	if err != nil {
		panic(err)
	}

	content := ""
	if reply != nil {
		if d, ok := reply.Data.(*copilot.AssistantMessageData); ok {
			content = d.Content
		}
	}
	fmt.Printf("Response: %s\n", content)
}

