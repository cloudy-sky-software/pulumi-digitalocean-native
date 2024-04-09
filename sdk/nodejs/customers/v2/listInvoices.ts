// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listInvoices(args?: ListInvoicesArgs, opts?: pulumi.InvokeOptions): Promise<ListInvoicesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:customers/v2:listInvoices", {
    }, opts);
}

export interface ListInvoicesArgs {
}

export interface ListInvoicesResult {
    readonly items: outputs.customers.v2.ListInvoicesItems;
}
export function listInvoicesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListInvoicesResult> {
    return pulumi.output(listInvoices(opts))
}
