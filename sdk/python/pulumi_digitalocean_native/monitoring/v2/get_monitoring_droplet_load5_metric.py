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
    'Metrics',
    'AwaitableMetrics',
    'get_monitoring_droplet_load5_metric',
    'get_monitoring_droplet_load5_metric_output',
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


def get_monitoring_droplet_load5_metric(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMetrics:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric', __args__, opts=opts, typ=Metrics).value

    return AwaitableMetrics(
        data=pulumi.get(__ret__, 'data'),
        status=pulumi.get(__ret__, 'status'))
def get_monitoring_droplet_load5_metric_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[Metrics]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:monitoring/v2:getMonitoringDropletLoad5Metric', __args__, opts=opts, typ=Metrics)
    return __ret__.apply(lambda __response__: Metrics(
        data=pulumi.get(__response__, 'data'),
        status=pulumi.get(__response__, 'status')))
