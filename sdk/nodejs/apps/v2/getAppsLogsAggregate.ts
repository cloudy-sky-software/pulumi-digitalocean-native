// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAppsLogsAggregate(args: GetAppsLogsAggregateArgs, opts?: pulumi.InvokeOptions): Promise<outputs.apps.v2.AppsGetLogsResponse> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:getAppsLogsAggregate", {
        "appId": args.appId,
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetAppsLogsAggregateArgs {
    /**
     * The app ID
     */
    appId: string;
    /**
     * The deployment ID
     */
    deploymentId: string;
}
export function getAppsLogsAggregateOutput(args: GetAppsLogsAggregateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.apps.v2.AppsGetLogsResponse> {
    return pulumi.output(args).apply((a: any) => getAppsLogsAggregate(a, opts))
}

export interface GetAppsLogsAggregateOutputArgs {
    /**
     * The app ID
     */
    appId: pulumi.Input<string>;
    /**
     * The deployment ID
     */
    deploymentId: pulumi.Input<string>;
}
