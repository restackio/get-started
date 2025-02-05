import { log } from "@restackio/ai/function";

export const getResult = async ({
  todoTitle,
  todoId,
}: {
  todoTitle: string;
  todoId: string;
}) => {
  const statuses = ["completed", "failed"];
  const status = statuses[Math.floor(Math.random() * statuses.length)];
  log.info("getResult", { todoId, todoTitle, status });
  return {
    todoId,
    status,
  };
};
