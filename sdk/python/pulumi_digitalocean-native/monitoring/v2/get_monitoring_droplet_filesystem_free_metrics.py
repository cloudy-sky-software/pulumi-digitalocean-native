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
    'GetMonitoringDropletFilesystemFreeMetricsResult',
    'AwaitableGetMonitoringDropletFilesystemFreeMetricsResult',
    'get_monitoring_droplet_filesystem_free_metrics',
    'get_monitoring_droplet_filesystem_free_metrics_output',
]

@pulumi.output_type
class GetMonitoringDropletFilesystemFreeMetricsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.Metrics':
        return pulumi.get(self, "items")


class AwaitableGetMonitoringDropletFilesystemFreeMetricsResult(GetMonitoringDropletFilesystemFreeMetricsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMonitoringDropletFilesystemFreeMetricsResult(
            items=self.items)


def get_monitoring_droplet_filesystem_free_metrics(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitoringDropletFilesystemFreeMetricsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemFreeMetrics', __args__, opts=opts, typ=GetMonitoringDropletFilesystemFreeMetricsResult).value

    return AwaitableGetMonitoringDropletFilesystemFreeMetricsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_monitoring_droplet_filesystem_free_metrics)
def get_monitoring_droplet_filesystem_free_metrics_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMonitoringDropletFilesystemFreeMetricsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
