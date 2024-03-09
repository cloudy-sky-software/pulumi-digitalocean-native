// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listFloatingIPsAction(args: ListFloatingIPsActionArgs, opts?: pulumi.InvokeOptions): Promise<ListFloatingIPsActionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:floating_ips/v2:listFloatingIPsAction", {
        "floatingIp": args.floatingIp,
    }, opts);
}

export interface ListFloatingIPsActionArgs {
    /**
     * A floating IP address.
     */
    floatingIp: string;
}

export interface ListFloatingIPsActionResult {
    readonly items: outputs.floating_ips.v2.ListFloatingIPsAction;
}
export function listFloatingIPsActionOutput(args: ListFloatingIPsActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListFloatingIPsActionResult> {
    return pulumi.output(args).apply((a: any) => listFloatingIPsAction(a, opts))
}

export interface ListFloatingIPsActionOutputArgs {
    /**
     * A floating IP address.
     */
    floatingIp: pulumi.Input<string>;
}