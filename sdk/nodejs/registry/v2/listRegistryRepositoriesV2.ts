// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listRegistryRepositoriesV2(args: ListRegistryRepositoriesV2Args, opts?: pulumi.InvokeOptions): Promise<ListRegistryRepositoriesV2Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:registry/v2:listRegistryRepositoriesV2", {
        "registryName": args.registryName,
    }, opts);
}

export interface ListRegistryRepositoriesV2Args {
    /**
     * The name of a container registry.
     */
    registryName: string;
}

export interface ListRegistryRepositoriesV2Result {
    readonly items: outputs.registry.v2.ListRegistryRepositoriesV2Items;
}
export function listRegistryRepositoriesV2Output(args: ListRegistryRepositoriesV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListRegistryRepositoriesV2Result> {
    return pulumi.output(args).apply((a: any) => listRegistryRepositoriesV2(a, opts))
}

export interface ListRegistryRepositoriesV2OutputArgs {
    /**
     * The name of a container registry.
     */
    registryName: pulumi.Input<string>;
}
