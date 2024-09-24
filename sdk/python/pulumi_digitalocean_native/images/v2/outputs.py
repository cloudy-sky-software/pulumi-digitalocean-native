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
    'Action',
    'ActionRegionSlug',
    'Image',
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
    'Region',
    'RegionSlug',
]

@pulumi.output_type
class Action(dict):
    def __init__(__self__, *,
                 completed_at: Optional[str] = None,
                 id: Optional[int] = None,
                 region: Optional['outputs.Region'] = None,
                 region_slug: Optional['outputs.ActionRegionSlug'] = None,
                 resource_id: Optional[int] = None,
                 resource_type: Optional[str] = None,
                 started_at: Optional[str] = None,
                 status: Optional['ActionStatus'] = None,
                 type: Optional[str] = None):
        """
        :param str completed_at: A time value given in ISO8601 combined date and time format that represents when the action was completed.
        :param int id: A unique numeric ID that can be used to identify and reference an action.
        :param int resource_id: A unique identifier for the resource that the action is associated with.
        :param str resource_type: The type of resource that the action is associated with.
        :param str started_at: A time value given in ISO8601 combined date and time format that represents when the action was initiated.
        :param 'ActionStatus' status: The current status of the action. This can be "in-progress", "completed", or "errored".
        :param str type: This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
        """
        if completed_at is not None:
            pulumi.set(__self__, "completed_at", completed_at)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if region_slug is not None:
            pulumi.set(__self__, "region_slug", region_slug)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if started_at is not None:
            pulumi.set(__self__, "started_at", started_at)
        if status is None:
            status = 'in-progress'
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="completedAt")
    def completed_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the action was completed.
        """
        return pulumi.get(self, "completed_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        A unique numeric ID that can be used to identify and reference an action.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional['outputs.Region']:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionSlug")
    def region_slug(self) -> Optional['outputs.ActionRegionSlug']:
        return pulumi.get(self, "region_slug")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[int]:
        """
        A unique identifier for the resource that the action is associated with.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The type of resource that the action is associated with.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="startedAt")
    def started_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the action was initiated.
        """
        return pulumi.get(self, "started_at")

    @property
    @pulumi.getter
    def status(self) -> Optional['ActionStatus']:
        """
        The current status of the action. This can be "in-progress", "completed", or "errored".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ActionRegionSlug(dict):
    def __init__(__self__):
        pass


