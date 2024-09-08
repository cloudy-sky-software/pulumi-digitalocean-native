# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'Action',
    'AwaitableAction',
    'get_image_action',
    'get_image_action_output',
]

@pulumi.output_type
class Action:
    def __init__(__self__, completed_at=None, id=None, region=None, region_slug=None, resource_id=None, resource_type=None, started_at=None, status=None, type=None):
        if completed_at and not isinstance(completed_at, str):
            raise TypeError("Expected argument 'completed_at' to be a str")
        pulumi.set(__self__, "completed_at", completed_at)
        if id and not isinstance(id, int):
            raise TypeError("Expected argument 'id' to be a int")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, dict):
            raise TypeError("Expected argument 'region' to be a dict")
        pulumi.set(__self__, "region", region)
        if region_slug and not isinstance(region_slug, dict):
            raise TypeError("Expected argument 'region_slug' to be a dict")
        pulumi.set(__self__, "region_slug", region_slug)
        if resource_id and not isinstance(resource_id, int):
            raise TypeError("Expected argument 'resource_id' to be a int")
        pulumi.set(__self__, "resource_id", resource_id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if started_at and not isinstance(started_at, str):
            raise TypeError("Expected argument 'started_at' to be a str")
        pulumi.set(__self__, "started_at", started_at)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
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


class AwaitableAction(Action):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Action(
            completed_at=self.completed_at,
            id=self.id,
            region=self.region,
            region_slug=self.region_slug,
            resource_id=self.resource_id,
            resource_type=self.resource_type,
            started_at=self.started_at,
            status=self.status,
            type=self.type)


def get_image_action(action_id: Optional[str] = None,
                     image_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAction:
    """
    Use this data source to access information about an existing resource.

    :param str action_id: A unique numeric ID that can be used to identify and reference an action.
    :param str image_id: A unique number that can be used to identify and reference a specific image.
    """
    __args__ = dict()
    __args__['actionId'] = action_id
    __args__['imageId'] = image_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:images/v2:getImageAction', __args__, opts=opts, typ=Action).value

    return AwaitableAction(
        completed_at=pulumi.get(__ret__, 'completed_at'),
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        region_slug=pulumi.get(__ret__, 'region_slug'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        started_at=pulumi.get(__ret__, 'started_at'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_image_action)
def get_image_action_output(action_id: Optional[pulumi.Input[str]] = None,
                            image_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Action]:
    """
    Use this data source to access information about an existing resource.

    :param str action_id: A unique numeric ID that can be used to identify and reference an action.
    :param str image_id: A unique number that can be used to identify and reference a specific image.
    """
    ...
