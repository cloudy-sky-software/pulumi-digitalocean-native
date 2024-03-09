// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listCdnEndpoints(args?: ListCdnEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<ListCdnEndpointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:cdn/v2:listCdnEndpoints", {
    }, opts);
}

export interface ListCdnEndpointsArgs {
}

export interface ListCdnEndpointsResult {
    readonly items: outputs.cdn.v2.ListCdnEndpoints;
}
export function listCdnEndpointsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListCdnEndpointsResult> {
    return pulumi.output(listCdnEndpoints(opts))
}