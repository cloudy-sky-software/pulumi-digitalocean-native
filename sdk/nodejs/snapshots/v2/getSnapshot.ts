// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<outputs.snapshots.v2.GetSnapshotProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:snapshots/v2:getSnapshot", {
        "snapshotId": args.snapshotId,
    }, opts);
}

export interface GetSnapshotArgs {
    /**
     * Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
     */
    snapshotId: string;
}
export function getSnapshotOutput(args: GetSnapshotOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.snapshots.v2.GetSnapshotProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:snapshots/v2:getSnapshot", {
        "snapshotId": args.snapshotId,
    }, opts);
}

export interface GetSnapshotOutputArgs {
    /**
     * Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
     */
    snapshotId: pulumi.Input<string>;
}
