// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listAppsRegions(args?: ListAppsRegionsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.apps.v2.AppsListRegionsResponse> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:listAppsRegions", {
    }, opts);
}

export interface ListAppsRegionsArgs {
}
export function listAppsRegionsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.apps.v2.AppsListRegionsResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:apps/v2:listAppsRegions", {
    }, opts);
}

