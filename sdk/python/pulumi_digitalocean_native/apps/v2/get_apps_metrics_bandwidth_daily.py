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
    'AppMetricsBandwidthUsage',
    'AwaitableAppMetricsBandwidthUsage',
    'get_apps_metrics_bandwidth_daily',
    'get_apps_metrics_bandwidth_daily_output',
]

@pulumi.output_type
class AppMetricsBandwidthUsage:
    def __init__(__self__, app_bandwidth_usage=None, date=None):
        if app_bandwidth_usage and not isinstance(app_bandwidth_usage, list):
            raise TypeError("Expected argument 'app_bandwidth_usage' to be a list")
        pulumi.set(__self__, "app_bandwidth_usage", app_bandwidth_usage)
        if date and not isinstance(date, str):
            raise TypeError("Expected argument 'date' to be a str")
        pulumi.set(__self__, "date", date)

    @property
    @pulumi.getter(name="appBandwidthUsage")
    def app_bandwidth_usage(self) -> Optional[Sequence['outputs.AppMetricsBandwidthUsageDetails']]:
        """
        A list of bandwidth usage details by app.
        """
        return pulumi.get(self, "app_bandwidth_usage")

    @property
    @pulumi.getter
    def date(self) -> Optional[str]:
        """
        The date for the metrics data.
        """
        return pulumi.get(self, "date")


class AwaitableAppMetricsBandwidthUsage(AppMetricsBandwidthUsage):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppMetricsBandwidthUsage(
            app_bandwidth_usage=self.app_bandwidth_usage,
            date=self.date)


def get_apps_metrics_bandwidth_daily(app_id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppMetricsBandwidthUsage:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    """
    __args__ = dict()
    __args__['appId'] = app_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsMetricsBandwidthDaily', __args__, opts=opts, typ=AppMetricsBandwidthUsage).value

    return AwaitableAppMetricsBandwidthUsage(
        app_bandwidth_usage=pulumi.get(__ret__, 'app_bandwidth_usage'),
        date=pulumi.get(__ret__, 'date'))
def get_apps_metrics_bandwidth_daily_output(app_id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AppMetricsBandwidthUsage]:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    """
    __args__ = dict()
    __args__['appId'] = app_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:getAppsMetricsBandwidthDaily', __args__, opts=opts, typ=AppMetricsBandwidthUsage)
    return __ret__.apply(lambda __response__: AppMetricsBandwidthUsage(
        app_bandwidth_usage=pulumi.get(__response__, 'app_bandwidth_usage'),
        date=pulumi.get(__response__, 'date')))
