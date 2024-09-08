// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listVolumeSnapshots(args: ListVolumeSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.volumes.v2.ListVolumeSnapshotsItems> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:volumes/v2:listVolumeSnapshots", {
        "volumeId": args.volumeId,
    }, opts);
}

export interface ListVolumeSnapshotsArgs {
    /**
     * The ID of the block storage volume.
     */
    volumeId: string;
}
export function listVolumeSnapshotsOutput(args: ListVolumeSnapshotsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.volumes.v2.ListVolumeSnapshotsItems> {
    return pulumi.output(args).apply((a: any) => listVolumeSnapshots(a, opts))
}

export interface ListVolumeSnapshotsOutputArgs {
    /**
     * The ID of the block storage volume.
     */
    volumeId: pulumi.Input<string>;
}
