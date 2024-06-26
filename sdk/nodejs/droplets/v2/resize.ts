// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Resize extends pulumi.CustomResource {
    /**
     * Get an existing Resize resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Resize {
        return new Resize(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:droplets/v2:Resize';

    /**
     * Returns true if the given object is an instance of Resize.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resize {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resize.__pulumiType;
    }

    public /*out*/ readonly action!: pulumi.Output<outputs.droplets.v2.Action | undefined>;
    /**
     * When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
     */
    public readonly disk!: pulumi.Output<boolean | undefined>;
    /**
     * The slug identifier for the size to which you wish to resize the Droplet.
     */
    public readonly size!: pulumi.Output<string | undefined>;
    /**
     * The type of action to initiate for the Droplet.
     */
    public readonly type!: pulumi.Output<enums.droplets.v2.DropletActionType | undefined>;

    /**
     * Create a Resize resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResizeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["disk"] = args ? args.disk : undefined;
            resourceInputs["dropletId"] = args ? args.dropletId : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["action"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["disk"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resize.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Resize resource.
 */
export interface ResizeArgs {
    /**
     * When `true`, the Droplet's disk will be resized in addition to its RAM and CPU. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.
     */
    disk?: pulumi.Input<boolean>;
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId?: pulumi.Input<string>;
    /**
     * The slug identifier for the size to which you wish to resize the Droplet.
     */
    size?: pulumi.Input<string>;
    /**
     * The type of action to initiate for the Droplet.
     */
    type: pulumi.Input<enums.droplets.v2.DropletActionType>;
}
