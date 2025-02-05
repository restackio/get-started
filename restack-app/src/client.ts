import Restack from "@restackio/ai";

import "dotenv/config";

export const connectionOptions = {
  engineId: process.env.RESTACK_ENGINE_ID!,
  address: process.env.RESTACK_ENGINE_ADDRESS!,
  apiKey: process.env.RESTACK_ENGINE_API_KEY!,
  apiAddress: process.env.RESTACK_ENGINE_API_ADDRESS!,
};

export const client = new Restack(
  process.env.RESTACK_ENGINE_API_KEY ? connectionOptions : undefined
);
