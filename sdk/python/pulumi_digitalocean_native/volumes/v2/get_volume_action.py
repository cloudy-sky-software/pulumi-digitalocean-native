# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetVolumeActionProperties',
    'AwaitableGetVolumeActionProperties',
    'get_volume_action',
    'get_volume_action_output',
]

@pulumi.output_type
class GetVolumeActionProperties:
    def __init__(__self__, action=None):
        if action and not isinstance(action, dict):
            raise TypeError("Expected argument 'action' to be a dict")
        pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter
    def action(self) -> Optional['outputs.VolumeAction']:
        return pulumi.get(self, "action")


class AwaitableGetVolumeActionProperties(GetVolumeActionProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeActionProperties(
            action=self.action)


def get_volume_action(action_id: Optional[builtins.str] = None,
                      volume_id: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeActionProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str action_id: A unique numeric ID that can be used to identify and reference an action.
    :param builtins.str volume_id: The ID of the block storage volume.
    """
    __args__ = dict()
    __args__['actionId'] = action_id
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:volumes/v2:getVolumeAction', __args__, opts=opts, typ=GetVolumeActionProperties).value

    return AwaitableGetVolumeActionProperties(
        action=pulumi.get(__ret__, 'action'))
def get_volume_action_output(action_id: Optional[pulumi.Input[builtins.str]] = None,
                             volume_id: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVolumeActionProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str action_id: A unique numeric ID that can be used to identify and reference an action.
    :param builtins.str volume_id: The ID of the block storage volume.
    """
    __args__ = dict()
    __args__['actionId'] = action_id
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:volumes/v2:getVolumeAction', __args__, opts=opts, typ=GetVolumeActionProperties)
    return __ret__.apply(lambda __response__: GetVolumeActionProperties(
        action=pulumi.get(__response__, 'action')))
