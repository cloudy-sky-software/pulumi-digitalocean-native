// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class FunctionsTrigger extends pulumi.CustomResource {
    /**
     * Get an existing FunctionsTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FunctionsTrigger {
        return new FunctionsTrigger(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:functions/v2:FunctionsTrigger';

    /**
     * Returns true if the given object is an instance of FunctionsTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionsTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionsTrigger.__pulumiType;
    }

    /**
     * Name of function(action) that exists in the given namespace.
     */
    public readonly function!: pulumi.Output<string>;
    /**
     * Indicates weather the trigger is paused or unpaused.
     */
    public readonly isEnabled!: pulumi.Output<boolean>;
    /**
     * The trigger's unique name within the namespace.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Trigger details for SCHEDULED type, where body is optional.
     */
    public readonly scheduledDetails!: pulumi.Output<outputs.functions.v2.ScheduledDetails>;
    public /*out*/ readonly trigger!: pulumi.Output<outputs.functions.v2.TriggerInfo | undefined>;
    /**
     * One of different type of triggers. Currently only SCHEDULED is supported.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a FunctionsTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionsTriggerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.function === undefined) && !opts.urn) {
                throw new Error("Missing required property 'function'");
            }
            if ((!args || args.isEnabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isEnabled'");
            }
            if ((!args || args.scheduledDetails === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduledDetails'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["function"] = args ? args.function : undefined;
            resourceInputs["isEnabled"] = args ? args.isEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["scheduledDetails"] = args ? args.scheduledDetails : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["trigger"] = undefined /*out*/;
        } else {
            resourceInputs["function"] = undefined /*out*/;
            resourceInputs["isEnabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["scheduledDetails"] = undefined /*out*/;
            resourceInputs["trigger"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionsTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FunctionsTrigger resource.
 */
export interface FunctionsTriggerArgs {
    /**
     * Name of function(action) that exists in the given namespace.
     */
    function: pulumi.Input<string>;
    /**
     * Indicates weather the trigger is paused or unpaused.
     */
    isEnabled: pulumi.Input<boolean>;
    /**
     * The trigger's unique name within the namespace.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the namespace to be managed.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * Trigger details for SCHEDULED type, where body is optional.
     */
    scheduledDetails: pulumi.Input<inputs.functions.v2.ScheduledDetailsArgs>;
    /**
     * One of different type of triggers. Currently only SCHEDULED is supported.
     */
    type: pulumi.Input<string>;
}
