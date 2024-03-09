// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDomainsRecord(args: GetDomainsRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsRecordResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:domains/v2:getDomainsRecord", {
        "domainName": args.domainName,
        "domainRecordId": args.domainRecordId,
    }, opts);
}

export interface GetDomainsRecordArgs {
    /**
     * The name of the domain itself.
     */
    domainName: string;
    /**
     * The unique identifier of the domain record.
     */
    domainRecordId: string;
}

export interface GetDomainsRecordResult {
    readonly items: outputs.domains.v2.GetDomainsRecordProperties;
}
export function getDomainsRecordOutput(args: GetDomainsRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsRecordResult> {
    return pulumi.output(args).apply((a: any) => getDomainsRecord(a, opts))
}

export interface GetDomainsRecordOutputArgs {
    /**
     * The name of the domain itself.
     */
    domainName: pulumi.Input<string>;
    /**
     * The unique identifier of the domain record.
     */
    domainRecordId: pulumi.Input<string>;
}
