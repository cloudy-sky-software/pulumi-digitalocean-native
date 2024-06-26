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
    'ListAppsInstanceSizesResult',
    'AwaitableListAppsInstanceSizesResult',
    'list_apps_instance_sizes',
    'list_apps_instance_sizes_output',
]

@pulumi.output_type
class ListAppsInstanceSizesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.AppsListInstanceSizesResponse':
        return pulumi.get(self, "items")


class AwaitableListAppsInstanceSizesResult(ListAppsInstanceSizesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAppsInstanceSizesResult(
            items=self.items)


def list_apps_instance_sizes(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListAppsInstanceSizesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:listAppsInstanceSizes', __args__, opts=opts, typ=ListAppsInstanceSizesResult).value

    return AwaitableListAppsInstanceSizesResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_apps_instance_sizes)
def list_apps_instance_sizes_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListAppsInstanceSizesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
