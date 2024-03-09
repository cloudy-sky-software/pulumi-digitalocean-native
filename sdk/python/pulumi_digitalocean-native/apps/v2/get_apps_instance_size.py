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
    'GetAppsInstanceSizeResult',
    'AwaitableGetAppsInstanceSizeResult',
    'get_apps_instance_size',
    'get_apps_instance_size_output',
]

@pulumi.output_type
class GetAppsInstanceSizeResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.AppsGetInstanceSizeResponse':
        return pulumi.get(self, "items")


class AwaitableGetAppsInstanceSizeResult(GetAppsInstanceSizeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppsInstanceSizeResult(
            items=self.items)


def get_apps_instance_size(slug: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppsInstanceSizeResult:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the instance size
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsInstanceSize', __args__, opts=opts, typ=GetAppsInstanceSizeResult).value

    return AwaitableGetAppsInstanceSizeResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_apps_instance_size)
def get_apps_instance_size_output(slug: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppsInstanceSizeResult]:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the instance size
    """
    ...
