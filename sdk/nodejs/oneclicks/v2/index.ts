// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ListOneClicksArgs, ListOneClicksResult } from "./listOneClicks";
export const listOneClicks: typeof import("./listOneClicks").listOneClicks = null as any;
export const listOneClicksOutput: typeof import("./listOneClicks").listOneClicksOutput = null as any;
utilities.lazyLoad(exports, ["listOneClicks","listOneClicksOutput"], () => require("./listOneClicks"));

export { OneClicksInstallKubernetesArgs } from "./oneClicksInstallKubernetes";
export type OneClicksInstallKubernetes = import("./oneClicksInstallKubernetes").OneClicksInstallKubernetes;
export const OneClicksInstallKubernetes: typeof import("./oneClicksInstallKubernetes").OneClicksInstallKubernetes = null as any;
utilities.lazyLoad(exports, ["OneClicksInstallKubernetes"], () => require("./oneClicksInstallKubernetes"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:1-clicks/v2:OneClicksInstallKubernetes":
                return new OneClicksInstallKubernetes(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "1-clicks/v2", _module)