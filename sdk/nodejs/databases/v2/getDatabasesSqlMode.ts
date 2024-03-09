// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabasesSqlMode(args: GetDatabasesSqlModeArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesSqlModeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabasesSqlMode", {
        "databaseClusterUuid": args.databaseClusterUuid,
    }, opts);
}

export interface GetDatabasesSqlModeArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
}

export interface GetDatabasesSqlModeResult {
    readonly items: outputs.databases.v2.SqlMode;
}
export function getDatabasesSqlModeOutput(args: GetDatabasesSqlModeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesSqlModeResult> {
    return pulumi.output(args).apply((a: any) => getDatabasesSqlMode(a, opts))
}

export interface GetDatabasesSqlModeOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
}