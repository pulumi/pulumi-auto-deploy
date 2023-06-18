# Pulumi Auto Deploy

[![Slack](https://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fauto-deploy.svg)](https://www.npmjs.com/package/@pulumi/auto-deploy)
[![Python version](https://badge.fury.io/py/pulumi-auto-deploy.svg)](https://pypi.org/project/pulumi-auto-deploy)
[![NuGet version](https://badge.fury.io/nu/pulumi.auto-deploy.svg)](https://badge.fury.io/nu/pulumi.auto-deploy)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-auto-deploy/sdk/go/auto-deploy)](https://pkg.go.dev/github.com/pulumi/pulumi-auto-deploy/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fauto-deploy.svg)](https://github.com/pulumi/pulumi-auto-deploy/blob/main/LICENSE)

A Pulumi Component for configuring automated updates of dependent stacks using [Pulumi Deployments](https://www.pulumi.com/docs/pulumi-cloud/deployments/). It lets you simply express dependencies between stacks, and takes care of creating and updating the necessary Deployment Webhooks under the hood. Each stack that you configure must have [Deployment Settings](https://www.pulumi.com/docs/pulumi-cloud/deployments/reference/#deployment-settings).

```ts
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
    downstreamRefs: [d.ref, e.ref, f.ref],

});

export const a = new autodeploy.AutoDeployer("auto-deployer-a", {
    organization,
    project,
    stack: "a",
    downstreamRefs: [b.ref, c.ref],
});

```

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/auto-deploy

or `yarn`:

    $ yarn add @pulumi/auto-deploy

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-auto-deploy

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-auto-deploy/sdk/go

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.AutoDeploy