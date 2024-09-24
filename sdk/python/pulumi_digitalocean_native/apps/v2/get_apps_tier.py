# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = [
    'AppsGetTierResponse',
    'AwaitableAppsGetTierResponse',
    'get_apps_tier',
    'get_apps_tier_output',
]

@pulumi.output_type
class AppsGetTierResponse:
    def __init__(__self__, tier=None):
        if tier and not isinstance(tier, dict):
            raise TypeError("Expected argument 'tier' to be a dict")
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def tier(self) -> Optional['outputs.AppsTier']:
        return pulumi.get(self, "tier")


class AwaitableAppsGetTierResponse(AppsGetTierResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsGetTierResponse(
            tier=self.tier)


def get_apps_tier(slug: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsGetTierResponse:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the tier
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsTier', __args__, opts=opts, typ=AppsGetTierResponse).value

    return AwaitableAppsGetTierResponse(
        tier=pulumi.get(__ret__, 'tier'))
def get_apps_tier_output(slug: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AppsGetTierResponse]:
    """
    Use this data source to access information about an existing resource.

    :param str slug: The slug of the tier
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:getAppsTier', __args__, opts=opts, typ=AppsGetTierResponse)
    return __ret__.apply(lambda __response__: AppsGetTierResponse(
        tier=pulumi.get(__response__, 'tier')))
