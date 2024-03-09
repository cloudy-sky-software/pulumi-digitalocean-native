// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

export class DatabasesClusterSize extends pulumi.CustomResource {
    /**
     * Get an existing DatabasesClusterSize resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabasesClusterSize {
        return new DatabasesClusterSize(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:databases/v2:DatabasesClusterSize';

    /**
     * Returns true if the given object is an instance of DatabasesClusterSize.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabasesClusterSize {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabasesClusterSize.__pulumiType;
    }

    /**
     * The number of nodes in the database cluster. Valid values are are 1-3. In addition to the primary node, up to two standby nodes may be added for highly available configurations.
     */
    public readonly numNodes!: pulumi.Output<number>;
    /**
     * A slug identifier representing desired the size of the nodes in the database cluster.
     */
    public readonly size!: pulumi.Output<string>;

    /**
     * Create a DatabasesClusterSize resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabasesClusterSizeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.numNodes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numNodes'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["databaseClusterUuid"] = args ? args.databaseClusterUuid : undefined;
            resourceInputs["numNodes"] = args ? args.numNodes : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
        } else {
            resourceInputs["numNodes"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabasesClusterSize.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabasesClusterSize resource.
 */
export interface DatabasesClusterSizeArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid?: pulumi.Input<string>;
    /**
     * The number of nodes in the database cluster. Valid values are are 1-3. In addition to the primary node, up to two standby nodes may be added for highly available configurations.
     */
    numNodes: pulumi.Input<number>;
    /**
     * A slug identifier representing desired the size of the nodes in the database cluster.
     */
    size: pulumi.Input<string>;
}