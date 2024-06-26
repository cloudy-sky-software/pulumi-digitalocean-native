// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getMonitoringDropletBandwidthMetric(args?: GetMonitoringDropletBandwidthMetricArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoringDropletBandwidthMetricResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:getMonitoringDropletBandwidthMetric", {
    }, opts);
}

export interface GetMonitoringDropletBandwidthMetricArgs {
}

export interface GetMonitoringDropletBandwidthMetricResult {
    readonly items: outputs.monitoring.v2.Metrics;
}
export function getMonitoringDropletBandwidthMetricOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitoringDropletBandwidthMetricResult> {
    return pulumi.output(getMonitoringDropletBandwidthMetric(opts))
}
