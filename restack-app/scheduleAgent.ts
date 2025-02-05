import { client } from "./src/client";
import { agentChatTool } from "./src/agents/agent";
export type InputSchedule = {
  name: string;
};

async function scheduleAgent(input: InputSchedule) {
  try {
    const agentId = `${Date.now()}-${agentChatTool.name}`;
    const runId = await client.scheduleAgent({
      agentName: agentChatTool.name,
      agentId,
      input,
    });

    const result = await client.getAgentResult({ agentId, runId });

    console.log("Agent result:", result);

    process.exit(0); // Exit the process successfully
  } catch (error) {
    console.error("Error scheduling agent:", error);
    process.exit(1); // Exit the process with an error code
  }
}

scheduleAgent({
  name: "test",
});
