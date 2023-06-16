import * as autodeploy from "@pulumi/auto-deploy";
import * as pulumi from "@pulumi/pulumi";

/**
 *
 * The following example configures automatic deployment of stacks with the following dependency graph:
    a
    ├── b
    │   ├── d
    │   ├── e
    │   └── f
    └── c
 * Whenever a node in the graph is updated, 
 * all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
 */


const organization = pulumi.getOrganization();
const project = "dependency-example"

export const f = new autodeploy.AutoDeployer("auto-deployer-f", {
    organization,
    project,
    stack: "f",
    downstreamRefs: [],
});

export const e = new autodeploy.AutoDeployer("auto-deployer-e", {
    organization,
    project,
    stack: "e",
    downstreamRefs: [],
});

export const d = new autodeploy.AutoDeployer("auto-deployer-d", {
    organization,
    project,
    stack: "d",
    downstreamRefs: [],
});

export const c = new autodeploy.AutoDeployer("auto-deployer-c", {
    organization,
    project,
    stack: "c",
    downstreamRefs: [],
});

export const b = new autodeploy.AutoDeployer("auto-deployer-b", {
    organization,
    project,
    stack: "b",
    downstreamRefs: [d.downstreamRef, e.downstreamRef, f.downstreamRef],

});

export const a = new autodeploy.AutoDeployer("auto-deployer-a", {
    organization,
    project,
    stack: "a",
    downstreamRefs: [b.downstreamRef, c.downstreamRef],
});

