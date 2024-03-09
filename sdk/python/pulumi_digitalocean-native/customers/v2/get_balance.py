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

__all__ = [
    'GetBalanceResult',
    'AwaitableGetBalanceResult',
    'get_balance',
    'get_balance_output',
]

@pulumi.output_type
class GetBalanceResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.Balance':
        return pulumi.get(self, "items")


class AwaitableGetBalanceResult(GetBalanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBalanceResult(
            items=self.items)


def get_balance(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBalanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:customers/v2:getBalance', __args__, opts=opts, typ=GetBalanceResult).value

    return AwaitableGetBalanceResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_balance)
def get_balance_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBalanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...