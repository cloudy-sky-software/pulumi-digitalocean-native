// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


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
