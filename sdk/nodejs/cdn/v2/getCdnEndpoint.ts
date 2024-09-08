// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getCdnEndpoint(args: GetCdnEndpointArgs, opts?: pulumi.InvokeOptions): Promise<outputs.cdn.v2.GetCdnEndpointProperties> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:cdn/v2:getCdnEndpoint", {
        "cdnId": args.cdnId,
    }, opts);
}

export interface GetCdnEndpointArgs {
    /**
     * A unique identifier for a CDN endpoint.
     */
    cdnId: string;
}
export function getCdnEndpointOutput(args: GetCdnEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.cdn.v2.GetCdnEndpointProperties> {
    return pulumi.output(args).apply((a: any) => getCdnEndpoint(a, opts))
}

export interface GetCdnEndpointOutputArgs {
    /**
     * A unique identifier for a CDN endpoint.
     */
    cdnId: pulumi.Input<string>;
}
