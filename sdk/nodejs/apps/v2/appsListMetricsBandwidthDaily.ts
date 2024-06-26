// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class AppsListMetricsBandwidthDaily extends pulumi.CustomResource {
    /**
     * Get an existing AppsListMetricsBandwidthDaily resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppsListMetricsBandwidthDaily {
        return new AppsListMetricsBandwidthDaily(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:apps/v2:AppsListMetricsBandwidthDaily';

    /**
     * Returns true if the given object is an instance of AppsListMetricsBandwidthDaily.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppsListMetricsBandwidthDaily {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppsListMetricsBandwidthDaily.__pulumiType;
    }

    /**
     * A list of bandwidth usage details by app.
     */
    public /*out*/ readonly appBandwidthUsage!: pulumi.Output<outputs.apps.v2.AppMetricsBandwidthUsageDetails[] | undefined>;
    /**
     * A list of app IDs to query bandwidth metrics for.
     */
    public readonly appIds!: pulumi.Output<string[]>;
    /**
     * The date for the metrics data.
     */
    public readonly date!: pulumi.Output<string | undefined>;

    /**
     * Create a AppsListMetricsBandwidthDaily resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppsListMetricsBandwidthDailyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appIds'");
            }
            resourceInputs["appIds"] = args ? args.appIds : undefined;
            resourceInputs["date"] = args ? args.date : undefined;
            resourceInputs["appBandwidthUsage"] = undefined /*out*/;
        } else {
            resourceInputs["appBandwidthUsage"] = undefined /*out*/;
            resourceInputs["appIds"] = undefined /*out*/;
            resourceInputs["date"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppsListMetricsBandwidthDaily.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppsListMetricsBandwidthDaily resource.
 */
export interface AppsListMetricsBandwidthDailyArgs {
    /**
     * A list of app IDs to query bandwidth metrics for.
     */
    appIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional day to query. Only the date component of the timestamp will be considered. Default: yesterday.
     */
    date?: pulumi.Input<string>;
}
