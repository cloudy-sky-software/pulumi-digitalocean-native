// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDropletsSnapshots(args: ListDropletsSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.droplets.v2.ListDropletsSnapshotsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:listDropletsSnapshots", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface ListDropletsSnapshotsArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}
export function listDropletsSnapshotsOutput(args: ListDropletsSnapshotsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.droplets.v2.ListDropletsSnapshotsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:droplets/v2:listDropletsSnapshots", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface ListDropletsSnapshotsOutputArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}
