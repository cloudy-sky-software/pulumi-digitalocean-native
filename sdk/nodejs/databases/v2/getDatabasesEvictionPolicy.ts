// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabasesEvictionPolicy(args: GetDatabasesEvictionPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesEvictionPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabasesEvictionPolicy", {
        "databaseClusterUuid": args.databaseClusterUuid,
    }, opts);
}

export interface GetDatabasesEvictionPolicyArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
}

export interface GetDatabasesEvictionPolicyResult {
    readonly items: outputs.databases.v2.GetDatabasesEvictionPolicyProperties;
}
export function getDatabasesEvictionPolicyOutput(args: GetDatabasesEvictionPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesEvictionPolicyResult> {
    return pulumi.output(args).apply((a: any) => getDatabasesEvictionPolicy(a, opts))
}

export interface GetDatabasesEvictionPolicyOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
}
