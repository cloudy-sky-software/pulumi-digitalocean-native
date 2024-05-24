// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Detach extends pulumi.CustomResource {
    /**
     * Get an existing Detach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Detach {
        return new Detach(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:volumes/v2:Detach';

    /**
     * Returns true if the given object is an instance of Detach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Detach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Detach.__pulumiType;
    }

    public /*out*/ readonly action!: pulumi.Output<outputs.volumes.v2.VolumeAction | undefined>;
    /**
     * The unique identifier for the Droplet the volume will be attached or detached from.
     */
    public readonly dropletId!: pulumi.Output<number | undefined>;
    /**
     * The slug identifier for the region where the resource will initially be  available.
     */
    public readonly region!: pulumi.Output<enums.volumes.v2.VolumeActionCreateBaseRegion | undefined>;
    /**
     * The volume action to initiate.
     */
    public readonly type!: pulumi.Output<enums.volumes.v2.VolumeActionCreateBaseType | undefined>;

    /**
     * Create a Detach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DetachArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dropletId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dropletId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dropletId"] = args ? args.dropletId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["action"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["dropletId"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Detach.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Detach resource.
 */
export interface DetachArgs {
    /**
     * The unique identifier for the Droplet the volume will be attached or detached from.
     */
    dropletId: pulumi.Input<number>;
    /**
     * The slug identifier for the region where the resource will initially be  available.
     */
    region?: pulumi.Input<enums.volumes.v2.VolumeActionCreateBaseRegion>;
    /**
     * The volume action to initiate.
     */
    type: pulumi.Input<enums.volumes.v2.VolumeActionCreateBaseType>;
    /**
     * The ID of the block storage volume.
     */
    volumeId?: pulumi.Input<string>;
}
