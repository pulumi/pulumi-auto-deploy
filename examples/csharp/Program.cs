using System.Collections.Generic;
using System.Linq;
using Pulumi;
using AutoDeploy = Pulumi.AutoDeploy;

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

return await Deployment.RunAsync(() => 
{
    var projectVar = "dependency-example";

    var organization = "pulumi";

    var f = new AutoDeploy.AutoDeployer("f", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "f",
        DownstreamRefs = new[] {},
    });

    var e = new AutoDeploy.AutoDeployer("e", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "e",
        DownstreamRefs = new[] {},
    });

    var d = new AutoDeploy.AutoDeployer("d", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "d",
        DownstreamRefs = new[] {},
    });

    var c = new AutoDeploy.AutoDeployer("c", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "c",
        DownstreamRefs = new[] {},
    });

    var b = new AutoDeploy.AutoDeployer("b", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "b",
        DownstreamRefs = new[]
        {
            d.Ref,
            e.Ref,
            f.Ref,
        },
    });

    var a = new AutoDeploy.AutoDeployer("a", new()
    {
        Organization = organization,
        Project = projectVar,
        Stack = "a",
        DownstreamRefs = new[]
        {
            b.Ref,
            c.Ref,
        },
    });

});

