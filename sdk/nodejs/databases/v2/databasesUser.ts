// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class DatabasesUser extends pulumi.CustomResource {
    /**
     * Get an existing DatabasesUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabasesUser {
        return new DatabasesUser(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:databases/v2:DatabasesUser';

    /**
     * Returns true if the given object is an instance of DatabasesUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabasesUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabasesUser.__pulumiType;
    }

    public readonly mysqlSettings!: pulumi.Output<outputs.databases.v2.MysqlSettings | undefined>;
    /**
     * The name of a database user.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A randomly generated password for the database user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * For MongoDB clusters, set to `true` to create a read-only user.
     * This option is not currently supported for other database engines.
     */
    public readonly readonly!: pulumi.Output<boolean | undefined>;
    /**
     * A string representing the database user's role. The value will be either
     * "primary" or "normal".
     */
    public readonly role!: pulumi.Output<enums.databases.v2.DatabaseUserRole | undefined>;
    public /*out*/ readonly user!: pulumi.Output<outputs.databases.v2.DatabaseUser>;

    /**
     * Create a DatabasesUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabasesUserArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["databaseClusterUuid"] = args ? args.databaseClusterUuid : undefined;
            resourceInputs["mysqlSettings"] = args ? args.mysqlSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["readonly"] = args ? args.readonly : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["user"] = undefined /*out*/;
        } else {
            resourceInputs["mysqlSettings"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["password"] = undefined /*out*/;
            resourceInputs["readonly"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["user"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabasesUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabasesUser resource.
 */
export interface DatabasesUserArgs {
    /**
     * A unique identifier for a database cluster.
     */
    databaseClusterUuid?: pulumi.Input<string>;
    mysqlSettings?: pulumi.Input<inputs.databases.v2.MysqlSettingsArgs>;
    /**
     * The name of a database user.
     */
    name: pulumi.Input<string>;
    /**
     * A randomly generated password for the database user.
     */
    password?: pulumi.Input<string>;
    /**
     * For MongoDB clusters, set to `true` to create a read-only user.
     * This option is not currently supported for other database engines.
     */
    readonly?: pulumi.Input<boolean>;
    /**
     * A string representing the database user's role. The value will be either
     * "primary" or "normal".
     */
    role?: pulumi.Input<enums.databases.v2.DatabaseUserRole>;
}
