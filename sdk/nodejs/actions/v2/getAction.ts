// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAction(args: GetActionArgs, opts?: pulumi.InvokeOptions): Promise<outputs.actions.v2.GetActionProperties> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:actions/v2:getAction", {
        "actionId": args.actionId,
    }, opts);
}

export interface GetActionArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: string;
}
export function getActionOutput(args: GetActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.actions.v2.GetActionProperties> {
    return pulumi.output(args).apply((a: any) => getAction(a, opts))
}

export interface GetActionOutputArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: pulumi.Input<string>;
}
