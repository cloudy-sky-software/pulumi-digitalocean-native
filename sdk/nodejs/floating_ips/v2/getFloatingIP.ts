// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getFloatingIP(args: GetFloatingIPArgs, opts?: pulumi.InvokeOptions): Promise<GetFloatingIPResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:floating_ips/v2:getFloatingIP", {
        "floatingIp": args.floatingIp,
    }, opts);
}

export interface GetFloatingIPArgs {
    /**
     * A floating IP address.
     */
    floatingIp: string;
}

export interface GetFloatingIPResult {
    readonly items: outputs.floating_ips.v2.GetFloatingIPProperties;
}
export function getFloatingIPOutput(args: GetFloatingIPOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFloatingIPResult> {
    return pulumi.output(args).apply((a: any) => getFloatingIP(a, opts))
}

export interface GetFloatingIPOutputArgs {
    /**
     * A floating IP address.
     */
    floatingIp: pulumi.Input<string>;
}