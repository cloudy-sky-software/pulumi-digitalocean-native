// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listKubernetesNodePools(args: ListKubernetesNodePoolsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.kubernetes.v2.ListKubernetesNodePoolsProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:kubernetes/v2:listKubernetesNodePools", {
        "clusterId": args.clusterId,
    }, opts);
}

export interface ListKubernetesNodePoolsArgs {
    /**
     * A unique ID that can be used to reference a Kubernetes cluster.
     */
    clusterId: string;
}
export function listKubernetesNodePoolsOutput(args: ListKubernetesNodePoolsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.kubernetes.v2.ListKubernetesNodePoolsProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:kubernetes/v2:listKubernetesNodePools", {
        "clusterId": args.clusterId,
    }, opts);
}

export interface ListKubernetesNodePoolsOutputArgs {
    /**
     * A unique ID that can be used to reference a Kubernetes cluster.
     */
    clusterId: pulumi.Input<string>;
}
