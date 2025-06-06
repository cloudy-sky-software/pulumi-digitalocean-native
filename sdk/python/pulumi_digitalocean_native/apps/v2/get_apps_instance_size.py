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
    'AppsGetInstanceSizeResponse',
    'AwaitableAppsGetInstanceSizeResponse',
    'get_apps_instance_size',
    'get_apps_instance_size_output',
]

@pulumi.output_type
class AppsGetInstanceSizeResponse:
    def __init__(__self__, instance_size=None):
        if instance_size and not isinstance(instance_size, dict):
            raise TypeError("Expected argument 'instance_size' to be a dict")
        pulumi.set(__self__, "instance_size", instance_size)

    @property
    @pulumi.getter(name="instanceSize")
    def instance_size(self) -> Optional['outputs.AppsInstanceSize']:
        return pulumi.get(self, "instance_size")


class AwaitableAppsGetInstanceSizeResponse(AppsGetInstanceSizeResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsGetInstanceSizeResponse(
            instance_size=self.instance_size)


def get_apps_instance_size(slug: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsGetInstanceSizeResponse:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str slug: The slug of the instance size
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsInstanceSize', __args__, opts=opts, typ=AppsGetInstanceSizeResponse).value

    return AwaitableAppsGetInstanceSizeResponse(
        instance_size=pulumi.get(__ret__, 'instance_size'))
def get_apps_instance_size_output(slug: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AppsGetInstanceSizeResponse]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str slug: The slug of the instance size
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:getAppsInstanceSize', __args__, opts=opts, typ=AppsGetInstanceSizeResponse)
    return __ret__.apply(lambda __response__: AppsGetInstanceSizeResponse(
        instance_size=pulumi.get(__response__, 'instance_size')))
