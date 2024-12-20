// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listImages(args?: ListImagesArgs, opts?: pulumi.InvokeOptions): Promise<outputs.images.v2.ListImagesItems> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:images/v2:listImages", {
    }, opts);
}

export interface ListImagesArgs {
}
export function listImagesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.images.v2.ListImagesItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:images/v2:listImages", {
    }, opts);
}

