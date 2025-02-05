import { zodFunction } from "openai/helpers/zod";
import { executeTodoWorkflow } from "../workflows/executeTodo";
import { createTodo } from "./createTodo";
import { CreateTodoInput, ExecuteTodoInput } from "./toolTypes";

export const getTools = async () => {
  const tools = [
    zodFunction({
      name: createTodo.name,
      parameters: CreateTodoInput,
    }),
    zodFunction({
      name: executeTodoWorkflow.name,
      parameters: ExecuteTodoInput,
    }),
  ];
  return tools;
};
