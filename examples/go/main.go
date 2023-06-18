package main

import (
	"github.com/pulumi/pulumi-auto-deploy/sdk/go/autodeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		projectVar := "dependency-example"
		organization := "pulumi"
		f, err := autodeploy.NewAutoDeployer(ctx, "f", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("f"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		e, err := autodeploy.NewAutoDeployer(ctx, "e", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("e"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		d, err := autodeploy.NewAutoDeployer(ctx, "d", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("d"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		c, err := autodeploy.NewAutoDeployer(ctx, "c", &autodeploy.AutoDeployerArgs{
			Organization:   pulumi.String(organization),
			Project:        pulumi.String(projectVar),
			Stack:          pulumi.String("c"),
			DownstreamRefs: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		b, err := autodeploy.NewAutoDeployer(ctx, "b", &autodeploy.AutoDeployerArgs{
			Organization: pulumi.String(organization),
			Project:      pulumi.String(projectVar),
			Stack:        pulumi.String("b"),
			DownstreamRefs: pulumi.StringArray{
				d.Ref,
				e.Ref,
				f.Ref,
			},
		})
		if err != nil {
			return err
		}
		_, err = autodeploy.NewAutoDeployer(ctx, "a", &autodeploy.AutoDeployerArgs{
			Organization: pulumi.String(organization),
			Project:      pulumi.String(projectVar),
			Stack:        pulumi.String("a"),
			DownstreamRefs: pulumi.StringArray{
				b.Ref,
				c.Ref,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
