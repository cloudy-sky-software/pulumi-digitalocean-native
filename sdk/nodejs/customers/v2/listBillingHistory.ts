// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listBillingHistory(args?: ListBillingHistoryArgs, opts?: pulumi.InvokeOptions): Promise<outputs.customers.v2.ListBillingHistoryItems> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:customers/v2:listBillingHistory", {
    }, opts);
}

export interface ListBillingHistoryArgs {
}
export function listBillingHistoryOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.customers.v2.ListBillingHistoryItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:customers/v2:listBillingHistory", {
    }, opts);
}

