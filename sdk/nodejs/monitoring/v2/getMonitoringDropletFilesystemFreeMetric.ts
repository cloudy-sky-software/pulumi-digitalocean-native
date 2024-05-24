// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getMonitoringDropletFilesystemFreeMetric(args?: GetMonitoringDropletFilesystemFreeMetricArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoringDropletFilesystemFreeMetricResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemFreeMetric", {
    }, opts);
}

export interface GetMonitoringDropletFilesystemFreeMetricArgs {
}

export interface GetMonitoringDropletFilesystemFreeMetricResult {
    readonly items: outputs.monitoring.v2.Metrics;
}
export function getMonitoringDropletFilesystemFreeMetricOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitoringDropletFilesystemFreeMetricResult> {
    return pulumi.output(getMonitoringDropletFilesystemFreeMetric(opts))
}