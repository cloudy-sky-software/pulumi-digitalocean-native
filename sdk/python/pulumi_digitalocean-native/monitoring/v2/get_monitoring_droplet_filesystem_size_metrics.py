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
    'GetMonitoringDropletFilesystemSizeMetricsResult',
    'AwaitableGetMonitoringDropletFilesystemSizeMetricsResult',
    'get_monitoring_droplet_filesystem_size_metrics',
    'get_monitoring_droplet_filesystem_size_metrics_output',
]

@pulumi.output_type
class GetMonitoringDropletFilesystemSizeMetricsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.Metrics':
        return pulumi.get(self, "items")


class AwaitableGetMonitoringDropletFilesystemSizeMetricsResult(GetMonitoringDropletFilesystemSizeMetricsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMonitoringDropletFilesystemSizeMetricsResult(
            items=self.items)


def get_monitoring_droplet_filesystem_size_metrics(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitoringDropletFilesystemSizeMetricsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:getMonitoringDropletFilesystemSizeMetrics', __args__, opts=opts, typ=GetMonitoringDropletFilesystemSizeMetricsResult).value

    return AwaitableGetMonitoringDropletFilesystemSizeMetricsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_monitoring_droplet_filesystem_size_metrics)
def get_monitoring_droplet_filesystem_size_metrics_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMonitoringDropletFilesystemSizeMetricsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
