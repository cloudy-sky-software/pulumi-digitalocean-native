// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getApp(args: GetAppArgs, opts?: pulumi.InvokeOptions): Promise<GetAppResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:getApp", {
        "id": args.id,
    }, opts);
}

export interface GetAppArgs {
    /**
     * The ID of the app
     */
    id: string;
}

export interface GetAppResult {
    readonly items: outputs.apps.v2.AppResponse;
}
export function getAppOutput(args: GetAppOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppResult> {
    return pulumi.output(args).apply((a: any) => getApp(a, opts))
}

export interface GetAppOutputArgs {
    /**
     * The ID of the app
     */
    id: pulumi.Input<string>;
}