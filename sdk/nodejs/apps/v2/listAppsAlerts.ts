// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listAppsAlerts(args: ListAppsAlertsArgs, opts?: pulumi.InvokeOptions): Promise<ListAppsAlertsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:listAppsAlerts", {
        "appId": args.appId,
    }, opts);
}

export interface ListAppsAlertsArgs {
    /**
     * The app ID
     */
    appId: string;
}

export interface ListAppsAlertsResult {
    readonly items: outputs.apps.v2.AppsListAlertsResponse;
}
export function listAppsAlertsOutput(args: ListAppsAlertsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListAppsAlertsResult> {
    return pulumi.output(args).apply((a: any) => listAppsAlerts(a, opts))
}

export interface ListAppsAlertsOutputArgs {
    /**
     * The app ID
     */
    appId: pulumi.Input<string>;
}
