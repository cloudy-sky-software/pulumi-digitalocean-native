// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetTagsArgs, GetTagsResult, GetTagsOutputArgs } from "./getTags";
export const getTags: typeof import("./getTags").getTags = null as any;
export const getTagsOutput: typeof import("./getTags").getTagsOutput = null as any;
utilities.lazyLoad(exports, ["getTags","getTagsOutput"], () => require("./getTags"));

export { ListTagsArgs, ListTagsResult } from "./listTags";
export const listTags: typeof import("./listTags").listTags = null as any;
export const listTagsOutput: typeof import("./listTags").listTagsOutput = null as any;
utilities.lazyLoad(exports, ["listTags","listTagsOutput"], () => require("./listTags"));

export { TagsArgs } from "./tags";
export type Tags = import("./tags").Tags;
export const Tags: typeof import("./tags").Tags = null as any;
utilities.lazyLoad(exports, ["Tags"], () => require("./tags"));

export { TagsAssignResourcesArgs } from "./tagsAssignResources";
export type TagsAssignResources = import("./tagsAssignResources").TagsAssignResources;
export const TagsAssignResources: typeof import("./tagsAssignResources").TagsAssignResources = null as any;
utilities.lazyLoad(exports, ["TagsAssignResources"], () => require("./tagsAssignResources"));


// Export enums:
export * from "../../types/enums/tags/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:tags/v2:Tags":
                return new Tags(name, <any>undefined, { urn })
            case "digitalocean-native:tags/v2:TagsAssignResources":
                return new TagsAssignResources(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "tags/v2", _module)