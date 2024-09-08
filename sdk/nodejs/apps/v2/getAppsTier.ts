// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAppsTier(args: GetAppsTierArgs, opts?: pulumi.InvokeOptions): Promise<outputs.apps.v2.AppsGetTierResponse> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:getAppsTier", {
        "slug": args.slug,
    }, opts);
}

export interface GetAppsTierArgs {
    /**
     * The slug of the tier
     */
    slug: string;
}
export function getAppsTierOutput(args: GetAppsTierOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.apps.v2.AppsGetTierResponse> {
    return pulumi.output(args).apply((a: any) => getAppsTier(a, opts))
}

export interface GetAppsTierOutputArgs {
    /**
     * The slug of the tier
     */
    slug: pulumi.Input<string>;
}
