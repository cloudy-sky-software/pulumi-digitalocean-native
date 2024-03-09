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
    'ListDropletsResult',
    'AwaitableListDropletsResult',
    'list_droplets',
    'list_droplets_output',
]

@pulumi.output_type
class ListDropletsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDroplets':
        return pulumi.get(self, "items")


class AwaitableListDropletsResult(ListDropletsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDropletsResult(
            items=self.items)


def list_droplets(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDropletsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:listDroplets', __args__, opts=opts, typ=ListDropletsResult).value

    return AwaitableListDropletsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_droplets)
def list_droplets_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDropletsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...