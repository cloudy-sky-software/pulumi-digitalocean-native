// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

export class KubernetesRegistry extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KubernetesRegistry {
        return new KubernetesRegistry(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:kubernetes/v2:KubernetesRegistry';

    /**
     * Returns true if the given object is an instance of KubernetesRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesRegistry.__pulumiType;
    }

    /**
     * An array containing the UUIDs of Kubernetes clusters.
     */
    public readonly clusterUuids!: pulumi.Output<string[] | undefined>;

    /**
     * Create a KubernetesRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KubernetesRegistryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["clusterUuids"] = args ? args.clusterUuids : undefined;
        } else {
            resourceInputs["clusterUuids"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubernetesRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KubernetesRegistry resource.
 */
export interface KubernetesRegistryArgs {
    /**
     * An array containing the UUIDs of Kubernetes clusters.
     */
    clusterUuids?: pulumi.Input<pulumi.Input<string>[]>;
}
