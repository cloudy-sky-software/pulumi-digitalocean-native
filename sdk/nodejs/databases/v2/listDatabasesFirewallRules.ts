// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDatabasesFirewallRules(args: ListDatabasesFirewallRulesArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabasesFirewallRulesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:listDatabasesFirewallRules", {
        "databaseClusterUuid": args.databaseClusterUuid,
    }, opts);
}

export interface ListDatabasesFirewallRulesArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: string;
}

export interface ListDatabasesFirewallRulesResult {
    readonly items: outputs.databases.v2.ListDatabasesFirewallRulesProperties;
}
export function listDatabasesFirewallRulesOutput(args: ListDatabasesFirewallRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListDatabasesFirewallRulesResult> {
    return pulumi.output(args).apply((a: any) => listDatabasesFirewallRules(a, opts))
}

export interface ListDatabasesFirewallRulesOutputArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid: pulumi.Input<string>;
}
