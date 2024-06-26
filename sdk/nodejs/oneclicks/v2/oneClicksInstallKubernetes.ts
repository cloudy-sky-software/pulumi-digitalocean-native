// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

export class OneClicksInstallKubernetes extends pulumi.CustomResource {
    /**
     * Get an existing OneClicksInstallKubernetes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OneClicksInstallKubernetes {
        return new OneClicksInstallKubernetes(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:1-clicks/v2:OneClicksInstallKubernetes';

    /**
     * Returns true if the given object is an instance of OneClicksInstallKubernetes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OneClicksInstallKubernetes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OneClicksInstallKubernetes.__pulumiType;
    }

    /**
     * An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
     */
    public readonly addonSlugs!: pulumi.Output<string[]>;
    /**
     * A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
     */
    public readonly clusterUuid!: pulumi.Output<string>;
    /**
     * A message about the result of the request.
     */
    public /*out*/ readonly message!: pulumi.Output<string | undefined>;

    /**
     * Create a OneClicksInstallKubernetes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OneClicksInstallKubernetesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.addonSlugs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addonSlugs'");
            }
            if ((!args || args.clusterUuid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterUuid'");
            }
            resourceInputs["addonSlugs"] = args ? args.addonSlugs : undefined;
            resourceInputs["clusterUuid"] = args ? args.clusterUuid : undefined;
            resourceInputs["message"] = undefined /*out*/;
        } else {
            resourceInputs["addonSlugs"] = undefined /*out*/;
            resourceInputs["clusterUuid"] = undefined /*out*/;
            resourceInputs["message"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OneClicksInstallKubernetes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OneClicksInstallKubernetes resource.
 */
export interface OneClicksInstallKubernetesArgs {
    /**
     * An array of 1-Click Application slugs to be installed to the Kubernetes cluster.
     */
    addonSlugs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique ID for the Kubernetes cluster to which the 1-Click Applications will be installed.
     */
    clusterUuid: pulumi.Input<string>;
}
