// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AssignArgs } from "./assign";
export type Assign = import("./assign").Assign;
export const Assign: typeof import("./assign").Assign = null as any;
utilities.lazyLoad(exports, ["Assign"], () => require("./assign"));

export { GetReservedIPArgs, GetReservedIPOutputArgs } from "./getReservedIP";
export const getReservedIP: typeof import("./getReservedIP").getReservedIP = null as any;
export const getReservedIPOutput: typeof import("./getReservedIP").getReservedIPOutput = null as any;
utilities.lazyLoad(exports, ["getReservedIP","getReservedIPOutput"], () => require("./getReservedIP"));

export { GetReservedIPsActionArgs, GetReservedIPsActionOutputArgs } from "./getReservedIPsAction";
export const getReservedIPsAction: typeof import("./getReservedIPsAction").getReservedIPsAction = null as any;
export const getReservedIPsActionOutput: typeof import("./getReservedIPsAction").getReservedIPsActionOutput = null as any;
utilities.lazyLoad(exports, ["getReservedIPsAction","getReservedIPsActionOutput"], () => require("./getReservedIPsAction"));

export { ListReservedIPsArgs } from "./listReservedIPs";
export const listReservedIPs: typeof import("./listReservedIPs").listReservedIPs = null as any;
export const listReservedIPsOutput: typeof import("./listReservedIPs").listReservedIPsOutput = null as any;
utilities.lazyLoad(exports, ["listReservedIPs","listReservedIPsOutput"], () => require("./listReservedIPs"));

export { ListReservedIPsActionsArgs, ListReservedIPsActionsOutputArgs } from "./listReservedIPsActions";
export const listReservedIPsActions: typeof import("./listReservedIPsActions").listReservedIPsActions = null as any;
export const listReservedIPsActionsOutput: typeof import("./listReservedIPsActions").listReservedIPsActionsOutput = null as any;
utilities.lazyLoad(exports, ["listReservedIPsActions","listReservedIPsActionsOutput"], () => require("./listReservedIPsActions"));

export { ReservedIPArgs } from "./reservedIP";
export type ReservedIP = import("./reservedIP").ReservedIP;
export const ReservedIP: typeof import("./reservedIP").ReservedIP = null as any;
utilities.lazyLoad(exports, ["ReservedIP"], () => require("./reservedIP"));

export { UnassignArgs } from "./unassign";
export type Unassign = import("./unassign").Unassign;
export const Unassign: typeof import("./unassign").Unassign = null as any;
utilities.lazyLoad(exports, ["Unassign"], () => require("./unassign"));


// Export enums:
export * from "../../types/enums/reserved_ips/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:reserved_ips/v2:Assign":
                return new Assign(name, <any>undefined, { urn })
            case "digitalocean-native:reserved_ips/v2:ReservedIP":
                return new ReservedIP(name, <any>undefined, { urn })
            case "digitalocean-native:reserved_ips/v2:Unassign":
                return new Unassign(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "reserved_ips/v2", _module)
