// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getImageAction(args: GetImageActionArgs, opts?: pulumi.InvokeOptions): Promise<outputs.images.v2.Action> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:images/v2:getImageAction", {
        "actionId": args.actionId,
        "imageId": args.imageId,
    }, opts);
}

export interface GetImageActionArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: string;
    /**
     * A unique number that can be used to identify and reference a specific image.
     */
    imageId: string;
}
export function getImageActionOutput(args: GetImageActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.images.v2.Action> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:images/v2:getImageAction", {
        "actionId": args.actionId,
        "imageId": args.imageId,
    }, opts);
}

export interface GetImageActionOutputArgs {
    /**
     * A unique numeric ID that can be used to identify and reference an action.
     */
    actionId: pulumi.Input<string>;
    /**
     * A unique number that can be used to identify and reference a specific image.
     */
    imageId: pulumi.Input<string>;
}
