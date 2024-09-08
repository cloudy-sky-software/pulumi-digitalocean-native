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
    'ListMonitoringAlertPolicyItems',
    'AwaitableListMonitoringAlertPolicyItems',
    'list_monitoring_alert_policy',
    'list_monitoring_alert_policy_output',
]

@pulumi.output_type
class ListMonitoringAlertPolicyItems:
    def __init__(__self__, links=None, meta=None, policies=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.AlertPolicy']:
        return pulumi.get(self, "policies")


class AwaitableListMonitoringAlertPolicyItems(ListMonitoringAlertPolicyItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListMonitoringAlertPolicyItems(
            links=self.links,
            meta=self.meta,
            policies=self.policies)


def list_monitoring_alert_policy(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListMonitoringAlertPolicyItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:monitoring/v2:listMonitoringAlertPolicy', __args__, opts=opts, typ=ListMonitoringAlertPolicyItems).value

    return AwaitableListMonitoringAlertPolicyItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        policies=pulumi.get(__ret__, 'policies'))


@_utilities.lift_output_func(list_monitoring_alert_policy)
def list_monitoring_alert_policy_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListMonitoringAlertPolicyItems]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
