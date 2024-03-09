// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getAccount(args?: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:account/v2:getAccount", {
    }, opts);
}

export interface GetAccountArgs {
}

export interface GetAccountResult {
    readonly items: outputs.account.v2.GetAccountProperties;
}
export function getAccountOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(getAccount(opts))
}