// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Specifies the action that will be taken on the Droplet.
 */
export class DropletActionsEnableIpv6 extends pulumi.CustomResource {
    /**
     * Get an existing DropletActionsEnableIpv6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DropletActionsEnableIpv6 {
        return new DropletActionsEnableIpv6(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:droplets/v2:DropletActionsEnableIpv6';

    /**
     * Returns true if the given object is an instance of DropletActionsEnableIpv6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DropletActionsEnableIpv6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DropletActionsEnableIpv6.__pulumiType;
    }

    public /*out*/ readonly action!: pulumi.Output<outputs.droplets.v2.Action | undefined>;
    /**
     * The type of action to initiate for the Droplet.
     */
    public readonly type!: pulumi.Output<enums.droplets.v2.Type>;

    /**
     * Create a DropletActionsEnableIpv6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DropletActionsEnableIpv6Args, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dropletId"] = args ? args.dropletId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["action"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DropletActionsEnableIpv6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DropletActionsEnableIpv6 resource.
 */
export interface DropletActionsEnableIpv6Args {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId?: pulumi.Input<string>;
    /**
     * The type of action to initiate for the Droplet.
     */
    type: pulumi.Input<enums.droplets.v2.Type>;
}
