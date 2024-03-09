// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getVolumeSnapshotsById(args: GetVolumeSnapshotsByIdArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeSnapshotsByIdResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:volumes/v2:getVolumeSnapshotsById", {
        "snapshotId": args.snapshotId,
    }, opts);
}

export interface GetVolumeSnapshotsByIdArgs {
    /**
     * Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
     */
    snapshotId: string;
}

export interface GetVolumeSnapshotsByIdResult {
    readonly items: outputs.volumes.v2.GetVolumeSnapshotsByIdProperties;
}
export function getVolumeSnapshotsByIdOutput(args: GetVolumeSnapshotsByIdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeSnapshotsByIdResult> {
    return pulumi.output(args).apply((a: any) => getVolumeSnapshotsById(a, opts))
}

export interface GetVolumeSnapshotsByIdOutputArgs {
    /**
     * Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
     */
    snapshotId: pulumi.Input<string>;
}
