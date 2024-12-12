// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listProjects(args?: ListProjectsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.projects.v2.ListProjectsItems> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:projects/v2:listProjects", {
    }, opts);
}

export interface ListProjectsArgs {
}
export function listProjectsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.projects.v2.ListProjectsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:projects/v2:listProjects", {
    }, opts);
}

