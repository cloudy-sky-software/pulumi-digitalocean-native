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
    'GetAppsDeploymentResult',
    'AwaitableGetAppsDeploymentResult',
    'get_apps_deployment',
    'get_apps_deployment_output',
]

@pulumi.output_type
class GetAppsDeploymentResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.AppsDeploymentResponse':
        return pulumi.get(self, "items")


class AwaitableGetAppsDeploymentResult(GetAppsDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppsDeploymentResult(
            items=self.items)


def get_apps_deployment(app_id: Optional[str] = None,
                        deployment_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppsDeploymentResult:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    :param str deployment_id: The deployment ID
    """
    __args__ = dict()
    __args__['appId'] = app_id
    __args__['deploymentId'] = deployment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsDeployment', __args__, opts=opts, typ=GetAppsDeploymentResult).value

    return AwaitableGetAppsDeploymentResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_apps_deployment)
def get_apps_deployment_output(app_id: Optional[pulumi.Input[str]] = None,
                               deployment_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppsDeploymentResult]:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    :param str deployment_id: The deployment ID
    """
    ...
