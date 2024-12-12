// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDropletsNeighborsIds(args?: ListDropletsNeighborsIdsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.reports.v2.NeighborIds> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:reports/v2:listDropletsNeighborsIds", {
    }, opts);
}

export interface ListDropletsNeighborsIdsArgs {
}
export function listDropletsNeighborsIdsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.reports.v2.NeighborIds> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:reports/v2:listDropletsNeighborsIds", {
    }, opts);
}

