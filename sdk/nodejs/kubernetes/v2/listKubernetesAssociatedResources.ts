// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listKubernetesAssociatedResources(args: ListKubernetesAssociatedResourcesArgs, opts?: pulumi.InvokeOptions): Promise<ListKubernetesAssociatedResourcesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:kubernetes/v2:listKubernetesAssociatedResources", {
        "clusterId": args.clusterId,
    }, opts);
}

export interface ListKubernetesAssociatedResourcesArgs {
    /**
     * A unique ID that can be used to reference a Kubernetes cluster.
     */
    clusterId: string;
}

export interface ListKubernetesAssociatedResourcesResult {
    readonly items: outputs.kubernetes.v2.AssociatedKubernetesResources;
}
export function listKubernetesAssociatedResourcesOutput(args: ListKubernetesAssociatedResourcesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListKubernetesAssociatedResourcesResult> {
    return pulumi.output(args).apply((a: any) => listKubernetesAssociatedResources(a, opts))
}

export interface ListKubernetesAssociatedResourcesOutputArgs {
    /**
     * A unique ID that can be used to reference a Kubernetes cluster.
     */
    clusterId: pulumi.Input<string>;
}
