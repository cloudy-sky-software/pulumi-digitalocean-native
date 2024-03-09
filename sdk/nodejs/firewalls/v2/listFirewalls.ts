// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listFirewalls(args?: ListFirewallsArgs, opts?: pulumi.InvokeOptions): Promise<ListFirewallsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:firewalls/v2:listFirewalls", {
    }, opts);
}

export interface ListFirewallsArgs {
}

export interface ListFirewallsResult {
    readonly items: outputs.firewalls.v2.ListFirewalls;
}
export function listFirewallsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListFirewallsResult> {
    return pulumi.output(listFirewalls(opts))
}
