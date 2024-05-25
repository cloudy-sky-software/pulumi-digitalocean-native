// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AttachArgs } from "./attach";
export type Attach = import("./attach").Attach;
export const Attach: typeof import("./attach").Attach = null as any;
utilities.lazyLoad(exports, ["Attach"], () => require("./attach"));

export { DetachArgs } from "./detach";
export type Detach = import("./detach").Detach;
export const Detach: typeof import("./detach").Detach = null as any;
utilities.lazyLoad(exports, ["Detach"], () => require("./detach"));

export { GetVolumeArgs, GetVolumeResult, GetVolumeOutputArgs } from "./getVolume";
export const getVolume: typeof import("./getVolume").getVolume = null as any;
export const getVolumeOutput: typeof import("./getVolume").getVolumeOutput = null as any;
utilities.lazyLoad(exports, ["getVolume","getVolumeOutput"], () => require("./getVolume"));

export { GetVolumeActionArgs, GetVolumeActionResult, GetVolumeActionOutputArgs } from "./getVolumeAction";
export const getVolumeAction: typeof import("./getVolumeAction").getVolumeAction = null as any;
export const getVolumeActionOutput: typeof import("./getVolumeAction").getVolumeActionOutput = null as any;
utilities.lazyLoad(exports, ["getVolumeAction","getVolumeActionOutput"], () => require("./getVolumeAction"));

export { GetVolumeSnapshotsByIdArgs, GetVolumeSnapshotsByIdResult, GetVolumeSnapshotsByIdOutputArgs } from "./getVolumeSnapshotsById";
export const getVolumeSnapshotsById: typeof import("./getVolumeSnapshotsById").getVolumeSnapshotsById = null as any;
export const getVolumeSnapshotsByIdOutput: typeof import("./getVolumeSnapshotsById").getVolumeSnapshotsByIdOutput = null as any;
utilities.lazyLoad(exports, ["getVolumeSnapshotsById","getVolumeSnapshotsByIdOutput"], () => require("./getVolumeSnapshotsById"));

export { ListVolumeActionsArgs, ListVolumeActionsResult, ListVolumeActionsOutputArgs } from "./listVolumeActions";
export const listVolumeActions: typeof import("./listVolumeActions").listVolumeActions = null as any;
export const listVolumeActionsOutput: typeof import("./listVolumeActions").listVolumeActionsOutput = null as any;
utilities.lazyLoad(exports, ["listVolumeActions","listVolumeActionsOutput"], () => require("./listVolumeActions"));

export { ListVolumeSnapshotsArgs, ListVolumeSnapshotsResult, ListVolumeSnapshotsOutputArgs } from "./listVolumeSnapshots";
export const listVolumeSnapshots: typeof import("./listVolumeSnapshots").listVolumeSnapshots = null as any;
export const listVolumeSnapshotsOutput: typeof import("./listVolumeSnapshots").listVolumeSnapshotsOutput = null as any;
utilities.lazyLoad(exports, ["listVolumeSnapshots","listVolumeSnapshotsOutput"], () => require("./listVolumeSnapshots"));

export { ListVolumesArgs, ListVolumesResult } from "./listVolumes";
export const listVolumes: typeof import("./listVolumes").listVolumes = null as any;
export const listVolumesOutput: typeof import("./listVolumes").listVolumesOutput = null as any;
utilities.lazyLoad(exports, ["listVolumes","listVolumesOutput"], () => require("./listVolumes"));

export { ResizeArgs } from "./resize";
export type Resize = import("./resize").Resize;
export const Resize: typeof import("./resize").Resize = null as any;
utilities.lazyLoad(exports, ["Resize"], () => require("./resize"));

export { VolumeSnapshotsArgs } from "./volumeSnapshots";
export type VolumeSnapshots = import("./volumeSnapshots").VolumeSnapshots;
export const VolumeSnapshots: typeof import("./volumeSnapshots").VolumeSnapshots = null as any;
utilities.lazyLoad(exports, ["VolumeSnapshots"], () => require("./volumeSnapshots"));

export { VolumesArgs } from "./volumes";
export type Volumes = import("./volumes").Volumes;
export const Volumes: typeof import("./volumes").Volumes = null as any;
utilities.lazyLoad(exports, ["Volumes"], () => require("./volumes"));


// Export enums:
export * from "../../types/enums/volumes/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "digitalocean-native:volumes/v2:Attach":
                return new Attach(name, <any>undefined, { urn })
            case "digitalocean-native:volumes/v2:Detach":
                return new Detach(name, <any>undefined, { urn })
            case "digitalocean-native:volumes/v2:Resize":
                return new Resize(name, <any>undefined, { urn })
            case "digitalocean-native:volumes/v2:VolumeSnapshots":
                return new VolumeSnapshots(name, <any>undefined, { urn })
            case "digitalocean-native:volumes/v2:Volumes":
                return new Volumes(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("digitalocean-native", "volumes/v2", _module)
