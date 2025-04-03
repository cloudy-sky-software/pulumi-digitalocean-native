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
    'GetReservedIPProperties',
    'AwaitableGetReservedIPProperties',
    'get_reserved_ip',
    'get_reserved_ip_output',
]

@pulumi.output_type
class GetReservedIPProperties:
    def __init__(__self__, reserved_ip=None):
        if reserved_ip and not isinstance(reserved_ip, dict):
            raise TypeError("Expected argument 'reserved_ip' to be a dict")
        pulumi.set(__self__, "reserved_ip", reserved_ip)

    @property
    @pulumi.getter(name="reservedIp")
    def reserved_ip(self) -> Optional['outputs.ReservedIp']:
        return pulumi.get(self, "reserved_ip")


class AwaitableGetReservedIPProperties(GetReservedIPProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReservedIPProperties(
            reserved_ip=self.reserved_ip)


def get_reserved_ip(reserved_ip: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReservedIPProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str reserved_ip: A reserved IP address.
    """
    __args__ = dict()
    __args__['reservedIp'] = reserved_ip
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:reserved_ips/v2:getReservedIP', __args__, opts=opts, typ=GetReservedIPProperties).value

    return AwaitableGetReservedIPProperties(
        reserved_ip=pulumi.get(__ret__, 'reserved_ip'))
def get_reserved_ip_output(reserved_ip: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetReservedIPProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str reserved_ip: A reserved IP address.
    """
    __args__ = dict()
    __args__['reservedIp'] = reserved_ip
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:reserved_ips/v2:getReservedIP', __args__, opts=opts, typ=GetReservedIPProperties)
    return __ret__.apply(lambda __response__: GetReservedIPProperties(
        reserved_ip=pulumi.get(__response__, 'reserved_ip')))
