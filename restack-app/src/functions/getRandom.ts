import { log } from "@restackio/ai/function";

export const getRandom = async ({ todoTitle }: { todoTitle: string }) => {
  const random = Math.random() * 100;
  log.info("getRandom", { todoTitle, random });
  return `The random number for ${todoTitle} is ${random}`;
};
