// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getMonitoringDropletLoad1Metrics(args?: GetMonitoringDropletLoad1MetricsArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoringDropletLoad1MetricsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:monitoring/v2:getMonitoringDropletLoad1Metrics", {
    }, opts);
}

export interface GetMonitoringDropletLoad1MetricsArgs {
}

export interface GetMonitoringDropletLoad1MetricsResult {
    readonly items: outputs.monitoring.v2.Metrics;
}
export function getMonitoringDropletLoad1MetricsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitoringDropletLoad1MetricsResult> {
    return pulumi.output(getMonitoringDropletLoad1Metrics(opts))
}
