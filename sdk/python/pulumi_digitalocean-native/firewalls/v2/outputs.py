# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'Firewall',
    'FirewallPropertiesPendingChangesItemProperties',
    'FirewallPropertiesTags',
    'FirewallRulesInboundRulesItem',
    'FirewallRulesOutboundRulesItem',
    'GetFirewallsProperties',
    'ListFirewalls',
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'Tags',
]

@pulumi.output_type
class Firewall(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "dropletIds":
            suggest = "droplet_ids"
        elif key == "inboundRules":
            suggest = "inbound_rules"
        elif key == "outboundRules":
            suggest = "outbound_rules"
        elif key == "pendingChanges":
            suggest = "pending_changes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Firewall. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Firewall.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Firewall.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 droplet_ids: Optional[Sequence[int]] = None,
                 id: Optional[str] = None,
                 inbound_rules: Optional[Sequence['outputs.FirewallRulesInboundRulesItem']] = None,
                 name: Optional[str] = None,
                 outbound_rules: Optional[Sequence['outputs.FirewallRulesOutboundRulesItem']] = None,
                 pending_changes: Optional[Sequence['outputs.FirewallPropertiesPendingChangesItemProperties']] = None,
                 status: Optional['FirewallPropertiesStatus'] = None,
                 tags: Optional['outputs.FirewallPropertiesTags'] = None):
        """
        :param str created_at: A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        :param Sequence[int] droplet_ids: An array containing the IDs of the Droplets assigned to the firewall.
        :param str id: A unique ID that can be used to identify and reference a firewall.
        :param str name: A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        :param Sequence['FirewallPropertiesPendingChangesItemProperties'] pending_changes: An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        :param 'FirewallPropertiesStatus' status: A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if droplet_ids is not None:
            pulumi.set(__self__, "droplet_ids", droplet_ids)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)
        if pending_changes is not None:
            pulumi.set(__self__, "pending_changes", pending_changes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the firewall was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> Optional[Sequence[int]]:
        """
        An array containing the IDs of the Droplets assigned to the firewall.
        """
        return pulumi.get(self, "droplet_ids")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        A unique ID that can be used to identify and reference a firewall.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[Sequence['outputs.FirewallRulesInboundRulesItem']]:
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A human-readable name for a firewall. The name must begin with an alphanumeric character. Subsequent characters must either be alphanumeric characters, a period (.), or a dash (-).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[Sequence['outputs.FirewallRulesOutboundRulesItem']]:
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="pendingChanges")
    def pending_changes(self) -> Optional[Sequence['outputs.FirewallPropertiesPendingChangesItemProperties']]:
        """
        An array of objects each containing the fields "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.
        """
        return pulumi.get(self, "pending_changes")

    @property
    @pulumi.getter
    def status(self) -> Optional['FirewallPropertiesStatus']:
        """
        A status string indicating the current state of the firewall. This can be "waiting", "succeeded", or "failed".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional['outputs.FirewallPropertiesTags']:
        return pulumi.get(self, "tags")


@pulumi.output_type
class FirewallPropertiesPendingChangesItemProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dropletId":
            suggest = "droplet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FirewallPropertiesPendingChangesItemProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FirewallPropertiesPendingChangesItemProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FirewallPropertiesPendingChangesItemProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 droplet_id: Optional[int] = None,
                 removing: Optional[bool] = None,
                 status: Optional[str] = None):
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)
        if removing is not None:
            pulumi.set(__self__, "removing", removing)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[int]:
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter
    def removing(self) -> Optional[bool]:
        return pulumi.get(self, "removing")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


@pulumi.output_type
class FirewallPropertiesTags(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class FirewallRulesInboundRulesItem(dict):
    def __init__(__self__, *,
                 ports: str,
                 protocol: 'FirewallRuleBaseProtocol'):
        """
        :param str ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param 'FirewallRuleBaseProtocol' protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> str:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def protocol(self) -> 'FirewallRuleBaseProtocol':
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")


@pulumi.output_type
class FirewallRulesOutboundRulesItem(dict):
    def __init__(__self__, *,
                 ports: str,
                 protocol: 'FirewallRuleBaseProtocol'):
        """
        :param str ports: The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        :param 'FirewallRuleBaseProtocol' protocol: The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def ports(self) -> str:
        """
        The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "0" when all ports are open for a protocol. For ICMP rules this parameter will always return "0".
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def protocol(self) -> 'FirewallRuleBaseProtocol':
        """
        The type of traffic to be allowed. This may be one of `tcp`, `udp`, or `icmp`.
        """
        return pulumi.get(self, "protocol")


@pulumi.output_type
class GetFirewallsProperties(dict):
    def __init__(__self__, *,
                 firewall: Optional['outputs.Firewall'] = None):
        if firewall is not None:
            pulumi.set(__self__, "firewall", firewall)

    @property
    @pulumi.getter
    def firewall(self) -> Optional['outputs.Firewall']:
        return pulumi.get(self, "firewall")


@pulumi.output_type
class ListFirewalls(dict):
    def __init__(__self__, *,
                 meta: 'outputs.MetaMeta',
                 firewalls: Optional[Sequence['outputs.Firewall']] = None,
                 links: Optional['outputs.PageLinks'] = None):
        pulumi.set(__self__, "meta", meta)
        if firewalls is not None:
            pulumi.set(__self__, "firewalls", firewalls)
        if links is not None:
            pulumi.set(__self__, "links", links)

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def firewalls(self) -> Optional[Sequence['outputs.Firewall']]:
        return pulumi.get(self, "firewalls")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")


@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[int] = None):
        """
        :param int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class PageLinks(dict):
    def __init__(__self__, *,
                 pages: Optional['outputs.PageLinksPagesProperties'] = None):
        if pages is not None:
            pulumi.set(__self__, "pages", pages)

    @property
    @pulumi.getter
    def pages(self) -> Optional['outputs.PageLinksPagesProperties']:
        return pulumi.get(self, "pages")


@pulumi.output_type
class PageLinksPagesProperties(dict):
    def __init__(__self__, *,
                 first: Optional[str] = None,
                 last: Optional[str] = None,
                 next: Optional[str] = None,
                 prev: Optional[str] = None):
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if prev is not None:
            pulumi.set(__self__, "prev", prev)

    @property
    @pulumi.getter
    def first(self) -> Optional[str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class Tags(dict):
    def __init__(__self__):
        pass


