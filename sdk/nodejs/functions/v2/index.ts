// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { FunctionsNamespaceArgs } from "./functionsNamespace";
export type FunctionsNamespace = import("./functionsNamespace").FunctionsNamespace;
export const FunctionsNamespace: typeof import("./functionsNamespace").FunctionsNamespace = null as any;
utilities.lazyLoad(exports, ["FunctionsNamespace"], () => require("./functionsNamespace"));

export { FunctionsTriggerArgs } from "./functionsTrigger";
export type FunctionsTrigger = import("./functionsTrigger").FunctionsTrigger;
export const FunctionsTrigger: typeof import("./functionsTrigger").FunctionsTrigger = null as any;
utilities.lazyLoad(exports, ["FunctionsTrigger"], () => require("./functionsTrigger"));

export { GetFunctionsNamespaceArgs, GetFunctionsNamespaceOutputArgs } from "./getFunctionsNamespace";
export const getFunctionsNamespace: typeof import("./getFunctionsNamespace").getFunctionsNamespace = null as any;
export const getFunctionsNamespaceOutput: typeof import("./getFunctionsNamespace").getFunctionsNamespaceOutput = null as any;
utilities.lazyLoad(exports, ["getFunctionsNamespace","getFunctionsNamespaceOutput"], () => require("./getFunctionsNamespace"));

export { GetFunctionsTriggerArgs, GetFunctionsTriggerOutputArgs } from "./getFunctionsTrigger";
export const getFunctionsTrigger: typeof import("./getFunctionsTrigger").getFunctionsTrigger = null as any;
export const getFunctionsTriggerOutput: typeof import("./getFunctionsTrigger").getFunctionsTriggerOutput = null as any;
utilities.lazyLoad(exports, ["getFunctionsTrigger","getFunctionsTriggerOutput"], () => require("./getFunctionsTrigger"));

export { ListFunctionsNamespacesArgs } from "./listFunctionsNamespaces";
export const listFunctionsNamespaces: typeof import("./listFunctionsNamespaces").listFunctionsNamespaces = null as any;
export const listFunctionsNamespacesOutput: typeof import("./listFunctionsNamespaces").listFunctionsNamespacesOutput = null as any;
utilities.lazyLoad(exports, ["listFunctionsNamespaces","listFunctionsNamespacesOutput"], () => require("./listFunctionsNamespaces"));

export { ListFunctionsTriggersArgs, ListFunctionsTriggersOutputArgs } from "./listFunctionsTriggers";
export const listFunctionsTriggers: typeof import("./listFunctionsTriggers").listFunctionsTriggers = null as any;
export const listFunctionsTriggersOutput: typeof import("./listFunctionsTriggers").listFunctionsTriggersOutput = null as any;
utilities.lazyLoad(exports, ["listFunctionsTriggers","listFunctionsTriggersOutput"], () => require("./listFunctionsTriggers"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:functions/v2:FunctionsNamespace":
                return new FunctionsNamespace(name, <any>undefined, { urn })
            case "digitalocean-native:functions/v2:FunctionsTrigger":
                return new FunctionsTrigger(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "functions/v2", _module)
