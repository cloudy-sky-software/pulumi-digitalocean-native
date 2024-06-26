// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDomains(args?: ListDomainsArgs, opts?: pulumi.InvokeOptions): Promise<ListDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:domains/v2:listDomains", {
    }, opts);
}

export interface ListDomainsArgs {
}

export interface ListDomainsResult {
    readonly items: outputs.domains.v2.ListDomainsItems;
}
export function listDomainsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListDomainsResult> {
    return pulumi.output(listDomains(opts))
}
