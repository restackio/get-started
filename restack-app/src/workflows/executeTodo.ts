import { log, sleep, step } from "@restackio/ai/workflow";
import * as functions from "../functions";

type Input = {
  todoTitle: string;
  todoId: string;
};

type Output = {
  todoId: string;
  todoTitle: string;
  details: string;
  status: string;
};

export async function executeTodoWorkflow({
  todoTitle,
  todoId,
}: Input): Promise<Output> {
  const random = await step<typeof functions>({}).getRandom({
    todoTitle,
  });

  await sleep(2000);

  const result = await step<typeof functions>({}).getResult({
    todoTitle,
    todoId,
  });

  const todoDetails = {
    todoId,
    todoTitle,
    details: random,
    status: result.status,
  };

  log.info("Todo Details", { todoDetails });

  return todoDetails;
}
