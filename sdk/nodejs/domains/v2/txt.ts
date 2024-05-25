// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export class Txt extends pulumi.CustomResource {
    /**
     * Get an existing Txt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Txt {
        return new Txt(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean-native:domains/v2:Txt';

    /**
     * Returns true if the given object is an instance of Txt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Txt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Txt.__pulumiType;
    }

    /**
     * Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
     */
    public readonly data!: pulumi.Output<string | undefined>;
    public /*out*/ readonly domainRecord!: pulumi.Output<outputs.domains.v2.DomainRecord | undefined>;
    /**
     * An unsigned integer between 0-255 used for CAA records.
     */
    public readonly flags!: pulumi.Output<number | undefined>;
    /**
     * The host name, alias, or service being defined by the record.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The port for SRV records.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The priority for SRV and MX records.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
     */
    public readonly tag!: pulumi.Output<string | undefined>;
    /**
     * This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
     */
    public readonly ttl!: pulumi.Output<number | undefined>;
    /**
     * The type of the DNS record. For example: A, CNAME, TXT, ...
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The weight for SRV records.
     */
    public readonly weight!: pulumi.Output<number | undefined>;

    /**
     * Create a Txt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TxtArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["domainRecord"] = undefined /*out*/;
        } else {
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["domainRecord"] = undefined /*out*/;
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["tag"] = undefined /*out*/;
            resourceInputs["ttl"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["weight"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Txt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Txt resource.
 */
export interface TxtArgs {
    /**
     * Variable data depending on record type. For example, the "data" value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.
     */
    data?: pulumi.Input<string>;
    /**
     * The name of the domain itself.
     */
    domainName?: pulumi.Input<string>;
    /**
     * An unsigned integer between 0-255 used for CAA records.
     */
    flags?: pulumi.Input<number>;
    /**
     * The host name, alias, or service being defined by the record.
     */
    name?: pulumi.Input<string>;
    /**
     * The port for SRV records.
     */
    port?: pulumi.Input<number>;
    /**
     * The priority for SRV and MX records.
     */
    priority?: pulumi.Input<number>;
    /**
     * The parameter tag for CAA records. Valid values are "issue", "issuewild", or "iodef"
     */
    tag?: pulumi.Input<string>;
    /**
     * This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of the DNS record. For example: A, CNAME, TXT, ...
     */
    type: pulumi.Input<string>;
    /**
     * The weight for SRV records.
     */
    weight?: pulumi.Input<number>;
}
