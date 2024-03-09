// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDomains(args: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:domains/v2:getDomains", {
        "domainName": args.domainName,
    }, opts);
}

export interface GetDomainsArgs {
    /**
     * The name of the domain itself.
     */
    domainName: string;
}

export interface GetDomainsResult {
    readonly items: outputs.domains.v2.GetDomainsProperties;
}
export function getDomainsOutput(args: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply((a: any) => getDomains(a, opts))
}

export interface GetDomainsOutputArgs {
    /**
     * The name of the domain itself.
     */
    domainName: pulumi.Input<string>;
}
