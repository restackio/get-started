import { log, step } from "@restackio/ai/workflow";
import * as functions from "../functions";

type Input = {
  name: string;
};

type Output = {
  messages: {
    greetMessage: string;
  };
};

export async function greetingWorkflow({ name }: Input): Promise<Output> {
  // Step 1: Create hello message with simple function
  const { message: greetMessage } = await step<typeof functions>({}).hello({
    name,
  });

  log.info("Hello", { greetMessage });

  return {
    messages: { greetMessage },
  };
}
