// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabases(args: GetDatabasesArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabases", {
        "databaseClusterUuid": args.databaseClusterUuid,
        "databaseName": args.databaseName,
    }, opts);
}

export interface GetDatabasesArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
    /**
     * The name of the database.
     */
    databaseName: string;
}

export interface GetDatabasesResult {
    readonly items: outputs.databases.v2.GetDatabasesProperties;
}
export function getDatabasesOutput(args: GetDatabasesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesResult> {
    return pulumi.output(args).apply((a: any) => getDatabases(a, opts))
}

export interface GetDatabasesOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    databaseName: pulumi.Input<string>;
}
