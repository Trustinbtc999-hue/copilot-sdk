#:project ../src/GitHub.Copilot.SDK.csproj

using GitHub.Copilot;

// Demonstrate using an MCP server (filesystem) with the Copilot SDK.
// The @modelcontextprotocol/server-filesystem package must be available via npx.

await using var client = new CopilotClient();
await using var session = await client.CreateSessionAsync(new SessionConfig
{
    OnPermissionRequest = PermissionHandler.ApproveAll,
    McpServers = new Dictionary<string, McpServerConfig>
    {
        ["filesystem"] = new McpStdioServerConfig
        {
            Command = "npx",
            Args = ["-y", "@modelcontextprotocol/server-filesystem", "/tmp"],
            Tools = ["*"],
        },
    },
});

Console.WriteLine($"Session created: {session.SessionId}");

var result = await session.SendAndWaitAsync(new MessageOptions
{
    Prompt = "List the files in the allowed directory",
});

Console.WriteLine($"Response: {result?.Data.Content}");
