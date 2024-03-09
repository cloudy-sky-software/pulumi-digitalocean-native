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
    'ListUptimeAlertsResult',
    'AwaitableListUptimeAlertsResult',
    'list_uptime_alerts',
    'list_uptime_alerts_output',
]

@pulumi.output_type
class ListUptimeAlertsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListUptimeAlerts':
        return pulumi.get(self, "items")


class AwaitableListUptimeAlertsResult(ListUptimeAlertsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListUptimeAlertsResult(
            items=self.items)


def list_uptime_alerts(check_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListUptimeAlertsResult:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    __args__ = dict()
    __args__['checkId'] = check_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:uptime/v2:listUptimeAlerts', __args__, opts=opts, typ=ListUptimeAlertsResult).value

    return AwaitableListUptimeAlertsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_uptime_alerts)
def list_uptime_alerts_output(check_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListUptimeAlertsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str check_id: A unique identifier for a check.
    """
    ...