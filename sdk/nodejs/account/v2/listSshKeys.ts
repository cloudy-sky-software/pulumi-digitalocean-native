// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listSshKeys(args?: ListSshKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListSshKeysResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:account/v2:listSshKeys", {
    }, opts);
}

export interface ListSshKeysArgs {
}

export interface ListSshKeysResult {
    readonly items: outputs.account.v2.ListSshKeys;
}
export function listSshKeysOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListSshKeysResult> {
    return pulumi.output(listSshKeys(opts))
}
