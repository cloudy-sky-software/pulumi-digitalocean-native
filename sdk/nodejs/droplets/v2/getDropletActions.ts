// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDropletActions(args: GetDropletActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDropletActionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:getDropletActions", {
        "actionId": args.actionId,
        "dropletId": args.dropletId,
    }, opts);
}

export interface GetDropletActionsArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: string;
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}

export interface GetDropletActionsResult {
    readonly items: outputs.droplets.v2.GetDropletActionsProperties;
}
export function getDropletActionsOutput(args: GetDropletActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDropletActionsResult> {
    return pulumi.output(args).apply((a: any) => getDropletActions(a, opts))
}

export interface GetDropletActionsOutputArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: pulumi.Input<string>;
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}
