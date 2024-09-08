// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetVpcArgs, GetVpcOutputArgs } from "./getVpc";
export const getVpc: typeof import("./getVpc").getVpc = null as any;
export const getVpcOutput: typeof import("./getVpc").getVpcOutput = null as any;
utilities.lazyLoad(exports, ["getVpc","getVpcOutput"], () => require("./getVpc"));

export { ListVpcsArgs } from "./listVpcs";
export const listVpcs: typeof import("./listVpcs").listVpcs = null as any;
export const listVpcsOutput: typeof import("./listVpcs").listVpcsOutput = null as any;
utilities.lazyLoad(exports, ["listVpcs","listVpcsOutput"], () => require("./listVpcs"));

export { ListVpcsMembersArgs, ListVpcsMembersOutputArgs } from "./listVpcsMembers";
export const listVpcsMembers: typeof import("./listVpcsMembers").listVpcsMembers = null as any;
export const listVpcsMembersOutput: typeof import("./listVpcsMembers").listVpcsMembersOutput = null as any;
utilities.lazyLoad(exports, ["listVpcsMembers","listVpcsMembersOutput"], () => require("./listVpcsMembers"));

export { VpcArgs } from "./vpc";
export type Vpc = import("./vpc").Vpc;
export const Vpc: typeof import("./vpc").Vpc = null as any;
utilities.lazyLoad(exports, ["Vpc"], () => require("./vpc"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:vpcs/v2:Vpc":
                return new Vpc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "vpcs/v2", _module)
