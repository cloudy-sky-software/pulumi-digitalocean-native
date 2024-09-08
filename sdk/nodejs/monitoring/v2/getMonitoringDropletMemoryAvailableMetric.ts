// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getMonitoringDropletMemoryAvailableMetric(args?: GetMonitoringDropletMemoryAvailableMetricArgs, opts?: pulumi.InvokeOptions): Promise<outputs.monitoring.v2.Metrics> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:getMonitoringDropletMemoryAvailableMetric", {
    }, opts);
}

export interface GetMonitoringDropletMemoryAvailableMetricArgs {
}
export function getMonitoringDropletMemoryAvailableMetricOutput(opts?: pulumi.InvokeOptions): pulumi.Output<outputs.monitoring.v2.Metrics> {
    return pulumi.output(getMonitoringDropletMemoryAvailableMetric(opts))
}
