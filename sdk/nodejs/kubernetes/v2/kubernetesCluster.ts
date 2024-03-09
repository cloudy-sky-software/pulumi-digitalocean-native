// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class KubernetesCluster extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KubernetesCluster {
        return new KubernetesCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:kubernetes/v2:KubernetesCluster';

    /**
     * Returns true if the given object is an instance of KubernetesCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesCluster.__pulumiType;
    }

    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     */
    public readonly autoUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster in CIDR notation.
     */
    public /*out*/ readonly clusterSubnet!: pulumi.Output<string | undefined>;
    /**
     * A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * The base URL of the API server on the Kubernetes master node.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
     */
    public readonly ha!: pulumi.Output<boolean | undefined>;
    /**
     * The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
     */
    public /*out*/ readonly ipv4!: pulumi.Output<string | undefined>;
    public /*out*/ readonly kubernetesCluster!: pulumi.Output<outputs.kubernetes.v2.Cluster | undefined>;
    /**
     * An object specifying the maintenance window policy for the Kubernetes cluster.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.kubernetes.v2.MaintenancePolicy | undefined>;
    /**
     * A human-readable name for a Kubernetes cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An object specifying the details of the worker nodes available to the Kubernetes cluster.
     */
    public readonly nodePools!: pulumi.Output<outputs.kubernetes.v2.KubernetesNodePool[]>;
    /**
     * The slug identifier for the region where the Kubernetes cluster is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A read-only boolean value indicating if a container registry is integrated with the cluster.
     */
    public /*out*/ readonly registryEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster in CIDR notation.
     */
    public /*out*/ readonly serviceSubnet!: pulumi.Output<string | undefined>;
    /**
     * An object containing a `state` attribute whose value is set to a string indicating the current status of the cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.kubernetes.v2.StatusProperties | undefined>;
    /**
     * A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
     */
    public readonly surgeUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * A time value given in ISO8601 combined date and time format that represents when the Kubernetes cluster was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string | undefined>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
     */
    public readonly vpcUuid!: pulumi.Output<string | undefined>;

    /**
     * Create a KubernetesCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.nodePools === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodePools'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["autoUpgrade"] = (args ? args.autoUpgrade : undefined) ?? false;
            resourceInputs["ha"] = (args ? args.ha : undefined) ?? false;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodePools"] = args ? args.nodePools : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["surgeUpgrade"] = (args ? args.surgeUpgrade : undefined) ?? false;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcUuid"] = args ? args.vpcUuid : undefined;
            resourceInputs["clusterSubnet"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["ipv4"] = undefined /*out*/;
            resourceInputs["kubernetesCluster"] = undefined /*out*/;
            resourceInputs["registryEnabled"] = undefined /*out*/;
            resourceInputs["serviceSubnet"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["autoUpgrade"] = undefined /*out*/;
            resourceInputs["clusterSubnet"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["ha"] = undefined /*out*/;
            resourceInputs["ipv4"] = undefined /*out*/;
            resourceInputs["kubernetesCluster"] = undefined /*out*/;
            resourceInputs["maintenancePolicy"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodePools"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["registryEnabled"] = undefined /*out*/;
            resourceInputs["serviceSubnet"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["surgeUpgrade"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["vpcUuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubernetesCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KubernetesCluster resource.
 */
export interface KubernetesClusterArgs {
    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     */
    autoUpgrade?: pulumi.Input<boolean>;
    /**
     * A boolean value indicating whether the control plane is run in a highly available configuration in the cluster. Highly available control planes incur less downtime. The property cannot be disabled.
     */
    ha?: pulumi.Input<boolean>;
    /**
     * An object specifying the maintenance window policy for the Kubernetes cluster.
     */
    maintenancePolicy?: pulumi.Input<inputs.kubernetes.v2.MaintenancePolicyArgs>;
    /**
     * A human-readable name for a Kubernetes cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * An object specifying the details of the worker nodes available to the Kubernetes cluster.
     */
    nodePools: pulumi.Input<pulumi.Input<inputs.kubernetes.v2.KubernetesNodePoolArgs>[]>;
    /**
     * The slug identifier for the region where the Kubernetes cluster is located.
     */
    region: pulumi.Input<string>;
    /**
     * A boolean value indicating whether surge upgrade is enabled/disabled for the cluster. Surge upgrade makes cluster upgrades fast and reliable by bringing up new nodes before destroying the outdated nodes.
     */
    surgeUpgrade?: pulumi.Input<boolean>;
    /**
     * An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `k8s` and `k8s:$K8S_CLUSTER_ID`.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster. If set to a minor version (e.g. "1.14"), the latest version within it will be used (e.g. "1.14.6-do.1"); if set to "latest", the latest published version will be used. See the `/v2/kubernetes/options` endpoint to find all currently available versions.
     */
    version: pulumi.Input<string>;
    /**
     * A string specifying the UUID of the VPC to which the Kubernetes cluster is assigned.
     */
    vpcUuid?: pulumi.Input<string>;
}
