// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export function getDroplet(args: GetDropletArgs, opts?: pulumi.InvokeOptions): Promise<outputs.droplets.v2.GetDropletProperties> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean-native:droplets/v2:getDroplet", {
        "dropletId": args.dropletId,
    }, opts);
}

export interface GetDropletArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: string;
}
export function getDropletOutput(args: GetDropletOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.droplets.v2.GetDropletProperties> {
    return pulumi.output(args).apply((a: any) => getDroplet(a, opts))
}

export interface GetDropletOutputArgs {
    /**
     * A unique identifier for a Droplet instance.
     */
    dropletId: pulumi.Input<string>;
}
