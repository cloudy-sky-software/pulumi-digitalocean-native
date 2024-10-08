# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import outputs

__all__ = [
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'Vpc',
    'VpcMember',
]

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
class Vpc(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "ipRange":
            suggest = "ip_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Vpc. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Vpc.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Vpc.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 default: Optional[bool] = None,
                 description: Optional[str] = None,
                 id: Optional[str] = None,
                 ip_range: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 urn: Optional[str] = None):
        """
        :param str created_at: A time value given in ISO8601 combined date and time format.
        :param bool default: A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined.
        :param str description: A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.
        :param str id: A unique ID that can be used to identify and reference the VPC.
        :param str ip_range: The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/28` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.
        :param str name: The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.
        :param str region: The slug identifier for the region where the VPC will be created.
        :param str urn: The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip_range is not None:
            pulumi.set(__self__, "ip_range", ip_range)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def default(self) -> Optional[bool]:
        """
        A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        A unique ID that can be used to identify and reference the VPC.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> Optional[str]:
        """
        The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/28` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The slug identifier for the region where the VPC will be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        """
        The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class VpcMember(dict):
    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 name: Optional[str] = None,
                 urn: Optional[str] = None):
        """
        :param str created_at: A time value given in ISO8601 combined date and time format that represents when the resource was created.
        :param str name: The name of the resource.
        :param str urn: The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the resource was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        """
        The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        return pulumi.get(self, "urn")


