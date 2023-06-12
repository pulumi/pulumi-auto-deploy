import * as autodeploy from "@pulumi/auto-deploy";

const random = new autodeploy.Random("my-random", { length: 24 });

export const output = random.result;