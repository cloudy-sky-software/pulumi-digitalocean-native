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
    'ListCdnEndpointsResult',
    'AwaitableListCdnEndpointsResult',
    'list_cdn_endpoints',
    'list_cdn_endpoints_output',
]

@pulumi.output_type
class ListCdnEndpointsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListCdnEndpointsItems':
        return pulumi.get(self, "items")


class AwaitableListCdnEndpointsResult(ListCdnEndpointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListCdnEndpointsResult(
            items=self.items)


def list_cdn_endpoints(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListCdnEndpointsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:cdn/v2:listCdnEndpoints', __args__, opts=opts, typ=ListCdnEndpointsResult).value

    return AwaitableListCdnEndpointsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_cdn_endpoints)
def list_cdn_endpoints_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListCdnEndpointsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
