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
export class DropletActionsByTagEnableBackups extends pulumi.CustomResource {
    /**
     * Get an existing DropletActionsByTagEnableBackups resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DropletActionsByTagEnableBackups {
        return new DropletActionsByTagEnableBackups(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:droplets/v2:DropletActionsByTagEnableBackups';

    /**
     * Returns true if the given object is an instance of DropletActionsByTagEnableBackups.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DropletActionsByTagEnableBackups {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DropletActionsByTagEnableBackups.__pulumiType;
    }

    public /*out*/ readonly actions!: pulumi.Output<outputs.droplets.v2.Action[] | undefined>;
    /**
     * The type of action to initiate for the Droplet.
     */
    public readonly type!: pulumi.Output<enums.droplets.v2.Type>;

    /**
     * Create a DropletActionsByTagEnableBackups resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DropletActionsByTagEnableBackupsArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["actions"] = undefined /*out*/;
        } else {
            resourceInputs["actions"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DropletActionsByTagEnableBackups.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DropletActionsByTagEnableBackups resource.
 */
export interface DropletActionsByTagEnableBackupsArgs {
    /**
     * The type of action to initiate for the Droplet.
     */
    type: pulumi.Input<enums.droplets.v2.Type>;
}
