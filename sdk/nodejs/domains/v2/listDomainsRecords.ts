// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listDomainsRecords(args: ListDomainsRecordsArgs, opts?: pulumi.InvokeOptions): Promise<ListDomainsRecordsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:domains/v2:listDomainsRecords", {
        "domainName": args.domainName,
    }, opts);
}

export interface ListDomainsRecordsArgs {
    /**
     * The name of the domain itself.
     */
    domainName: string;
}

export interface ListDomainsRecordsResult {
    readonly items: outputs.domains.v2.ListDomainsRecords;
}
export function listDomainsRecordsOutput(args: ListDomainsRecordsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListDomainsRecordsResult> {
    return pulumi.output(args).apply((a: any) => listDomainsRecords(a, opts))
}

export interface ListDomainsRecordsOutputArgs {
    /**
     * The name of the domain itself.
     */
    domainName: pulumi.Input<string>;
}
