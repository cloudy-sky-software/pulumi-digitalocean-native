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
    'GetMonitoringDropletLoad15MetricResult',
    'AwaitableGetMonitoringDropletLoad15MetricResult',
    'get_monitoring_droplet_load15_metric',
    'get_monitoring_droplet_load15_metric_output',
]

@pulumi.output_type
class GetMonitoringDropletLoad15MetricResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.Metrics':
        return pulumi.get(self, "items")


class AwaitableGetMonitoringDropletLoad15MetricResult(GetMonitoringDropletLoad15MetricResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMonitoringDropletLoad15MetricResult(
            items=self.items)


def get_monitoring_droplet_load15_metric(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitoringDropletLoad15MetricResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:getMonitoringDropletLoad15Metric', __args__, opts=opts, typ=GetMonitoringDropletLoad15MetricResult).value

    return AwaitableGetMonitoringDropletLoad15MetricResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_monitoring_droplet_load15_metric)
def get_monitoring_droplet_load15_metric_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMonitoringDropletLoad15MetricResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...