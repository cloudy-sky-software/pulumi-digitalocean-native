// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getRegistry(args?: GetRegistryArgs, opts?: pulumi.InvokeOptions): Promise<outputs.registry.v2.GetRegistryProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:registry/v2:getRegistry", {
    }, opts);
}

export interface GetRegistryArgs {
}
export function getRegistryOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.registry.v2.GetRegistryProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:registry/v2:getRegistry", {
    }, opts);
}

