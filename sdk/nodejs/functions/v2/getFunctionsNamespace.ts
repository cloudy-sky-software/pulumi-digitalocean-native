// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getFunctionsNamespace(args: GetFunctionsNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<outputs.functions.v2.GetFunctionsNamespaceProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:functions/v2:getFunctionsNamespace", {
        "namespaceId": args.namespaceId,
    }, opts);
}

export interface GetFunctionsNamespaceArgs {
    /**
     * The ID of the namespace to be managed.
     */
    namespaceId: string;
}
export function getFunctionsNamespaceOutput(args: GetFunctionsNamespaceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.functions.v2.GetFunctionsNamespaceProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:functions/v2:getFunctionsNamespace", {
        "namespaceId": args.namespaceId,
    }, opts);
}

export interface GetFunctionsNamespaceOutputArgs {
    /**
     * The ID of the namespace to be managed.
     */
    namespaceId: pulumi.Input<string>;
}
