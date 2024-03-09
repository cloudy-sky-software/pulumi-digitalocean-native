// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getVolumes(args: GetVolumesArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:volumes/v2:getVolumes", {
        "volumeId": args.volumeId,
    }, opts);
}

export interface GetVolumesArgs {
    /**
     * The ID of the block storage volume.
     */
    volumeId: string;
}

export interface GetVolumesResult {
    readonly items: outputs.volumes.v2.GetVolumesProperties;
}
export function getVolumesOutput(args: GetVolumesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumesResult> {
    return pulumi.output(args).apply((a: any) => getVolumes(a, opts))
}

export interface GetVolumesOutputArgs {
    /**
     * The ID of the block storage volume.
     */
    volumeId: pulumi.Input<string>;
}