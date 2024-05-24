// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ActionStatus = {
    InProgress: "in-progress",
    Completed: "completed",
    Errored: "errored",
} as const;

/**
 * The current status of the action. This can be "in-progress", "completed", or "errored".
 */
export type ActionStatus = (typeof ActionStatus)[keyof typeof ActionStatus];

export const DropletActionType = {
    EnableBackups: "enable_backups",
    DisableBackups: "disable_backups",
    Reboot: "reboot",
    PowerCycle: "power_cycle",
    Shutdown: "shutdown",
    PowerOff: "power_off",
    PowerOn: "power_on",
    Restore: "restore",
    PasswordReset: "password_reset",
    Resize: "resize",
    Rebuild: "rebuild",
    Rename: "rename",
    ChangeKernel: "change_kernel",
    EnableIpv6: "enable_ipv6",
    Snapshot: "snapshot",
} as const;

/**
 * The type of action to initiate for the Droplet.
 */
export type DropletActionType = (typeof DropletActionType)[keyof typeof DropletActionType];

export const DropletSnapshotPropertiesType = {
    Snapshot: "snapshot",
    Backup: "backup",
} as const;

/**
 * Describes the kind of image. It may be one of `snapshot` or `backup`. This specifies whether an image is a user-generated Droplet snapshot or automatically created Droplet backup.
 */
export type DropletSnapshotPropertiesType = (typeof DropletSnapshotPropertiesType)[keyof typeof DropletSnapshotPropertiesType];

export const DropletStatus = {
    New: "new",
    Active: "active",
    Off: "off",
    Archive: "archive",
} as const;

/**
 * A status string indicating the state of the Droplet instance. This may be "new", "active", "off", or "archive".
 */
export type DropletStatus = (typeof DropletStatus)[keyof typeof DropletStatus];

export const FirewallPropertiesStatus = {
    Waiting: "waiting",
    Succeeded: "succeeded",
    Failed: "failed",
} as const;

/**
 * A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
 */
export type FirewallPropertiesStatus = (typeof FirewallPropertiesStatus)[keyof typeof FirewallPropertiesStatus];

export const FirewallRuleBaseProtocol = {
    Tcp: "tcp",
    Udp: "udp",
    Icmp: "icmp",
} as const;

/**
 * The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
 */
export type FirewallRuleBaseProtocol = (typeof FirewallRuleBaseProtocol)[keyof typeof FirewallRuleBaseProtocol];

export const ImageDistribution = {
    ArchLinux: "Arch Linux",
    CentOS: "CentOS",
    CoreOS: "CoreOS",
    Debian: "Debian",
    Fedora: "Fedora",
    FedoraAtomic: "Fedora Atomic",
    FreeBSD: "FreeBSD",
    Gentoo: "Gentoo",
    OpenSUSE: "openSUSE",
    RancherOS: "RancherOS",
    RockyLinux: "Rocky Linux",
    Ubuntu: "Ubuntu",
    Unknown: "Unknown",
} as const;

/**
 * The name of a custom image's distribution. Currently, the valid values are  `Arch Linux`, `CentOS`, `CoreOS`, `Debian`, `Fedora`, `Fedora Atomic`,  `FreeBSD`, `Gentoo`, `openSUSE`, `RancherOS`, `Rocky Linux`, `Ubuntu`, and `Unknown`.  Any other value will be accepted but ignored, and `Unknown` will be used in its place.
 */
export type ImageDistribution = (typeof ImageDistribution)[keyof typeof ImageDistribution];

export const ImageRegionsItem = {
    Ams1: "ams1",
    Ams2: "ams2",
    Ams3: "ams3",
    Blr1: "blr1",
    Fra1: "fra1",
    Lon1: "lon1",
    Nyc1: "nyc1",
    Nyc2: "nyc2",
    Nyc3: "nyc3",
    Sfo1: "sfo1",
    Sfo2: "sfo2",
    Sfo3: "sfo3",
    Sgp1: "sgp1",
    Tor1: "tor1",
} as const;

/**
 * The slug identifier for the region where the resource will initially be  available.
 */
export type ImageRegionsItem = (typeof ImageRegionsItem)[keyof typeof ImageRegionsItem];

export const ImageStatus = {
    New: "NEW",
    Available: "available",
    Pending: "pending",
    Deleted: "deleted",
    Retired: "retired",
} as const;

/**
 * A status string indicating the state of a custom image. This may be `NEW`,
 *  `available`, `pending`, `deleted`, or `retired`.
 */
export type ImageStatus = (typeof ImageStatus)[keyof typeof ImageStatus];

export const ImageType = {
    Base: "base",
    Snapshot: "snapshot",
    Backup: "backup",
    Custom: "custom",
    Admin: "admin",
} as const;

/**
 * Describes the kind of image. It may be one of `base`, `snapshot`, `backup`, `custom`, or `admin`. Respectively, this specifies whether an image is a DigitalOcean base OS image, user-generated Droplet snapshot, automatically created Droplet backup, user-provided virtual machine image, or an image used for DigitalOcean managed resources (e.g. DOKS worker nodes).
 */
export type ImageType = (typeof ImageType)[keyof typeof ImageType];

export const NetworkV4Type = {
    Public: "public",
    Private: "private",
} as const;

/**
 * The type of the IPv4 network interface.
 */
export type NetworkV4Type = (typeof NetworkV4Type)[keyof typeof NetworkV4Type];

export const NetworkV6Type = {
    Public: "public",
} as const;

/**
 * The type of the IPv6 network interface.
 *
 * **Note**: IPv6 private  networking is not currently supported.
 */
export type NetworkV6Type = (typeof NetworkV6Type)[keyof typeof NetworkV6Type];

export const Type = {
    EnableBackups: "enable_backups",
    DisableBackups: "disable_backups",
    Reboot: "reboot",
    PowerCycle: "power_cycle",
    Shutdown: "shutdown",
    PowerOff: "power_off",
    PowerOn: "power_on",
    Restore: "restore",
    PasswordReset: "password_reset",
    Resize: "resize",
    Rebuild: "rebuild",
    Rename: "rename",
    ChangeKernel: "change_kernel",
    EnableIpv6: "enable_ipv6",
    Snapshot: "snapshot",
} as const;

/**
 * The type of action to initiate for the Droplet.
 */
export type Type = (typeof Type)[keyof typeof Type];
