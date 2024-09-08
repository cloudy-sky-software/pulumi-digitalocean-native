// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDatabasesClusters(args?: ListDatabasesClustersArgs, opts?: pulumi.InvokeOptions): Promise<outputs.databases.v2.ListDatabasesClustersProperties> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:listDatabasesClusters", {
    }, opts);
}

export interface ListDatabasesClustersArgs {
}
export function listDatabasesClustersOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.databases.v2.ListDatabasesClustersProperties> {
    return pulumi.output(listDatabasesClusters(opts))
}
