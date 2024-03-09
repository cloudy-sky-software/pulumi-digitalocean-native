// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { FloatingIPsArgs } from "./floatingIPs";
export type FloatingIPs = import("./floatingIPs").FloatingIPs;
export const FloatingIPs: typeof import("./floatingIPs").FloatingIPs = null as any;
utilities.lazyLoad(exports, ["FloatingIPs"], () => require("./floatingIPs"));

export { FloatingIPsActionAssignArgs } from "./floatingIPsActionAssign";
export type FloatingIPsActionAssign = import("./floatingIPsActionAssign").FloatingIPsActionAssign;
export const FloatingIPsActionAssign: typeof import("./floatingIPsActionAssign").FloatingIPsActionAssign = null as any;
utilities.lazyLoad(exports, ["FloatingIPsActionAssign"], () => require("./floatingIPsActionAssign"));

export { FloatingIPsActionUnassignArgs } from "./floatingIPsActionUnassign";
export type FloatingIPsActionUnassign = import("./floatingIPsActionUnassign").FloatingIPsActionUnassign;
export const FloatingIPsActionUnassign: typeof import("./floatingIPsActionUnassign").FloatingIPsActionUnassign = null as any;
utilities.lazyLoad(exports, ["FloatingIPsActionUnassign"], () => require("./floatingIPsActionUnassign"));

export { GetFloatingIPsArgs, GetFloatingIPsResult, GetFloatingIPsOutputArgs } from "./getFloatingIPs";
export const getFloatingIPs: typeof import("./getFloatingIPs").getFloatingIPs = null as any;
export const getFloatingIPsOutput: typeof import("./getFloatingIPs").getFloatingIPsOutput = null as any;
utilities.lazyLoad(exports, ["getFloatingIPs","getFloatingIPsOutput"], () => require("./getFloatingIPs"));

export { GetFloatingIPsActionArgs, GetFloatingIPsActionResult, GetFloatingIPsActionOutputArgs } from "./getFloatingIPsAction";
export const getFloatingIPsAction: typeof import("./getFloatingIPsAction").getFloatingIPsAction = null as any;
export const getFloatingIPsActionOutput: typeof import("./getFloatingIPsAction").getFloatingIPsActionOutput = null as any;
utilities.lazyLoad(exports, ["getFloatingIPsAction","getFloatingIPsActionOutput"], () => require("./getFloatingIPsAction"));

export { ListFloatingIPsArgs, ListFloatingIPsResult } from "./listFloatingIPs";
export const listFloatingIPs: typeof import("./listFloatingIPs").listFloatingIPs = null as any;
export const listFloatingIPsOutput: typeof import("./listFloatingIPs").listFloatingIPsOutput = null as any;
utilities.lazyLoad(exports, ["listFloatingIPs","listFloatingIPsOutput"], () => require("./listFloatingIPs"));

export { ListFloatingIPsActionArgs, ListFloatingIPsActionResult, ListFloatingIPsActionOutputArgs } from "./listFloatingIPsAction";
export const listFloatingIPsAction: typeof import("./listFloatingIPsAction").listFloatingIPsAction = null as any;
export const listFloatingIPsActionOutput: typeof import("./listFloatingIPsAction").listFloatingIPsActionOutput = null as any;
utilities.lazyLoad(exports, ["listFloatingIPsAction","listFloatingIPsActionOutput"], () => require("./listFloatingIPsAction"));


// Export enums:
export * from "../../types/enums/floating_ips/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:floating_ips/v2:FloatingIPs":
                return new FloatingIPs(name, <any>undefined, { urn })
            case "digitalocean-native:floating_ips/v2:FloatingIPsActionAssign":
                return new FloatingIPsActionAssign(name, <any>undefined, { urn })
            case "digitalocean-native:floating_ips/v2:FloatingIPsActionUnassign":
                return new FloatingIPsActionUnassign(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "floating_ips/v2", _module)
