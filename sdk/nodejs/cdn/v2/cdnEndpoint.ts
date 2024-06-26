// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class CdnEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing CdnEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CdnEndpoint {
        return new CdnEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:cdn/v2:CdnEndpoint';

    /**
     * Returns true if the given object is an instance of CdnEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CdnEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CdnEndpoint.__pulumiType;
    }

    /**
     * The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * A time value given in ISO8601 combined date and time format that represents when the CDN endpoint was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * The fully qualified domain name (FQDN) of the custom subdomain used with the CDN endpoint.
     */
    public readonly customDomain!: pulumi.Output<string | undefined>;
    public /*out*/ readonly endpoint!: pulumi.Output<outputs.cdn.v2.CdnEndpoint | undefined>;
    /**
     * The fully qualified domain name (FQDN) for the origin server which provides the content for the CDN. This is currently restricted to a Space.
     */
    public readonly origin!: pulumi.Output<string>;
    /**
     * The amount of time the content is cached by the CDN's edge servers in seconds. TTL must be one of 60, 600, 3600, 86400, or 604800. Defaults to 3600 (one hour) when excluded.
     */
    public readonly ttl!: pulumi.Output<enums.cdn.v2.Ttl | undefined>;

    /**
     * Create a CdnEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CdnEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.origin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'origin'");
            }
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["customDomain"] = args ? args.customDomain : undefined;
            resourceInputs["origin"] = args ? args.origin : undefined;
            resourceInputs["ttl"] = (args ? args.ttl : undefined) ?? 3600;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
        } else {
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["customDomain"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["origin"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CdnEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CdnEndpoint resource.
 */
export interface CdnEndpointArgs {
    /**
     * The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The fully qualified domain name (FQDN) of the custom subdomain used with the CDN endpoint.
     */
    customDomain?: pulumi.Input<string>;
    /**
     * The fully qualified domain name (FQDN) for the origin server which provides the content for the CDN. This is currently restricted to a Space.
     */
    origin: pulumi.Input<string>;
    /**
     * The amount of time the content is cached by the CDN's edge servers in seconds. TTL must be one of 60, 600, 3600, 86400, or 604800. Defaults to 3600 (one hour) when excluded.
     */
    ttl?: pulumi.Input<enums.cdn.v2.Ttl>;
}
