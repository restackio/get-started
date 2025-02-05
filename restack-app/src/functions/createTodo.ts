import { log } from "@restackio/ai/function";

export const createTodo = async ({ todoTitle }: { todoTitle: string }) => {
  const todo_id = `todo-${Math.floor(Math.random() * 10000)}`;
  log.info("createTodo", { todo_id, todoTitle });
  return `Created the todo '${todoTitle}' with id: ${todo_id}`;
};
