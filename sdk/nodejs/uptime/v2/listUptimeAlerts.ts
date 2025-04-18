// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listUptimeAlerts(args: ListUptimeAlertsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.uptime.v2.ListUptimeAlertsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:uptime/v2:listUptimeAlerts", {
        "checkId": args.checkId,
    }, opts);
}

export interface ListUptimeAlertsArgs {
    /**
     * A unique identifier for a check.
     */
    checkId: string;
}
export function listUptimeAlertsOutput(args: ListUptimeAlertsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.uptime.v2.ListUptimeAlertsItems> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean-native:uptime/v2:listUptimeAlerts", {
        "checkId": args.checkId,
    }, opts);
}

export interface ListUptimeAlertsOutputArgs {
    /**
     * A unique identifier for a check.
     */
    checkId: pulumi.Input<string>;
}
