import pulumi
import pulumi_auto_deploy as auto_deploy

'''

The following example configures automatic deployment of stacks with the following dependency graph:
    a
    ├── b
    │   ├── d
    │   ├── e
    │   └── f
    └── c
Whenever a node in the graph is updated, 
all downstream nodes will be automatically updated via a webhook triggering Pulumi Deployments.
'''

project_var = "dependency-example"
organization = pulumi.get_organization()
f = auto_deploy.AutoDeployer("f",
    organization=organization,
    project=project_var,
    stack="f",
    downstream_refs=[])
e = auto_deploy.AutoDeployer("e",
    organization=organization,
    project=project_var,
    stack="e",
    downstream_refs=[])
d = auto_deploy.AutoDeployer("d",
    organization=organization,
    project=project_var,
    stack="d",
    downstream_refs=[])
c = auto_deploy.AutoDeployer("c",
    organization=organization,
    project=project_var,
    stack="c",
    downstream_refs=[])
b = auto_deploy.AutoDeployer("b",
    organization=organization,
    project=project_var,
    stack="b",
    downstream_refs=[
        d.ref,
        e.ref,
        f.ref,
    ])
a = auto_deploy.AutoDeployer("a",
    organization=organization,
    project=project_var,
    stack="a",
    downstream_refs=[
        b.ref,
        c.ref,
    ])
