import { CopilotClient, approveAll } from "@github/copilot-sdk";

// Demonstrate using an MCP server (filesystem) with the Copilot SDK.
// The @modelcontextprotocol/server-filesystem package must be available via npx.

async function main() {
    const client = new CopilotClient();

    const session = await client.createSession({
        onPermissionRequest: approveAll,
        mcpServers: {
            filesystem: {
                type: "local",
                command: "npx",
                args: ["-y", "@modelcontextprotocol/server-filesystem", "/tmp"],
                tools: ["*"],
            },
        },
    });

    console.log(`Session created: ${session.sessionId}`);

    const result = await session.sendAndWait({
        prompt: "List the files in the allowed directory",
    });

    console.log(`Response: ${result?.data.content}`);

    await session.disconnect();
    await client.stop();
}

main().catch(console.error);
