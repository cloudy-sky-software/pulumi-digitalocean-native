// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDropletsFirewalls(args: ListDropletsFirewallsArgs, opts?: pulumi.InvokeOptions): Promise<ListDropletsFirewallsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:listDropletsFirewalls", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface ListDropletsFirewallsArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}

export interface ListDropletsFirewallsResult {
    readonly items: outputs.droplets.v2.ListDropletsFirewallsItems;
}
export function listDropletsFirewallsOutput(args: ListDropletsFirewallsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListDropletsFirewallsResult> {
    return pulumi.output(args).apply((a: any) => listDropletsFirewalls(a, opts))
}

export interface ListDropletsFirewallsOutputArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}
