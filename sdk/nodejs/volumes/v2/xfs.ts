// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Xfs extends pulumi.CustomResource {
    /**
     * Get an existing Xfs resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Xfs {
        return new Xfs(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:volumes/v2:Xfs';

    /**
     * Returns true if the given object is an instance of Xfs.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Xfs {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Xfs.__pulumiType;
    }

    /**
     * A time value given in ISO8601 combined date and time format that represents when the block storage volume was created.
     */
    public readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * An optional free-form text field to describe a block storage volume.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet.
     */
    public readonly dropletIds!: pulumi.Output<number[] | undefined>;
    public readonly filesystemLabel!: pulumi.Output<string | undefined>;
    /**
     * The name of the filesystem type to be used on the volume. When provided, the volume will automatically be formatted to the specified filesystem type. Currently, the available options are `ext4` and `xfs`. Pre-formatted volumes are automatically mounted when attached to Ubuntu, Debian, Fedora, Fedora Atomic, and CentOS Droplets created on or after April 26, 2018. Attaching pre-formatted volumes to other Droplets is not recommended.
     */
    public readonly filesystemType!: pulumi.Output<string | undefined>;
    /**
     * A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The slug identifier for the region where the resource will initially be  available.
     */
    public readonly region!: pulumi.Output<enums.volumes.v2.XfsPropertiesRegion | undefined>;
    /**
     * The size of the block storage volume in GiB (1024^3). This field does not apply  when creating a volume from a snapshot.
     */
    public readonly sizeGigabytes!: pulumi.Output<number | undefined>;
    /**
     * The unique identifier for the volume snapshot from which to create the volume.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly volume!: pulumi.Output<outputs.volumes.v2.VolumeFull | undefined>;

    /**
     * Create a Xfs resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: XfsArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["createdAt"] = args ? args.createdAt : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dropletIds"] = args ? args.dropletIds : undefined;
            resourceInputs["filesystemLabel"] = args ? args.filesystemLabel : undefined;
            resourceInputs["filesystemType"] = args ? args.filesystemType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sizeGigabytes"] = args ? args.sizeGigabytes : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volume"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dropletIds"] = undefined /*out*/;
            resourceInputs["filesystemLabel"] = undefined /*out*/;
            resourceInputs["filesystemType"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["sizeGigabytes"] = undefined /*out*/;
            resourceInputs["snapshotId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["volume"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Xfs.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Xfs resource.
 */
export interface XfsArgs {
    /**
     * A time value given in ISO8601 combined date and time format that represents when the block storage volume was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An optional free-form text field to describe a block storage volume.
     */
    description?: pulumi.Input<string>;
    /**
     * An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet.
     */
    dropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    filesystemLabel?: pulumi.Input<string>;
    /**
     * The name of the filesystem type to be used on the volume. When provided, the volume will automatically be formatted to the specified filesystem type. Currently, the available options are `ext4` and `xfs`. Pre-formatted volumes are automatically mounted when attached to Ubuntu, Debian, Fedora, Fedora Atomic, and CentOS Droplets created on or after April 26, 2018. Attaching pre-formatted volumes to other Droplets is not recommended.
     */
    filesystemType?: pulumi.Input<string>;
    /**
     * A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.
     */
    name?: pulumi.Input<string>;
    /**
     * The slug identifier for the region where the resource will initially be  available.
     */
    region: pulumi.Input<enums.volumes.v2.XfsPropertiesRegion>;
    /**
     * The size of the block storage volume in GiB (1024^3). This field does not apply  when creating a volume from a snapshot.
     */
    sizeGigabytes?: pulumi.Input<number>;
    /**
     * The unique identifier for the volume snapshot from which to create the volume.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
