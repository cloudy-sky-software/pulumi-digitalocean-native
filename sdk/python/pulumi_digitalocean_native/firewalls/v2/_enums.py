# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'FirewallPropertiesStatus',
    'FirewallRuleBaseProtocol',
]


class FirewallPropertiesStatus(builtins.str, Enum):
    """
    A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
    """
    WAITING = "waiting"
    SUCCEEDED = "succeeded"
    FAILED = "failed"


class FirewallRuleBaseProtocol(builtins.str, Enum):
    """
    The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
    """
    TCP = "tcp"
    UDP = "udp"
    ICMP = "icmp"
