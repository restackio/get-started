import { FunctionFailure, log } from "@restackio/ai/function";
import {
  ChatCompletionCreateParamsNonStreaming,
  ChatCompletionSystemMessageParam,
  ChatCompletionTool,
  ChatCompletionToolMessageParam,
  ChatCompletionUserMessageParam,
  ChatCompletionAssistantMessageParam,
  ChatCompletionMessage,
} from "openai/resources/chat/completions";

import { openaiClient } from "../utils/client";

export type Message =
  | ChatCompletionSystemMessageParam
  | ChatCompletionUserMessageParam
  | ChatCompletionAssistantMessageParam
  | ChatCompletionToolMessageParam;

export type OpenAIChatInput = {
  systemContent?: string;
  model?: string;
  messages: Message[];
  tools?: ChatCompletionTool[];
};

export const llmChat = async ({
  systemContent = "",
  model = "gpt-4o-mini",
  messages,
  tools,
}: OpenAIChatInput): Promise<ChatCompletionMessage> => {
  try {
    const openai = openaiClient({});

    const chatParams: ChatCompletionCreateParamsNonStreaming = {
      messages: [
        ...(systemContent
          ? [{ role: "system" as const, content: systemContent }]
          : []),
        ...(messages ?? []),
      ],
      model,
      tools,
    };

    log.debug("OpenAI chat completion params", {
      chatParams,
    });

    const completion = await openai.chat.completions.create(chatParams);

    const message = completion.choices[0].message;

    return message;
  } catch (error) {
    throw FunctionFailure.nonRetryable(`Error OpenAI chat: ${error}`);
  }
};
