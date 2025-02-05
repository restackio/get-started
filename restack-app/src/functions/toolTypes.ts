import { z } from "zod";

export const CreateTodoInput = z.object({
  todoTitle: z.string(),
});

export type CreateTodoInputType = z.infer<typeof CreateTodoInput>;

export const ExecuteTodoInput = z.object({
  todoId: z.string(),
  todoTitle: z.string(),
});

export type ExecuteTodoInputType = z.infer<typeof ExecuteTodoInput>;
