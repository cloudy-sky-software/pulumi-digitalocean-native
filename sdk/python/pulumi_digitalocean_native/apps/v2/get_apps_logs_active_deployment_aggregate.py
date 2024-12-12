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

__all__ = [
    'AppsGetLogsResponse',
    'AwaitableAppsGetLogsResponse',
    'get_apps_logs_active_deployment_aggregate',
    'get_apps_logs_active_deployment_aggregate_output',
]

@pulumi.output_type
class AppsGetLogsResponse:
    def __init__(__self__, historic_urls=None, live_url=None):
        if historic_urls and not isinstance(historic_urls, list):
            raise TypeError("Expected argument 'historic_urls' to be a list")
        pulumi.set(__self__, "historic_urls", historic_urls)
        if live_url and not isinstance(live_url, str):
            raise TypeError("Expected argument 'live_url' to be a str")
        pulumi.set(__self__, "live_url", live_url)

    @property
    @pulumi.getter(name="historicUrls")
    def historic_urls(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "historic_urls")

    @property
    @pulumi.getter(name="liveUrl")
    def live_url(self) -> Optional[str]:
        """
        A URL of the real-time live logs. This URL may use either the `https://` or `wss://` protocols and will keep pushing live logs as they become available.
        """
        return pulumi.get(self, "live_url")


class AwaitableAppsGetLogsResponse(AppsGetLogsResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsGetLogsResponse(
            historic_urls=self.historic_urls,
            live_url=self.live_url)


def get_apps_logs_active_deployment_aggregate(app_id: Optional[str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsGetLogsResponse:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    """
    __args__ = dict()
    __args__['appId'] = app_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:getAppsLogsActiveDeploymentAggregate', __args__, opts=opts, typ=AppsGetLogsResponse).value

    return AwaitableAppsGetLogsResponse(
        historic_urls=pulumi.get(__ret__, 'historic_urls'),
        live_url=pulumi.get(__ret__, 'live_url'))
def get_apps_logs_active_deployment_aggregate_output(app_id: Optional[pulumi.Input[str]] = None,
                                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AppsGetLogsResponse]:
    """
    Use this data source to access information about an existing resource.

    :param str app_id: The app ID
    """
    __args__ = dict()
    __args__['appId'] = app_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:getAppsLogsActiveDeploymentAggregate', __args__, opts=opts, typ=AppsGetLogsResponse)
    return __ret__.apply(lambda __response__: AppsGetLogsResponse(
        historic_urls=pulumi.get(__response__, 'historic_urls'),
        live_url=pulumi.get(__response__, 'live_url')))
