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
from . import outputs
from ._enums import *

__all__ = [
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'Snapshots',
]

@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[builtins.int] = None):
        """
        :param builtins.int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[builtins.int]:
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
                 first: Optional[builtins.str] = None,
                 last: Optional[builtins.str] = None,
                 next: Optional[builtins.str] = None,
                 prev: Optional[builtins.str] = None):
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
    def first(self) -> Optional[builtins.str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[builtins.str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[builtins.str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[builtins.str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class Snapshots(dict):
    def __init__(__self__, *,
                 created_at: builtins.str,
                 min_disk_size: builtins.int,
                 name: builtins.str,
                 regions: Sequence[builtins.str],
                 resource_id: builtins.str,
                 resource_type: 'SnapshotsPropertiesResourceType',
                 size_gigabytes: builtins.float,
                 tags: Sequence[builtins.str]):
        """
        :param builtins.str created_at: A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
        :param builtins.int min_disk_size: The minimum size in GB required for a volume or Droplet to use this snapshot.
        :param builtins.str name: A human-readable name for the snapshot.
        :param Sequence[builtins.str] regions: An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
        :param builtins.str resource_id: The unique identifier for the resource that the snapshot originated from.
        :param 'SnapshotsPropertiesResourceType' resource_type: The type of resource that the snapshot originated from.
        :param builtins.float size_gigabytes: The billable size of the snapshot in gigabytes.
        :param Sequence[builtins.str] tags: An array of Tags the snapshot has been tagged with.
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "min_disk_size", min_disk_size)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "regions", regions)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "size_gigabytes", size_gigabytes)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        """
        A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="minDiskSize")
    def min_disk_size(self) -> builtins.int:
        """
        The minimum size in GB required for a volume or Droplet to use this snapshot.
        """
        return pulumi.get(self, "min_disk_size")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        A human-readable name for the snapshot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[builtins.str]:
        """
        An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> builtins.str:
        """
        The unique identifier for the resource that the snapshot originated from.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> 'SnapshotsPropertiesResourceType':
        """
        The type of resource that the snapshot originated from.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="sizeGigabytes")
    def size_gigabytes(self) -> builtins.float:
        """
        The billable size of the snapshot in gigabytes.
        """
        return pulumi.get(self, "size_gigabytes")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[builtins.str]:
        """
        An array of Tags the snapshot has been tagged with.
        """
        return pulumi.get(self, "tags")


