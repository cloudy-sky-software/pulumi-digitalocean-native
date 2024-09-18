// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDropletActions(args: ListDropletActionsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.droplets.v2.ListDropletActionsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:listDropletActions", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface ListDropletActionsArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}
export function listDropletActionsOutput(args: ListDropletActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.droplets.v2.ListDropletActionsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:droplets/v2:listDropletActions", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface ListDropletActionsOutputArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}
