import { client } from "./src/client";

export type EventInput = {
  agentId: string;
  runId: string;
};

async function eventAgent(input: EventInput) {
  try {
    await client.sendAgentEvent({
      event: {
        name: "message",
        input: { content: "Sales on boots?" },
      },
      agent: {
        agentId: input.agentId,
        runId: input.runId,
      },
    });

    await client.sendAgentEvent({
      event: {
        name: "end",
      },
      agent: {
        agentId: input.agentId,
        runId: input.runId,
      },
    });

    process.exit(0); // Exit the process successfully
  } catch (error) {
    console.error("Error sending event to agent:", error);
    process.exit(1); // Exit the process with an error code
  }
}

eventAgent({
  agentId: "your-agent-id",
  runId: "your-run-id",
});
