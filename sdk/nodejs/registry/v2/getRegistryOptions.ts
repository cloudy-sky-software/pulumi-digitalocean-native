// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getRegistryOptions(args?: GetRegistryOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryOptionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:registry/v2:getRegistryOptions", {
    }, opts);
}

export interface GetRegistryOptionsArgs {
}

export interface GetRegistryOptionsResult {
    readonly items: outputs.registry.v2.GetRegistryOptionsProperties;
}
export function getRegistryOptionsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryOptionsResult> {
    return pulumi.output(getRegistryOptions(opts))
}