// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listFunctionsNamespaces(args?: ListFunctionsNamespacesArgs, opts?: pulumi.InvokeOptions): Promise<outputs.functions.v2.ListFunctionsNamespacesItems> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:functions/v2:listFunctionsNamespaces", {
    }, opts);
}

export interface ListFunctionsNamespacesArgs {
}
export function listFunctionsNamespacesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.functions.v2.ListFunctionsNamespacesItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:functions/v2:listFunctionsNamespaces", {
    }, opts);
}

