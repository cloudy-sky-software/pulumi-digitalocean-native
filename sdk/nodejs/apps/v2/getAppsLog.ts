// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAppsLog(args: GetAppsLogArgs, opts?: pulumi.InvokeOptions): Promise<outputs.apps.v2.AppsGetLogsResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:getAppsLog", {
        "appId": args.appId,
        "componentName": args.componentName,
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetAppsLogArgs {
    /**
     * The app ID
     */
    appId: string;
    /**
     * An optional component name. If set, logs will be limited to this component only.
     */
    componentName: string;
    /**
     * The deployment ID
     */
    deploymentId: string;
}
export function getAppsLogOutput(args: GetAppsLogOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.apps.v2.AppsGetLogsResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:apps/v2:getAppsLog", {
        "appId": args.appId,
        "componentName": args.componentName,
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetAppsLogOutputArgs {
    /**
     * The app ID
     */
    appId: pulumi.Input<string>;
    /**
     * An optional component name. If set, logs will be limited to this component only.
     */
    componentName: pulumi.Input<string>;
    /**
     * The deployment ID
     */
    deploymentId: pulumi.Input<string>;
}
