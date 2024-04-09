// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function listVpcsMembers(args: ListVpcsMembersArgs, opts?: pulumi.InvokeOptions): Promise<ListVpcsMembersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:vpcs/v2:listVpcsMembers", {
        "vpcId": args.vpcId,
    }, opts);
}

export interface ListVpcsMembersArgs {
    /**
     * A unique identifier for a VPC.
     */
    vpcId: string;
}

export interface ListVpcsMembersResult {
    readonly items: outputs.vpcs.v2.ListVpcsMembersItems;
}
export function listVpcsMembersOutput(args: ListVpcsMembersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListVpcsMembersResult> {
    return pulumi.output(args).apply((a: any) => listVpcsMembers(a, opts))
}

export interface ListVpcsMembersOutputArgs {
    /**
     * A unique identifier for a VPC.
     */
    vpcId: pulumi.Input<string>;
}
