// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class LoadBalancersForwardingRules extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancersForwardingRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancersForwardingRules {
        return new LoadBalancersForwardingRules(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:load_balancers/v2:LoadBalancersForwardingRules';

    /**
     * Returns true if the given object is an instance of LoadBalancersForwardingRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancersForwardingRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancersForwardingRules.__pulumiType;
    }

    public readonly forwardingRules!: pulumi.Output<outputs.load_balancers.v2.ForwardingRule[]>;

    /**
     * Create a LoadBalancersForwardingRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancersForwardingRulesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.forwardingRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forwardingRules'");
            }
            resourceInputs["forwardingRules"] = args ? args.forwardingRules : undefined;
            resourceInputs["lbId"] = args ? args.lbId : undefined;
        } else {
            resourceInputs["forwardingRules"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancersForwardingRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoadBalancersForwardingRules resource.
 */
export interface LoadBalancersForwardingRulesArgs {
    forwardingRules: pulumi.Input<pulumi.Input<inputs.load_balancers.v2.ForwardingRuleArgs>[]>;
    /**
     * A unique identifier for a load balancer.
     */
    lbId?: pulumi.Input<string>;
}
