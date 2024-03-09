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
    'GetReservedIPsResult',
    'AwaitableGetReservedIPsResult',
    'get_reserved_ips',
    'get_reserved_ips_output',
]

@pulumi.output_type
class GetReservedIPsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetReservedIPsProperties':
        return pulumi.get(self, "items")


class AwaitableGetReservedIPsResult(GetReservedIPsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReservedIPsResult(
            items=self.items)


def get_reserved_ips(reserved_ip: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReservedIPsResult:
    """
    Use this data source to access information about an existing resource.

    :param str reserved_ip: A reserved IP address.
    """
    __args__ = dict()
    __args__['reservedIp'] = reserved_ip
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:reserved_ips/v2:getReservedIPs', __args__, opts=opts, typ=GetReservedIPsResult).value

    return AwaitableGetReservedIPsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_reserved_ips)
def get_reserved_ips_output(reserved_ip: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReservedIPsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str reserved_ip: A reserved IP address.
    """
    ...