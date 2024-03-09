// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getVolumeActions(args: GetVolumeActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeActionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:volumes/v2:getVolumeActions", {
        "actionId": args.actionId,
        "volumeId": args.volumeId,
    }, opts);
}

export interface GetVolumeActionsArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: string;
    /**
     * The ID of the block storage volume.
     */
    volumeId: string;
}

export interface GetVolumeActionsResult {
    readonly items: outputs.volumes.v2.GetVolumeActionsProperties;
}
export function getVolumeActionsOutput(args: GetVolumeActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeActionsResult> {
    return pulumi.output(args).apply((a: any) => getVolumeActions(a, opts))
}

export interface GetVolumeActionsOutputArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: pulumi.Input<string>;
    /**
     * The ID of the block storage volume.
     */
    volumeId: pulumi.Input<string>;
}
