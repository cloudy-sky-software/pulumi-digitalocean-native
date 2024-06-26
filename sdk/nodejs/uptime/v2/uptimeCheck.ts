// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class UptimeCheck extends pulumi.CustomResource {
    /**
     * Get an existing UptimeCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UptimeCheck {
        return new UptimeCheck(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:uptime/v2:UptimeCheck';

    /**
     * Returns true if the given object is an instance of UptimeCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UptimeCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UptimeCheck.__pulumiType;
    }

    public /*out*/ readonly check!: pulumi.Output<outputs.uptime.v2.Check | undefined>;
    /**
     * A boolean value indicating whether the check is enabled/disabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * A human-friendly display name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array containing the selected regions to perform healthchecks from.
     */
    public readonly regions!: pulumi.Output<enums.uptime.v2.CheckUpdatableRegionsItem[]>;
    /**
     * The endpoint to perform healthchecks on.
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * The type of health check to perform.
     */
    public readonly type!: pulumi.Output<enums.uptime.v2.CheckUpdatableType>;

    /**
     * Create a UptimeCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UptimeCheckArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["enabled"] = (args ? args.enabled : undefined) ?? true;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["check"] = undefined /*out*/;
        } else {
            resourceInputs["check"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["regions"] = undefined /*out*/;
            resourceInputs["target"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UptimeCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UptimeCheck resource.
 */
export interface UptimeCheckArgs {
    /**
     * A boolean value indicating whether the check is enabled/disabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A human-friendly display name.
     */
    name?: pulumi.Input<string>;
    /**
     * An array containing the selected regions to perform healthchecks from.
     */
    regions?: pulumi.Input<pulumi.Input<enums.uptime.v2.CheckUpdatableRegionsItem>[]>;
    /**
     * The endpoint to perform healthchecks on.
     */
    target?: pulumi.Input<string>;
    /**
     * The type of health check to perform.
     */
    type?: pulumi.Input<enums.uptime.v2.CheckUpdatableType>;
}
