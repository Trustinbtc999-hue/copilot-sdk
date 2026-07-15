import asyncio

from copilot import CopilotClient
from copilot.session import PermissionHandler
from copilot.session_events import AssistantMessageData

# Demonstrate using an MCP server (filesystem) with the Copilot SDK.
# The @modelcontextprotocol/server-filesystem package must be available via npx.


async def main():
    client = CopilotClient()
    await client.start()

    session = await client.create_session(
        on_permission_request=PermissionHandler.approve_all,
        mcp_servers={
            "filesystem": {
                "command": "npx",
                "args": ["-y", "@modelcontextprotocol/server-filesystem", "/tmp"],
                "tools": ["*"],
            }
        },
    )

    print(f"Session created: {session.session_id}")

    result = await session.send_and_wait("List the files in the allowed directory")
    content = None
    if result:
        match result.data:
            case AssistantMessageData() as data:
                content = data.content
    print(f"Response: {content}")

    await client.stop()


if __name__ == "__main__":
    asyncio.run(main())
