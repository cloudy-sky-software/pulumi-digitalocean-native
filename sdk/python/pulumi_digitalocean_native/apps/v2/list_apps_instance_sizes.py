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
    'AppsListInstanceSizesResponse',
    'AwaitableAppsListInstanceSizesResponse',
    'list_apps_instance_sizes',
    'list_apps_instance_sizes_output',
]

@pulumi.output_type
class AppsListInstanceSizesResponse:
    def __init__(__self__, discount_percent=None, instance_sizes=None):
        if discount_percent and not isinstance(discount_percent, float):
            raise TypeError("Expected argument 'discount_percent' to be a float")
        pulumi.set(__self__, "discount_percent", discount_percent)
        if instance_sizes and not isinstance(instance_sizes, list):
            raise TypeError("Expected argument 'instance_sizes' to be a list")
        pulumi.set(__self__, "instance_sizes", instance_sizes)

    @property
    @pulumi.getter(name="discountPercent")
    def discount_percent(self) -> Optional[builtins.float]:
        return pulumi.get(self, "discount_percent")

    @property
    @pulumi.getter(name="instanceSizes")
    def instance_sizes(self) -> Optional[Sequence['outputs.AppsInstanceSize']]:
        return pulumi.get(self, "instance_sizes")


class AwaitableAppsListInstanceSizesResponse(AppsListInstanceSizesResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsListInstanceSizesResponse(
            discount_percent=self.discount_percent,
            instance_sizes=self.instance_sizes)


def list_apps_instance_sizes(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsListInstanceSizesResponse:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:listAppsInstanceSizes', __args__, opts=opts, typ=AppsListInstanceSizesResponse).value

    return AwaitableAppsListInstanceSizesResponse(
        discount_percent=pulumi.get(__ret__, 'discount_percent'),
        instance_sizes=pulumi.get(__ret__, 'instance_sizes'))
def list_apps_instance_sizes_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AppsListInstanceSizesResponse]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:listAppsInstanceSizes', __args__, opts=opts, typ=AppsListInstanceSizesResponse)
    return __ret__.apply(lambda __response__: AppsListInstanceSizesResponse(
        discount_percent=pulumi.get(__response__, 'discount_percent'),
        instance_sizes=pulumi.get(__response__, 'instance_sizes')))
