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
    'Resources',
    'ResourcesItemProperties',
    'Tags',
    'TagsResources',
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
class Resources(dict):
    """
    An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastTaggedUri":
            suggest = "last_tagged_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Resources. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Resources.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Resources.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 count: Optional[builtins.int] = None,
                 last_tagged_uri: Optional[builtins.str] = None):
        """
        An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
        :param builtins.int count: The number of tagged objects for this type of resource.
        :param builtins.str last_tagged_uri: The URI for the last tagged object for this type of resource.
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if last_tagged_uri is not None:
            pulumi.set(__self__, "last_tagged_uri", last_tagged_uri)

    @property
    @pulumi.getter
    def count(self) -> Optional[builtins.int]:
        """
        The number of tagged objects for this type of resource.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="lastTaggedUri")
    def last_tagged_uri(self) -> Optional[builtins.str]:
        """
        The URI for the last tagged object for this type of resource.
        """
        return pulumi.get(self, "last_tagged_uri")


@pulumi.output_type
class ResourcesItemProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceId":
            suggest = "resource_id"
        elif key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourcesItemProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourcesItemProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourcesItemProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 resource_id: Optional[builtins.str] = None,
                 resource_type: Optional['ResourcesItemPropertiesResourceType'] = None):
        """
        :param builtins.str resource_id: The identifier of a resource.
        :param 'ResourcesItemPropertiesResourceType' resource_type: The type of the resource.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[builtins.str]:
        """
        The identifier of a resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional['ResourcesItemPropertiesResourceType']:
        """
        The type of the resource.
        """
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class Tags(dict):
    """
    A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
    Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
    """
    def __init__(__self__, *,
                 name: Optional[builtins.str] = None,
                 resources: Optional['outputs.TagsResources'] = None):
        """
        A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
        Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
        :param builtins.str name: The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores.
               There is a limit of 255 characters per tag.
               
               **Note:** Tag names are case stable, which means the capitalization you use when you first create a tag is canonical.
               
               When working with tags in the API, you must use the tag's canonical capitalization. For example, if you create a tag named "PROD", the URL to add that tag to a resource would be `https://api.digitalocean.com/v2/tags/PROD/resources` (not `/v2/tags/prod/resources`).
               
               Tagged resources in the control panel will always display the canonical capitalization. For example, if you create a tag named "PROD", you can tag resources in the control panel by entering "prod". The tag will still display with its canonical capitalization, "PROD".
        :param 'TagsResources' resources: An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores.
        There is a limit of 255 characters per tag.

        **Note:** Tag names are case stable, which means the capitalization you use when you first create a tag is canonical.

        When working with tags in the API, you must use the tag's canonical capitalization. For example, if you create a tag named "PROD", the URL to add that tag to a resource would be `https://api.digitalocean.com/v2/tags/PROD/resources` (not `/v2/tags/prod/resources`).

        Tagged resources in the control panel will always display the canonical capitalization. For example, if you create a tag named "PROD", you can tag resources in the control panel by entering "prod". The tag will still display with its canonical capitalization, "PROD".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def resources(self) -> Optional['outputs.TagsResources']:
        """
        An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
        """
        return pulumi.get(self, "resources")


@pulumi.output_type
class TagsResources(dict):
    """
    An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastTaggedUri":
            suggest = "last_tagged_uri"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TagsResources. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TagsResources.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TagsResources.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 count: Optional[builtins.int] = None,
                 last_tagged_uri: Optional[builtins.str] = None):
        """
        An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a `last_tagged_uri` attribute set to the last resource tagged with the current tag.
        :param builtins.int count: The number of tagged objects for this type of resource.
        :param builtins.str last_tagged_uri: The URI for the last tagged object for this type of resource.
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if last_tagged_uri is not None:
            pulumi.set(__self__, "last_tagged_uri", last_tagged_uri)

    @property
    @pulumi.getter
    def count(self) -> Optional[builtins.int]:
        """
        The number of tagged objects for this type of resource.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="lastTaggedUri")
    def last_tagged_uri(self) -> Optional[builtins.str]:
        """
        The URI for the last tagged object for this type of resource.
        """
        return pulumi.get(self, "last_tagged_uri")


