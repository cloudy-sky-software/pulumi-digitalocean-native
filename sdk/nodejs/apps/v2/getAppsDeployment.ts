// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAppsDeployment(args: GetAppsDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<outputs.apps.v2.AppsDeploymentResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:apps/v2:getAppsDeployment", {
        "appId": args.appId,
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetAppsDeploymentArgs {
    /**
     * The app ID
     */
    appId: string;
    /**
     * The deployment ID
     */
    deploymentId: string;
}
export function getAppsDeploymentOutput(args: GetAppsDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.apps.v2.AppsDeploymentResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:apps/v2:getAppsDeployment", {
        "appId": args.appId,
        "deploymentId": args.deploymentId,
    }, opts);
}

export interface GetAppsDeploymentOutputArgs {
    /**
     * The app ID
     */
    appId: pulumi.Input<string>;
    /**
     * The deployment ID
     */
    deploymentId: pulumi.Input<string>;
}
