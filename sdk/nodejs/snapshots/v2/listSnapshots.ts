// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listSnapshots(args?: ListSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.snapshots.v2.ListSnapshotsItems> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:snapshots/v2:listSnapshots", {
    }, opts);
}

export interface ListSnapshotsArgs {
}
export function listSnapshotsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.snapshots.v2.ListSnapshotsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:snapshots/v2:listSnapshots", {
    }, opts);
}

