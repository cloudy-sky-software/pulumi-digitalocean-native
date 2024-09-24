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
class Snapshots(dict):
    def __init__(__self__, *,
                 created_at: str,
                 min_disk_size: int,
                 name: str,
                 regions: Sequence[str],
                 resource_id: str,
                 resource_type: 'SnapshotsPropertiesResourceType',
                 size_gigabytes: float,
                 tags: Sequence[str]):
        """
        :param str created_at: A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
        :param int min_disk_size: The minimum size in GB required for a volume or Droplet to use this snapshot.
        :param str name: A human-readable name for the snapshot.
        :param Sequence[str] regions: An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
        :param str resource_id: The unique identifier for the resource that the snapshot originated from.
        :param 'SnapshotsPropertiesResourceType' resource_type: The type of resource that the snapshot originated from.
        :param float size_gigabytes: The billable size of the snapshot in gigabytes.
        :param Sequence[str] tags: An array of Tags the snapshot has been tagged with.
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
    def created_at(self) -> str:
        """
        A time value given in ISO8601 combined date and time format that represents when the snapshot was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="minDiskSize")
    def min_disk_size(self) -> int:
        """
        The minimum size in GB required for a volume or Droplet to use this snapshot.
        """
        return pulumi.get(self, "min_disk_size")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-readable name for the snapshot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        An array of the regions that the snapshot is available in. The regions are represented by their identifying slug values.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
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
    def size_gigabytes(self) -> float:
        """
        The billable size of the snapshot in gigabytes.
        """
        return pulumi.get(self, "size_gigabytes")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        An array of Tags the snapshot has been tagged with.
        """
        return pulumi.get(self, "tags")


