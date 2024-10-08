// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetRegistryArgs } from "./getRegistry";
export const getRegistry: typeof import("./getRegistry").getRegistry = null as any;
export const getRegistryOutput: typeof import("./getRegistry").getRegistryOutput = null as any;
utilities.lazyLoad(exports, ["getRegistry","getRegistryOutput"], () => require("./getRegistry"));

export { GetRegistryDockerCredentialArgs } from "./getRegistryDockerCredential";
export const getRegistryDockerCredential: typeof import("./getRegistryDockerCredential").getRegistryDockerCredential = null as any;
export const getRegistryDockerCredentialOutput: typeof import("./getRegistryDockerCredential").getRegistryDockerCredentialOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryDockerCredential","getRegistryDockerCredentialOutput"], () => require("./getRegistryDockerCredential"));

export { GetRegistryOptionArgs } from "./getRegistryOption";
export const getRegistryOption: typeof import("./getRegistryOption").getRegistryOption = null as any;
export const getRegistryOptionOutput: typeof import("./getRegistryOption").getRegistryOptionOutput = null as any;
utilities.lazyLoad(exports, ["getRegistryOption","getRegistryOptionOutput"], () => require("./getRegistryOption"));

export { GetRegistrySubscriptionArgs } from "./getRegistrySubscription";
export const getRegistrySubscription: typeof import("./getRegistrySubscription").getRegistrySubscription = null as any;
export const getRegistrySubscriptionOutput: typeof import("./getRegistrySubscription").getRegistrySubscriptionOutput = null as any;
utilities.lazyLoad(exports, ["getRegistrySubscription","getRegistrySubscriptionOutput"], () => require("./getRegistrySubscription"));

export { ListRegistryGarbageCollectionsArgs, ListRegistryGarbageCollectionsOutputArgs } from "./listRegistryGarbageCollections";
export const listRegistryGarbageCollections: typeof import("./listRegistryGarbageCollections").listRegistryGarbageCollections = null as any;
export const listRegistryGarbageCollectionsOutput: typeof import("./listRegistryGarbageCollections").listRegistryGarbageCollectionsOutput = null as any;
utilities.lazyLoad(exports, ["listRegistryGarbageCollections","listRegistryGarbageCollectionsOutput"], () => require("./listRegistryGarbageCollections"));

export { ListRegistryRepositoriesArgs, ListRegistryRepositoriesOutputArgs } from "./listRegistryRepositories";
export const listRegistryRepositories: typeof import("./listRegistryRepositories").listRegistryRepositories = null as any;
export const listRegistryRepositoriesOutput: typeof import("./listRegistryRepositories").listRegistryRepositoriesOutput = null as any;
utilities.lazyLoad(exports, ["listRegistryRepositories","listRegistryRepositoriesOutput"], () => require("./listRegistryRepositories"));

export { ListRegistryRepositoriesV2Args, ListRegistryRepositoriesV2OutputArgs } from "./listRegistryRepositoriesV2";
export const listRegistryRepositoriesV2: typeof import("./listRegistryRepositoriesV2").listRegistryRepositoriesV2 = null as any;
export const listRegistryRepositoriesV2Output: typeof import("./listRegistryRepositoriesV2").listRegistryRepositoriesV2Output = null as any;
utilities.lazyLoad(exports, ["listRegistryRepositoriesV2","listRegistryRepositoriesV2Output"], () => require("./listRegistryRepositoriesV2"));

export { ListRegistryRepositoryManifestsArgs, ListRegistryRepositoryManifestsOutputArgs } from "./listRegistryRepositoryManifests";
export const listRegistryRepositoryManifests: typeof import("./listRegistryRepositoryManifests").listRegistryRepositoryManifests = null as any;
export const listRegistryRepositoryManifestsOutput: typeof import("./listRegistryRepositoryManifests").listRegistryRepositoryManifestsOutput = null as any;
utilities.lazyLoad(exports, ["listRegistryRepositoryManifests","listRegistryRepositoryManifestsOutput"], () => require("./listRegistryRepositoryManifests"));

export { ListRegistryRepositoryTagsArgs, ListRegistryRepositoryTagsOutputArgs } from "./listRegistryRepositoryTags";
export const listRegistryRepositoryTags: typeof import("./listRegistryRepositoryTags").listRegistryRepositoryTags = null as any;
export const listRegistryRepositoryTagsOutput: typeof import("./listRegistryRepositoryTags").listRegistryRepositoryTagsOutput = null as any;
utilities.lazyLoad(exports, ["listRegistryRepositoryTags","listRegistryRepositoryTagsOutput"], () => require("./listRegistryRepositoryTags"));

export { RegistryArgs } from "./registry";
export type Registry = import("./registry").Registry;
export const Registry: typeof import("./registry").Registry = null as any;
utilities.lazyLoad(exports, ["Registry"], () => require("./registry"));

export { RegistryUpdateSubscriptionArgs } from "./registryUpdateSubscription";
export type RegistryUpdateSubscription = import("./registryUpdateSubscription").RegistryUpdateSubscription;
export const RegistryUpdateSubscription: typeof import("./registryUpdateSubscription").RegistryUpdateSubscription = null as any;
utilities.lazyLoad(exports, ["RegistryUpdateSubscription"], () => require("./registryUpdateSubscription"));

export { RegistryValidateNameArgs } from "./registryValidateName";
export type RegistryValidateName = import("./registryValidateName").RegistryValidateName;
export const RegistryValidateName: typeof import("./registryValidateName").RegistryValidateName = null as any;
utilities.lazyLoad(exports, ["RegistryValidateName"], () => require("./registryValidateName"));


// Export enums:
export * from "../../types/enums/registry/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:registry/v2:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "digitalocean-native:registry/v2:RegistryUpdateSubscription":
                return new RegistryUpdateSubscription(name, <any>undefined, { urn })
            case "digitalocean-native:registry/v2:RegistryValidateName":
                return new RegistryValidateName(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "registry/v2", _module)
