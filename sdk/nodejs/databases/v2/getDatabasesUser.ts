// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDatabasesUser(args: GetDatabasesUserArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:getDatabasesUser", {
        "databaseClusterUuid": args.databaseClusterUuid,
        "username": args.username,
    }, opts);
}

export interface GetDatabasesUserArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
    /**
     * The name of the database user.
     */
    username: string;
}

export interface GetDatabasesUserResult {
    readonly items: outputs.databases.v2.GetDatabasesUserProperties;
}
export function getDatabasesUserOutput(args: GetDatabasesUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesUserResult> {
    return pulumi.output(args).apply((a: any) => getDatabasesUser(a, opts))
}

export interface GetDatabasesUserOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
    /**
     * The name of the database user.
     */
    username: pulumi.Input<string>;
}
