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
    'ListAppsRegionsResult',
    'AwaitableListAppsRegionsResult',
    'list_apps_regions',
    'list_apps_regions_output',
]

@pulumi.output_type
class ListAppsRegionsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.AppsListRegionsResponse':
        return pulumi.get(self, "items")


class AwaitableListAppsRegionsResult(ListAppsRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAppsRegionsResult(
            items=self.items)


def list_apps_regions(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListAppsRegionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:listAppsRegions', __args__, opts=opts, typ=ListAppsRegionsResult).value

    return AwaitableListAppsRegionsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_apps_regions)
def list_apps_regions_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListAppsRegionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
