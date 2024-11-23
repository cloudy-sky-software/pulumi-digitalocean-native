// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<outputs.tags.v2.GetTagProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:tags/v2:getTag", {
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagArgs {
    /**
     * The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag.
     */
    tagId: string;
}
export function getTagOutput(args: GetTagOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.tags.v2.GetTagProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:tags/v2:getTag", {
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagOutputArgs {
    /**
     * The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag.
     */
    tagId: pulumi.Input<string>;
}
