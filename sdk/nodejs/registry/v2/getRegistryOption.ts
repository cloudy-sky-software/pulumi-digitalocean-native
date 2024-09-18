// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getRegistryOption(args?: GetRegistryOptionArgs, opts?: pulumi.InvokeOptions): Promise<outputs.registry.v2.GetRegistryOptionProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:registry/v2:getRegistryOption", {
    }, opts);
}

export interface GetRegistryOptionArgs {
}
export function getRegistryOptionOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.registry.v2.GetRegistryOptionProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:registry/v2:getRegistryOption", {
    }, opts);
}

