// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDropletAction(args: GetDropletActionArgs, opts?: pulumi.InvokeOptions): Promise<GetDropletActionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:getDropletAction", {
        "actionId": args.actionId,
        "dropletId": args.dropletId,
    }, opts);
}

export interface GetDropletActionArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: string;
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}

export interface GetDropletActionResult {
    readonly items: outputs.droplets.v2.GetDropletActionProperties;
}
export function getDropletActionOutput(args: GetDropletActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDropletActionResult> {
    return pulumi.output(args).apply((a: any) => getDropletAction(a, opts))
}

export interface GetDropletActionOutputArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: pulumi.Input<string>;
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}