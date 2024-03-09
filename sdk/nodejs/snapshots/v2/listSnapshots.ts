// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listSnapshots(args?: ListSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<ListSnapshotsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:snapshots/v2:listSnapshots", {
    }, opts);
}

export interface ListSnapshotsArgs {
}

export interface ListSnapshotsResult {
    readonly items: outputs.snapshots.v2.ListSnapshots;
}
export function listSnapshotsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListSnapshotsResult> {
    return pulumi.output(listSnapshots(opts))
}