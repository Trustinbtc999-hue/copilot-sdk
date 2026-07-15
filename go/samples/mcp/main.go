package main

import (
	"context"
	"fmt"
	"path/filepath"

	copilot "github.com/github/copilot-sdk/go"
)

// Demonstrate using an MCP server (filesystem) with the Copilot SDK.
// The @modelcontextprotocol/server-filesystem package must be available via npx.

func main() {
	ctx := context.Background()
	cliPath := filepath.Join("..", "..", "..", "nodejs", "node_modules", "@github", "copilot", "index.js")
	client := copilot.NewClient(&copilot.ClientOptions{Connection: copilot.StdioConnection{Path: cliPath}})
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

	reply, _ := session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: "List the files in the allowed directory",
	})

	content := ""
	if reply != nil {
		if d, ok := reply.Data.(*copilot.AssistantMessageData); ok {
			content = d.Content
		}
	}
	fmt.Printf("Response: %s\n", content)
}
