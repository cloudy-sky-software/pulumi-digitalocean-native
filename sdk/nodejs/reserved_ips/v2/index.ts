// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetReservedIPsArgs, GetReservedIPsResult, GetReservedIPsOutputArgs } from "./getReservedIPs";
export const getReservedIPs: typeof import("./getReservedIPs").getReservedIPs = null as any;
export const getReservedIPsOutput: typeof import("./getReservedIPs").getReservedIPsOutput = null as any;
utilities.lazyLoad(exports, ["getReservedIPs","getReservedIPsOutput"], () => require("./getReservedIPs"));

export { GetReservedIPsActionsArgs, GetReservedIPsActionsResult, GetReservedIPsActionsOutputArgs } from "./getReservedIPsActions";
export const getReservedIPsActions: typeof import("./getReservedIPsActions").getReservedIPsActions = null as any;
export const getReservedIPsActionsOutput: typeof import("./getReservedIPsActions").getReservedIPsActionsOutput = null as any;
utilities.lazyLoad(exports, ["getReservedIPsActions","getReservedIPsActionsOutput"], () => require("./getReservedIPsActions"));

export { ListReservedIPsArgs, ListReservedIPsResult } from "./listReservedIPs";
export const listReservedIPs: typeof import("./listReservedIPs").listReservedIPs = null as any;
export const listReservedIPsOutput: typeof import("./listReservedIPs").listReservedIPsOutput = null as any;
utilities.lazyLoad(exports, ["listReservedIPs","listReservedIPsOutput"], () => require("./listReservedIPs"));

export { ListReservedIPsActionsArgs, ListReservedIPsActionsResult, ListReservedIPsActionsOutputArgs } from "./listReservedIPsActions";
export const listReservedIPsActions: typeof import("./listReservedIPsActions").listReservedIPsActions = null as any;
export const listReservedIPsActionsOutput: typeof import("./listReservedIPsActions").listReservedIPsActionsOutput = null as any;
utilities.lazyLoad(exports, ["listReservedIPsActions","listReservedIPsActionsOutput"], () => require("./listReservedIPsActions"));

export { ReservedIPsArgs } from "./reservedIPs";
export type ReservedIPs = import("./reservedIPs").ReservedIPs;
export const ReservedIPs: typeof import("./reservedIPs").ReservedIPs = null as any;
utilities.lazyLoad(exports, ["ReservedIPs"], () => require("./reservedIPs"));

export { ReservedIPsActionsAssignArgs } from "./reservedIPsActionsAssign";
export type ReservedIPsActionsAssign = import("./reservedIPsActionsAssign").ReservedIPsActionsAssign;
export const ReservedIPsActionsAssign: typeof import("./reservedIPsActionsAssign").ReservedIPsActionsAssign = null as any;
utilities.lazyLoad(exports, ["ReservedIPsActionsAssign"], () => require("./reservedIPsActionsAssign"));

export { ReservedIPsActionsUnassignArgs } from "./reservedIPsActionsUnassign";
export type ReservedIPsActionsUnassign = import("./reservedIPsActionsUnassign").ReservedIPsActionsUnassign;
export const ReservedIPsActionsUnassign: typeof import("./reservedIPsActionsUnassign").ReservedIPsActionsUnassign = null as any;
utilities.lazyLoad(exports, ["ReservedIPsActionsUnassign"], () => require("./reservedIPsActionsUnassign"));


// Export enums:
export * from "../../types/enums/reserved_ips/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:reserved_ips/v2:ReservedIPs":
                return new ReservedIPs(name, <any>undefined, { urn })
            case "digitalocean-native:reserved_ips/v2:ReservedIPsActionsAssign":
                return new ReservedIPsActionsAssign(name, <any>undefined, { urn })
            case "digitalocean-native:reserved_ips/v2:ReservedIPsActionsUnassign":
                return new ReservedIPsActionsUnassign(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "reserved_ips/v2", _module)
