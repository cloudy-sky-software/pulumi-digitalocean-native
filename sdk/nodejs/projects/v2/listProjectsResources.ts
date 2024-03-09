// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listProjectsResources(args: ListProjectsResourcesArgs, opts?: pulumi.InvokeOptions): Promise<ListProjectsResourcesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:projects/v2:listProjectsResources", {
        "projectId": args.projectId,
    }, opts);
}

export interface ListProjectsResourcesArgs {
    /**
     * A unique identifier for a project.
     */
    projectId: string;
}

export interface ListProjectsResourcesResult {
    readonly items: outputs.projects.v2.ListProjectsResources;
}
export function listProjectsResourcesOutput(args: ListProjectsResourcesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListProjectsResourcesResult> {
    return pulumi.output(args).apply((a: any) => listProjectsResources(a, opts))
}

export interface ListProjectsResourcesOutputArgs {
    /**
     * A unique identifier for a project.
     */
    projectId: pulumi.Input<string>;
}