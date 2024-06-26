// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:firewalls/v2:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * A time value given in ISO8601 combined date and time format that represents when the firewall was created.
     */
    public readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * An array containing the IDs of the Droplets assigned to the firewall.
     */
    public readonly dropletIds!: pulumi.Output<number[] | undefined>;
    public /*out*/ readonly firewall!: pulumi.Output<outputs.firewalls.v2.Firewall | undefined>;
    public readonly inboundRules!: pulumi.Output<outputs.firewalls.v2.FirewallRulesInboundRulesItem[] | undefined>;
    /**
     * A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly outboundRules!: pulumi.Output<outputs.firewalls.v2.FirewallRulesOutboundRulesItem[] | undefined>;
    /**
     * An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
     */
    public readonly pendingChanges!: pulumi.Output<outputs.firewalls.v2.FirewallPropertiesPendingChangesItemProperties[] | undefined>;
    /**
     * A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
     */
    public readonly status!: pulumi.Output<enums.firewalls.v2.FirewallPropertiesStatus | undefined>;
    public readonly tags!: pulumi.Output<outputs.firewalls.v2.FirewallPropertiesTags | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["createdAt"] = args ? args.createdAt : undefined;
            resourceInputs["dropletIds"] = args ? args.dropletIds : undefined;
            resourceInputs["inboundRules"] = args ? args.inboundRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outboundRules"] = args ? args.outboundRules : undefined;
            resourceInputs["pendingChanges"] = args ? args.pendingChanges : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["firewall"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dropletIds"] = undefined /*out*/;
            resourceInputs["firewall"] = undefined /*out*/;
            resourceInputs["inboundRules"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["outboundRules"] = undefined /*out*/;
            resourceInputs["pendingChanges"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Firewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * A time value given in ISO8601 combined date and time format that represents when the firewall was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An array containing the IDs of the Droplets assigned to the firewall.
     */
    dropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    inboundRules?: pulumi.Input<pulumi.Input<inputs.firewalls.v2.FirewallRulesInboundRulesItemArgs>[]>;
    /**
     * A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
     */
    name?: pulumi.Input<string>;
    outboundRules?: pulumi.Input<pulumi.Input<inputs.firewalls.v2.FirewallRulesOutboundRulesItemArgs>[]>;
    /**
     * An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
     */
    pendingChanges?: pulumi.Input<pulumi.Input<inputs.firewalls.v2.FirewallPropertiesPendingChangesItemPropertiesArgs>[]>;
    /**
     * A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
     */
    status?: pulumi.Input<enums.firewalls.v2.FirewallPropertiesStatus>;
    tags?: pulumi.Input<inputs.firewalls.v2.FirewallPropertiesTagsArgs>;
}
