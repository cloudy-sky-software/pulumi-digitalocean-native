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
    'ListUptimeAlertsItems',
    'AwaitableListUptimeAlertsItems',
    'list_uptime_alerts',
    'list_uptime_alerts_output',
]

@pulumi.output_type
class ListUptimeAlertsItems:
    def __init__(__self__, alerts=None, links=None, meta=None):
        if alerts and not isinstance(alerts, list):
            raise TypeError("Expected argument 'alerts' to be a list")
        pulumi.set(__self__, "alerts", alerts)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def alerts(self) -> Optional[Sequence['outputs.Alert']]:
        return pulumi.get(self, "alerts")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListUptimeAlertsItems(ListUptimeAlertsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListUptimeAlertsItems(
            alerts=self.alerts,
            links=self.links,
            meta=self.meta)


def list_uptime_alerts(check_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListUptimeAlertsItems:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    __args__ = dict()
    __args__['checkId'] = check_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:uptime/v2:listUptimeAlerts', __args__, opts=opts, typ=ListUptimeAlertsItems).value

    return AwaitableListUptimeAlertsItems(
        alerts=pulumi.get(__ret__, 'alerts'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))


@_utilities.lift_output_func(list_uptime_alerts)
def list_uptime_alerts_output(check_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListUptimeAlertsItems]:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    ...
