# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'FirewallPropertiesPendingChangesItemPropertiesArgs',
    'FirewallPropertiesTagsArgs',
    'FirewallRulesInboundRulesItemArgs',
    'FirewallRulesOutboundRulesItemArgs',
    'TagsArgs',
]

@pulumi.input_type
class FirewallPropertiesPendingChangesItemPropertiesArgs:
    def __init__(__self__, *,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 removing: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)
        if removing is not None:
            pulumi.set(__self__, "removing", removing)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter
    def removing(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "removing")

    @removing.setter
    def removing(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "removing", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class FirewallPropertiesTagsArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class FirewallRulesInboundRulesItemArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[str],
                 protocol: pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']):
        """
        :param pulumi.Input[str] ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol'] protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[str]:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[str]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']:
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class FirewallRulesOutboundRulesItemArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[str],
                 protocol: pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']):
        """
        :param pulumi.Input[str] ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol'] protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[str]:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[str]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']:
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['FirewallsRulesFirewallRuleBaseProtocol']):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class TagsArgs:
    def __init__(__self__):
        pass


