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
from ._enums import *

__all__ = [
    'Metrics',
    'AwaitableMetrics',
    'get_monitoring_droplet_memory_total_metric',
    'get_monitoring_droplet_memory_total_metric_output',
]

@pulumi.output_type
class Metrics:
    def __init__(__self__, data=None, status=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def data(self) -> 'outputs.MetricsData':
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def status(self) -> 'MetricsStatus':
        return pulumi.get(self, "status")


class AwaitableMetrics(Metrics):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Metrics(
            data=self.data,
            status=self.status)


def get_monitoring_droplet_memory_total_metric(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMetrics:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:getMonitoringDropletMemoryTotalMetric', __args__, opts=opts, typ=Metrics).value

    return AwaitableMetrics(
        data=pulumi.get(__ret__, 'data'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_monitoring_droplet_memory_total_metric)
def get_monitoring_droplet_memory_total_metric_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Metrics]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
