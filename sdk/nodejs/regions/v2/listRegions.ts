// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listRegions(args?: ListRegionsArgs, opts?: pulumi.InvokeOptions): Promise<ListRegionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:regions/v2:listRegions", {
    }, opts);
}

export interface ListRegionsArgs {
}

export interface ListRegionsResult {
    readonly items: outputs.regions.v2.ListRegions;
}
export function listRegionsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListRegionsResult> {
    return pulumi.output(listRegions(opts))
}