@pulumi.output_type
class Image(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "errorMessage":
            suggest = "error_message"
        elif key == "minDiskSize":
            suggest = "min_disk_size"
        elif key == "sizeGigabytes":
            suggest = "size_gigabytes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Image. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Image.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Image.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 description: Optional[str] = None,
                 distribution: Optional['ImageDistribution'] = None,
                 error_message: Optional[str] = None,
                 id: Optional[int] = None,
                 min_disk_size: Optional[int] = None,
                 name: Optional[str] = None,
                 public: Optional[bool] = None,
                 regions: Optional[Sequence['ImageRegionsItem']] = None,
                 size_gigabytes: Optional[float] = None,
                 slug: Optional[str] = None,
                 status: Optional['ImageStatus'] = None,
                 tags: Optional[Sequence[str]] = None,
                 type: Optional['ImageType'] = None):
        """
        :param str created_at: A time value given in ISO8601 combined date and time format that represents when the image was created.
        :param str description: An optional free-form text field to describe an image.
        :param 'ImageDistribution' distribution: The name of a custom image's distribution. Currently, the valid values are  `Arch Linux`, `CentOS`, `CoreOS`, `Debian`, `Fedora`, `Fedora Atomic`,  `FreeBSD`, `Gentoo`, `openSUSE`, `RancherOS`, `Rocky Linux`, `Ubuntu`, and `Unknown`.  Any other value will be accepted but ignored, and `Unknown` will be used in its place.
        :param str error_message: A string containing information about errors that may occur when importing
                a custom image.
        :param int id: A unique number that can be used to identify and reference a specific image.
        :param int min_disk_size: The minimum disk size in GB required for a Droplet to use this image.
        :param str name: The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.
        :param bool public: This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.
        :param Sequence['ImageRegionsItem'] regions: This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.
        :param float size_gigabytes: The size of the image in gigabytes.
        :param str slug: A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.
        :param 'ImageStatus' status: A status string indicating the state of a custom image. This may be `NEW`,
                `available`, `pending`, `deleted`, or `retired`.
        :param Sequence[str] tags: A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.
        :param 'ImageType' type: Describes the kind of image. It may be one of `base`, `snapshot`, `backup`, `custom`, or `admin`. Respectively, this specifies whether an image is a DigitalOcean base OS image, user-generated Droplet snapshot, automatically created Droplet backup, user-provided virtual machine image, or an image used for DigitalOcean managed resources (e.g. DOKS worker nodes).
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if distribution is not None:
            pulumi.set(__self__, "distribution", distribution)
        if error_message is not None:
            pulumi.set(__self__, "error_message", error_message)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if min_disk_size is not None:
            pulumi.set(__self__, "min_disk_size", min_disk_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if size_gigabytes is not None:
            pulumi.set(__self__, "size_gigabytes", size_gigabytes)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the image was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional free-form text field to describe an image.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distribution(self) -> Optional['ImageDistribution']:
        """
        The name of a custom image's distribution. Currently, the valid values are  `Arch Linux`, `CentOS`, `CoreOS`, `Debian`, `Fedora`, `Fedora Atomic`,  `FreeBSD`, `Gentoo`, `openSUSE`, `RancherOS`, `Rocky Linux`, `Ubuntu`, and `Unknown`.  Any other value will be accepted but ignored, and `Unknown` will be used in its place.
        """
        return pulumi.get(self, "distribution")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> Optional[str]:
        """
        A string containing information about errors that may occur when importing
         a custom image.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        A unique number that can be used to identify and reference a specific image.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minDiskSize")
    def min_disk_size(self) -> Optional[int]:
        """
        The minimum disk size in GB required for a Droplet to use this image.
        """
        return pulumi.get(self, "min_disk_size")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The display name that has been given to an image.  This is what is shown in the control panel and is generally a descriptive title for the image in question.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def public(self) -> Optional[bool]:
        """
        This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def regions(self) -> Optional[Sequence['ImageRegionsItem']]:
        """
        This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="sizeGigabytes")
    def size_gigabytes(self) -> Optional[float]:
        """
        The size of the image in gigabytes.
        """
        return pulumi.get(self, "size_gigabytes")

    @property
    @pulumi.getter
    def slug(self) -> Optional[str]:
        """
        A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def status(self) -> Optional['ImageStatus']:
        """
        A status string indicating the state of a custom image. This may be `NEW`,
         `available`, `pending`, `deleted`, or `retired`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional['ImageType']:
        """
        Describes the kind of image. It may be one of `base`, `snapshot`, `backup`, `custom`, or `admin`. Respectively, this specifies whether an image is a DigitalOcean base OS image, user-generated Droplet snapshot, automatically created Droplet backup, user-provided virtual machine image, or an image used for DigitalOcean managed resources (e.g. DOKS worker nodes).
        """
        return pulumi.get(self, "type")


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
class Region(dict):
    def __init__(__self__, *,
                 available: bool,
                 features: Sequence[str],
                 name: str,
                 sizes: Sequence[str],
                 slug: str):
        """
        :param bool available: This is a boolean value that represents whether new Droplets can be created in this region.
        :param Sequence[str] features: This attribute is set to an array which contains features available in this region
        :param str name: The display name of the region.  This will be a full name that is used in the control panel and other interfaces.
        :param Sequence[str] sizes: This attribute is set to an array which contains the identifying slugs for the sizes available in this region.
        :param str slug: A human-readable string that is used as a unique identifier for each region.
        """
        pulumi.set(__self__, "available", available)
        pulumi.set(__self__, "features", features)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sizes", sizes)
        pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter
    def available(self) -> bool:
        """
        This is a boolean value that represents whether new Droplets can be created in this region.
        """
        return pulumi.get(self, "available")

    @property
    @pulumi.getter
    def features(self) -> Sequence[str]:
        """
        This attribute is set to an array which contains features available in this region
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The display name of the region.  This will be a full name that is used in the control panel and other interfaces.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sizes(self) -> Sequence[str]:
        """
        This attribute is set to an array which contains the identifying slugs for the sizes available in this region.
        """
        return pulumi.get(self, "sizes")

    @property
    @pulumi.getter
    def slug(self) -> str:
        """
        A human-readable string that is used as a unique identifier for each region.
        """
        return pulumi.get(self, "slug")


@pulumi.output_type
class RegionSlug(dict):
    def __init__(__self__):
        pass


