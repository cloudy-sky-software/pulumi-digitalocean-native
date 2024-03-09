// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabasesReplica(args: GetDatabasesReplicaArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesReplicaResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabasesReplica", {
        "databaseClusterUuid": args.databaseClusterUuid,
        "replicaName": args.replicaName,
    }, opts);
}

export interface GetDatabasesReplicaArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
    /**
     * The name of the database replica.
     */
    replicaName: string;
}

export interface GetDatabasesReplicaResult {
    readonly items: outputs.databases.v2.GetDatabasesReplicaProperties;
}
export function getDatabasesReplicaOutput(args: GetDatabasesReplicaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesReplicaResult> {
    return pulumi.output(args).apply((a: any) => getDatabasesReplica(a, opts))
}

export interface GetDatabasesReplicaOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
    /**
     * The name of the database replica.
     */
    replicaName: pulumi.Input<string>;
}
