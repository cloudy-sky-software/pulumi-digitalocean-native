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
    'Project',
    'Resource',
    'ResourceLinksProperties',
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
class Project(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "isDefault":
            suggest = "is_default"
        elif key == "ownerId":
            suggest = "owner_id"
        elif key == "ownerUuid":
            suggest = "owner_uuid"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Project. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Project.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Project.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[builtins.str] = None,
                 description: Optional[builtins.str] = None,
                 environment: Optional['ProjectBaseEnvironment'] = None,
                 id: Optional[builtins.str] = None,
                 is_default: Optional[builtins.bool] = None,
                 name: Optional[builtins.str] = None,
                 owner_id: Optional[builtins.int] = None,
                 owner_uuid: Optional[builtins.str] = None,
                 purpose: Optional[builtins.str] = None,
                 updated_at: Optional[builtins.str] = None):
        """
        :param builtins.str created_at: A time value given in ISO8601 combined date and time format that represents when the project was created.
        :param builtins.str description: The description of the project. The maximum length is 255 characters.
        :param 'ProjectBaseEnvironment' environment: The environment of the project's resources.
        :param builtins.str id: The unique universal identifier of this project.
        :param builtins.bool is_default: If true, all resources will be added to this project if no project is specified.
        :param builtins.str name: The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        :param builtins.int owner_id: The integer id of the project owner.
        :param builtins.str owner_uuid: The unique universal identifier of the project owner.
        :param builtins.str purpose: The purpose of the project. The maximum length is 255 characters. It can
               have one of the following values:
               
               - Just trying out DigitalOcean
               - Class project / Educational purposes
               - Website or blog
               - Web Application
               - Service or API
               - Mobile Application
               - Machine learning / AI / Data processing
               - IoT
               - Operational / Developer tooling
               
               If another value for purpose is specified, for example, "your custom purpose",
               your purpose will be stored as `Other: your custom purpose`.
        :param builtins.str updated_at: A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if owner_uuid is not None:
            pulumi.set(__self__, "owner_uuid", owner_uuid)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the project. The maximum length is 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> Optional['ProjectBaseEnvironment']:
        """
        The environment of the project's resources.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        The unique universal identifier of this project.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[builtins.bool]:
        """
        If true, all resources will be added to this project if no project is specified.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The human-readable name for the project. The maximum length is 175 characters and the name must be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[builtins.int]:
        """
        The integer id of the project owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="ownerUuid")
    def owner_uuid(self) -> Optional[builtins.str]:
        """
        The unique universal identifier of the project owner.
        """
        return pulumi.get(self, "owner_uuid")

    @property
    @pulumi.getter
    def purpose(self) -> Optional[builtins.str]:
        """
        The purpose of the project. The maximum length is 255 characters. It can
        have one of the following values:

        - Just trying out DigitalOcean
        - Class project / Educational purposes
        - Website or blog
        - Web Application
        - Service or API
        - Mobile Application
        - Machine learning / AI / Data processing
        - IoT
        - Operational / Developer tooling

        If another value for purpose is specified, for example, "your custom purpose",
        your purpose will be stored as `Other: your custom purpose`.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was updated.
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class Resource(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "assignedAt":
            suggest = "assigned_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Resource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Resource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Resource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 assigned_at: Optional[builtins.str] = None,
                 links: Optional['outputs.ResourceLinksProperties'] = None,
                 status: Optional['ResourceStatus'] = None,
                 urn: Optional[builtins.str] = None):
        """
        :param builtins.str assigned_at: A time value given in ISO8601 combined date and time format that represents when the project was created.
        :param 'ResourceLinksProperties' links: The links object contains the `self` object, which contains the resource relationship.
        :param 'ResourceStatus' status: The status of assigning and fetching the resources.
        :param builtins.str urn: The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        if assigned_at is not None:
            pulumi.set(__self__, "assigned_at", assigned_at)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="assignedAt")
    def assigned_at(self) -> Optional[builtins.str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the project was created.
        """
        return pulumi.get(self, "assigned_at")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.ResourceLinksProperties']:
        """
        The links object contains the `self` object, which contains the resource relationship.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def status(self) -> Optional['ResourceStatus']:
        """
        The status of assigning and fetching the resources.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def urn(self) -> Optional[builtins.str]:
        """
        The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class ResourceLinksProperties(dict):
    """
    The links object contains the `self` object, which contains the resource relationship.
    """
    def __init__(__self__, *,
                 self: Optional[builtins.str] = None):
        """
        The links object contains the `self` object, which contains the resource relationship.
        :param builtins.str self: A URI that can be used to retrieve the resource.
        """
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter
    def self(self) -> Optional[builtins.str]:
        """
        A URI that can be used to retrieve the resource.
        """
        return pulumi.get(self, "self")


