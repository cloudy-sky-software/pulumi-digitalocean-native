# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from ._enums import *

__all__ = [
    'FirewallPropertiesPendingChangesItemPropertiesArgs',
    'FirewallPropertiesPendingChangesItemPropertiesArgsDict',
    'FirewallPropertiesTagsArgs',
    'FirewallPropertiesTagsArgsDict',
    'FirewallRulesInboundRulesItemArgs',
    'FirewallRulesInboundRulesItemArgsDict',
    'FirewallRulesOutboundRulesItemArgs',
    'FirewallRulesOutboundRulesItemArgsDict',
    'TagsArgs',
    'TagsArgsDict',
]

MYPY = False

if not MYPY:
    class FirewallPropertiesPendingChangesItemPropertiesArgsDict(TypedDict):
        droplet_id: NotRequired[pulumi.Input[builtins.int]]
        removing: NotRequired[pulumi.Input[builtins.bool]]
        status: NotRequired[pulumi.Input[builtins.str]]
elif False:
    FirewallPropertiesPendingChangesItemPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FirewallPropertiesPendingChangesItemPropertiesArgs:
    def __init__(__self__, *,
                 droplet_id: Optional[pulumi.Input[builtins.int]] = None,
                 removing: Optional[pulumi.Input[builtins.bool]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)
        if removing is not None:
            pulumi.set(__self__, "removing", removing)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter
    def removing(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "removing")

    @removing.setter
    def removing(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "removing", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


if not MYPY:
    class FirewallPropertiesTagsArgsDict(TypedDict):
        pass
elif False:
    FirewallPropertiesTagsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FirewallPropertiesTagsArgs:
    def __init__(__self__):
        pass


if not MYPY:
    class FirewallRulesInboundRulesItemArgsDict(TypedDict):
        ports: pulumi.Input[builtins.str]
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        protocol: pulumi.Input['FirewallRuleBaseProtocol']
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
elif False:
    FirewallRulesInboundRulesItemArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FirewallRulesInboundRulesItemArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[builtins.str],
                 protocol: pulumi.Input['FirewallRuleBaseProtocol']):
        """
        :param pulumi.Input[builtins.str] ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param pulumi.Input['FirewallRuleBaseProtocol'] protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[builtins.str]:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['FirewallRuleBaseProtocol']:
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['FirewallRuleBaseProtocol']):
        pulumi.set(self, "protocol", value)


if not MYPY:
    class FirewallRulesOutboundRulesItemArgsDict(TypedDict):
        ports: pulumi.Input[builtins.str]
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        protocol: pulumi.Input['FirewallRuleBaseProtocol']
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
elif False:
    FirewallRulesOutboundRulesItemArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FirewallRulesOutboundRulesItemArgs:
    def __init__(__self__, *,
                 ports: pulumi.Input[builtins.str],
                 protocol: pulumi.Input['FirewallRuleBaseProtocol']):
        """
        :param pulumi.Input[builtins.str] ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param pulumi.Input['FirewallRuleBaseProtocol'] protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Input[builtins.str]:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['FirewallRuleBaseProtocol']:
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['FirewallRuleBaseProtocol']):
        pulumi.set(self, "protocol", value)


if not MYPY:
    class TagsArgsDict(TypedDict):
        pass
elif False:
    TagsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TagsArgs:
    def __init__(__self__):
        pass


