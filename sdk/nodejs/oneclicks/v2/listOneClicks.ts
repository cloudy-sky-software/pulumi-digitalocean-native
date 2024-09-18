// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listOneClicks(args?: ListOneClicksArgs, opts?: pulumi.InvokeOptions): Promise<outputs.oneclicks.v2.ListOneClicksProperties> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:1-clicks/v2:listOneClicks", {
    }, opts);
}

export interface ListOneClicksArgs {
}
export function listOneClicksOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.oneclicks.v2.ListOneClicksProperties> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:1-clicks/v2:listOneClicks", {
    }, opts);
}

