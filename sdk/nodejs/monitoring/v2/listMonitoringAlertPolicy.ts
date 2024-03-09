// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listMonitoringAlertPolicy(args?: ListMonitoringAlertPolicyArgs, opts?: pulumi.InvokeOptions): Promise<ListMonitoringAlertPolicyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:listMonitoringAlertPolicy", {
    }, opts);
}

export interface ListMonitoringAlertPolicyArgs {
}

export interface ListMonitoringAlertPolicyResult {
    readonly items: outputs.monitoring.v2.ListMonitoringAlertPolicy;
}
export function listMonitoringAlertPolicyOutput(opts?: pulumi.InvokeOptions): pulumi.Output<ListMonitoringAlertPolicyResult> {
    return pulumi.output(listMonitoringAlertPolicy(opts))
}