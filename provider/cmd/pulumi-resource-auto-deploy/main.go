// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	pcloud "github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	p.RunProvider("auto-deploy", Version,
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Components: []infer.InferredComponent{
				infer.Component[*AutoDeployer, AutoDeployerArgs, *AutoDeployerOutput](),
			},
		}))
}

type AutoDeployer struct{}
type AutoDeployerArgs struct {
	Organization   pulumi.StringInput   `pulumi:"organization"`
	Project        pulumi.StringInput   `pulumi:"project"`
	Stack          pulumi.StringInput   `pulumi:"stack"`
	DownstreamRefs []pulumi.StringInput `pulumi:"downstreamRefs"`
}

type AutoDeployerOutput struct {
	pulumi.ResourceState
	Organization pulumi.StringInput `pulumi:"organization"`
	Project      pulumi.StringInput `pulumi:"project"`
	Stack        pulumi.StringInput `pulumi:"stack"`
	// Outputs
	DownstreamRefs     []pulumi.StringInput     `pulumi:"downstreamRefs"`
	Ref                pulumi.StringOutput      `pulumi:"ref"`
	DownstreamWebhooks pulumi.StringArrayOutput `pulumi:"downstreamWebhooks"`
}

func (r *AutoDeployer) Construct(ctx *pulumi.Context, name, typ string, args AutoDeployerArgs, opts pulumi.ResourceOption) (*AutoDeployerOutput, error) {
	comp := &AutoDeployerOutput{}
	err := ctx.RegisterComponentResource(typ, name, comp, opts)
	if err != nil {
		return nil, err
	}

	downstreamWebhooks := []pulumi.StringOutput{}

	for i, d := range args.DownstreamRefs {
		wh, err := pcloud.NewWebhook(ctx, fmt.Sprintf("%s-%d", name, i), &pcloud.WebhookArgs{
			OrganizationName: args.Organization,
			ProjectName:      args.Project,
			StackName:        args.Stack,
			Format:           pcloud.WebhookFormatPulumiDeployments,
			PayloadUrl:       d,
			Active:           pulumi.Bool(true),
			DisplayName:      d,
			Filters:          pcloud.WebhookFiltersArray{pcloud.WebhookFiltersUpdateSucceeded},
		}, pulumi.Parent(comp))
		if err != nil {
			return nil, err
		}
		downstreamWebhooks = append(downstreamWebhooks, wh.Name.Elem().ToStringOutput())
	}

	comp.Organization = args.Organization
	comp.Project = args.Project
	comp.Stack = args.Stack
	comp.DownstreamRefs = args.DownstreamRefs
	comp.Ref = pulumi.Sprintf("%s/%s", args.Project, args.Stack)
	comp.DownstreamWebhooks = pulumi.ToStringArrayOutput(downstreamWebhooks)

	return comp, nil
}

// Implementing Annotate lets you provide descriptions and default values for resources and they will
// be visible in the provider's schema and the generated SDKs.
func (c *AutoDeployer) Annotate(a infer.Annotator) {
	a.Describe(&c, "Automatically trigger downstream updates on dependent stacks via Pulumi Deployments.\n"+
		"AutoDeployer requires that stacks have Deployment Settings configured.")
}

// Annotate lets you provide descriptions and default values for fields and they will
// be visible in the provider's schema and the generated SDKs.
func (c *AutoDeployerArgs) Annotate(a infer.Annotator) {
	a.Describe(&c.Organization, "The organization name for the AutoDeployer stack.")
	a.Describe(&c.Project, "The project name for the AutoDeployer stack.")
	a.Describe(&c.Stack, "The stack name for this AutoDeployer.")
	a.Describe(&c.DownstreamRefs, "A list of `AutoDeployer.DownstreamRef` indicating which stacks should\n"+
		"automatically be updated via Pulumi Deployments when this stack is successfully updated.")
}

// Annotate lets you provide descriptions and default values for fields and they will
// be visible in the provider's schema and the generated SDKs.
func (c *AutoDeployerOutput) Annotate(a infer.Annotator) {
	a.Describe(&c.Organization, "The organization name for the AutoDeployer stack.")
	a.Describe(&c.Project, "The project name for the AutoDeployer stack.")
	a.Describe(&c.Stack, "The stack name for this AutoDeployer.")
	a.Describe(&c.DownstreamRefs, "A list of `AutoDeployer.DownstreamRef` indicating which stacks should\n"+
		"automatically be updated via Pulumi Deployments when this stack is successfully updated.")

	a.Describe(&c.Ref, "The output reference that can be passed to another AutoDeployer's downstreamRefs list\n"+
		"to configure depedent updates.")
	a.Describe(&c.DownstreamWebhooks, "A list of webhook URLs configured on this stack to trigger downstream deployments.")
}
