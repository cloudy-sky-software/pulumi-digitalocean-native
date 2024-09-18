// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabasesConfig(args: GetDatabasesConfigArgs, opts?: pulumi.InvokeOptions): Promise<outputs.databases.v2.DatabaseConfig> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabasesConfig", {
        "databaseClusterUuid": args.databaseClusterUuid,
    }, opts);
}

export interface GetDatabasesConfigArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
}
export function getDatabasesConfigOutput(args: GetDatabasesConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.databases.v2.DatabaseConfig> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:databases/v2:getDatabasesConfig", {
        "databaseClusterUuid": args.databaseClusterUuid,
    }, opts);
}

export interface GetDatabasesConfigOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
}
