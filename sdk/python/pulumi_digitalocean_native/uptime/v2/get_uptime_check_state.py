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
    'GetUptimeCheckStateResult',
    'AwaitableGetUptimeCheckStateResult',
    'get_uptime_check_state',
    'get_uptime_check_state_output',
]

@pulumi.output_type
class GetUptimeCheckStateResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetUptimeCheckStateProperties':
        return pulumi.get(self, "items")


class AwaitableGetUptimeCheckStateResult(GetUptimeCheckStateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUptimeCheckStateResult(
            items=self.items)


def get_uptime_check_state(check_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUptimeCheckStateResult:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    __args__ = dict()
    __args__['checkId'] = check_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:uptime/v2:getUptimeCheckState', __args__, opts=opts, typ=GetUptimeCheckStateResult).value

    return AwaitableGetUptimeCheckStateResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_uptime_check_state)
def get_uptime_check_state_output(check_id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUptimeCheckStateResult]:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    ...
