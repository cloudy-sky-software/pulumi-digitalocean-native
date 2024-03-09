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
    'GetLoadBalancersResult',
    'AwaitableGetLoadBalancersResult',
    'get_load_balancers',
    'get_load_balancers_output',
]

@pulumi.output_type
class GetLoadBalancersResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetLoadBalancersProperties':
        return pulumi.get(self, "items")


class AwaitableGetLoadBalancersResult(GetLoadBalancersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancersResult(
            items=self.items)


def get_load_balancers(lb_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancersResult:
    """
    Use this data source to access information about an existing resource.

    :param str lb_id: A unique identifier for a load balancer.
    """
    __args__ = dict()
    __args__['lbId'] = lb_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:load_balancers/v2:getLoadBalancers', __args__, opts=opts, typ=GetLoadBalancersResult).value

    return AwaitableGetLoadBalancersResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_load_balancers)
def get_load_balancers_output(lb_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancersResult]:
    """
    Use this data source to access information about an existing resource.

    :param str lb_id: A unique identifier for a load balancer.
    """
    ...
