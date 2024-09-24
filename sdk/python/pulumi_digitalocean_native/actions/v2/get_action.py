# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetActionProperties',
    'AwaitableGetActionProperties',
    'get_action',
    'get_action_output',
]

@pulumi.output_type
class GetActionProperties:
    def __init__(__self__, action=None):
        if action and not isinstance(action, dict):
            raise TypeError("Expected argument 'action' to be a dict")
        pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter
    def action(self) -> Optional['outputs.Action']:
        return pulumi.get(self, "action")


class AwaitableGetActionProperties(GetActionProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionProperties(
            action=self.action)


def get_action(action_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionProperties:
    """
    Use this data source to access information about an existing resource.

    :param str action_id: A unique numeric ID that can be used to identify and reference an action.
    """
    __args__ = dict()
    __args__['actionId'] = action_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:actions/v2:getAction', __args__, opts=opts, typ=GetActionProperties).value

    return AwaitableGetActionProperties(
        action=pulumi.get(__ret__, 'action'))
def get_action_output(action_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActionProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str action_id: A unique numeric ID that can be used to identify and reference an action.
    """
    __args__ = dict()
    __args__['actionId'] = action_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:actions/v2:getAction', __args__, opts=opts, typ=GetActionProperties)
    return __ret__.apply(lambda __response__: GetActionProperties(
        action=pulumi.get(__response__, 'action')))
