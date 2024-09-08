# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from ... import _utilities
from . import outputs

__all__ = [
    'AppsListTiersResponse',
    'AwaitableAppsListTiersResponse',
    'list_apps_tiers',
    'list_apps_tiers_output',
]

@pulumi.output_type
class AppsListTiersResponse:
    def __init__(__self__, tiers=None):
        if tiers and not isinstance(tiers, list):
            raise TypeError("Expected argument 'tiers' to be a list")
        pulumi.set(__self__, "tiers", tiers)

    @property
    @pulumi.getter
    def tiers(self) -> Optional[Sequence['outputs.AppsTier']]:
        return pulumi.get(self, "tiers")


class AwaitableAppsListTiersResponse(AppsListTiersResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsListTiersResponse(
            tiers=self.tiers)


def list_apps_tiers(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsListTiersResponse:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:listAppsTiers', __args__, opts=opts, typ=AppsListTiersResponse).value

    return AwaitableAppsListTiersResponse(
        tiers=pulumi.get(__ret__, 'tiers'))


@_utilities.lift_output_func(list_apps_tiers)
def list_apps_tiers_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AppsListTiersResponse]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
