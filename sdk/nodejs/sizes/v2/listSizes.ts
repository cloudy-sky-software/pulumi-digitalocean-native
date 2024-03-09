// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listSizes(args?: ListSizesArgs, opts?: pulumi.InvokeOptions): Promise<ListSizesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:sizes/v2:listSizes", {
    }, opts);
}

export interface ListSizesArgs {
}

export interface ListSizesResult {
    readonly items: outputs.sizes.v2.ListSizes;
}
export function listSizesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListSizesResult> {
    return pulumi.output(listSizes(opts))
}