// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getMonitoringDropletCpuMetric(args?: GetMonitoringDropletCpuMetricArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoringDropletCpuMetricResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:getMonitoringDropletCpuMetric", {
    }, opts);
}

export interface GetMonitoringDropletCpuMetricArgs {
}

export interface GetMonitoringDropletCpuMetricResult {
    readonly items: outputs.monitoring.v2.Metrics;
}
export function getMonitoringDropletCpuMetricOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitoringDropletCpuMetricResult> {
    return pulumi.output(getMonitoringDropletCpuMetric(opts))
}
