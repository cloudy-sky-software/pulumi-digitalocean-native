// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDatabasesOptions(args?: ListDatabasesOptionsArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabasesOptionsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:databases/v2:listDatabasesOptions", {
    }, opts);
}

export interface ListDatabasesOptionsArgs {
}

export interface ListDatabasesOptionsResult {
    readonly items: outputs.databases.v2.Options;
}
export function listDatabasesOptionsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListDatabasesOptionsResult> {
    return pulumi.output(listDatabasesOptions(opts))
}
