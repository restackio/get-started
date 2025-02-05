import OpenAI from "openai/index";
import "dotenv/config";

let openaiInstance: OpenAI | null = null;

export const openaiClient = ({
  apiKey = process.env.RESTACK_API_KEY,
}: {
  apiKey?: string;
}): OpenAI => {
  if (!apiKey) {
    throw new Error("API key is required to create OpenAI client.");
  }

  if (!openaiInstance) {
    openaiInstance = new OpenAI({
      baseURL: "https://ai.restack.io",
      apiKey,
    });
  }
  return openaiInstance;
};
