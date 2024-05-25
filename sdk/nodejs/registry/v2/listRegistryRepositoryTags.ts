// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listRegistryRepositoryTags(args: ListRegistryRepositoryTagsArgs, opts?: pulumi.InvokeOptions): Promise<ListRegistryRepositoryTagsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:registry/v2:listRegistryRepositoryTags", {
        "registryName": args.registryName,
        "repositoryName": args.repositoryName,
    }, opts);
}

export interface ListRegistryRepositoryTagsArgs {
    /**
     * The name of a container registry.
     */
    registryName: string;
    /**
     * The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
     */
    repositoryName: string;
}

export interface ListRegistryRepositoryTagsResult {
    readonly items: outputs.registry.v2.ListRegistryRepositoryTagsItems;
}
export function listRegistryRepositoryTagsOutput(args: ListRegistryRepositoryTagsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListRegistryRepositoryTagsResult> {
    return pulumi.output(args).apply((a: any) => listRegistryRepositoryTags(a, opts))
}

export interface ListRegistryRepositoryTagsOutputArgs {
    /**
     * The name of a container registry.
     */
    registryName: pulumi.Input<string>;
    /**
     * The name of a container registry repository. If the name contains `/` characters, they must be URL-encoded, e.g. `%2F`.
     */
    repositoryName: pulumi.Input<string>;
}
