# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'AlertPolicy',
    'Alerts',
    'MetaMeta',
    'MetricsData',
    'MetricsResult',
    'PageLinks',
    'PageLinksPagesProperties',
    'SlackDetails',
]

@pulumi.output_type
class AlertPolicy(dict):
    def __init__(__self__, *,
                 alerts: 'outputs.Alerts',
                 compare: 'AlertPolicyCompare',
                 description: builtins.str,
                 enabled: builtins.bool,
                 entities: Sequence[builtins.str],
                 tags: Sequence[builtins.str],
                 type: 'AlertPolicyType',
                 uuid: builtins.str,
                 value: builtins.float,
                 window: 'AlertPolicyWindow'):
        pulumi.set(__self__, "alerts", alerts)
        pulumi.set(__self__, "compare", compare)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "entities", entities)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "uuid", uuid)
        pulumi.set(__self__, "value", value)
        pulumi.set(__self__, "window", window)

    @property
    @pulumi.getter
    def alerts(self) -> 'outputs.Alerts':
        return pulumi.get(self, "alerts")

    @property
    @pulumi.getter
    def compare(self) -> 'AlertPolicyCompare':
        return pulumi.get(self, "compare")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def entities(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> 'AlertPolicyType':
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> builtins.str:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def value(self) -> builtins.float:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def window(self) -> 'AlertPolicyWindow':
        return pulumi.get(self, "window")


@pulumi.output_type
class Alerts(dict):
    def __init__(__self__, *,
                 email: Sequence[builtins.str],
                 slack: Sequence['outputs.SlackDetails']):
        """
        :param Sequence[builtins.str] email: An email to notify on an alert trigger.
        :param Sequence['SlackDetails'] slack: Slack integration details.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "slack", slack)

    @property
    @pulumi.getter
    def email(self) -> Sequence[builtins.str]:
        """
        An email to notify on an alert trigger.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def slack(self) -> Sequence['outputs.SlackDetails']:
        """
        Slack integration details.
        """
        return pulumi.get(self, "slack")


@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[builtins.int] = None):
        """
        :param builtins.int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[builtins.int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class MetricsData(dict):
    def __init__(__self__, *,
                 result: Sequence['outputs.MetricsResult'],
                 result_type: 'MetricsDataResultType'):
        """
        :param Sequence['MetricsResult'] result: Result of query.
        """
        pulumi.set(__self__, "result", result)
        pulumi.set(__self__, "result_type", result_type)

    @property
    @pulumi.getter
    def result(self) -> Sequence['outputs.MetricsResult']:
        """
        Result of query.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="resultType")
    def result_type(self) -> 'MetricsDataResultType':
        return pulumi.get(self, "result_type")


@pulumi.output_type
class MetricsResult(dict):
    def __init__(__self__, *,
                 metric: Any,
                 values: Sequence[Sequence[Any]]):
        """
        :param Any metric: An object containing the metric labels.
        """
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def metric(self) -> Any:
        """
        An object containing the metric labels.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def values(self) -> Sequence[Sequence[Any]]:
        return pulumi.get(self, "values")


@pulumi.output_type
class PageLinks(dict):
    def __init__(__self__, *,
                 pages: Optional['outputs.PageLinksPagesProperties'] = None):
        if pages is not None:
            pulumi.set(__self__, "pages", pages)

    @property
    @pulumi.getter
    def pages(self) -> Optional['outputs.PageLinksPagesProperties']:
        return pulumi.get(self, "pages")


@pulumi.output_type
class PageLinksPagesProperties(dict):
    def __init__(__self__, *,
                 first: Optional[builtins.str] = None,
                 last: Optional[builtins.str] = None,
                 next: Optional[builtins.str] = None,
                 prev: Optional[builtins.str] = None):
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if prev is not None:
            pulumi.set(__self__, "prev", prev)

    @property
    @pulumi.getter
    def first(self) -> Optional[builtins.str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[builtins.str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[builtins.str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[builtins.str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class SlackDetails(dict):
    def __init__(__self__, *,
                 channel: builtins.str,
                 url: builtins.str):
        """
        :param builtins.str channel: Slack channel to notify of an alert trigger.
        :param builtins.str url: Slack Webhook URL.
        """
        pulumi.set(__self__, "channel", channel)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def channel(self) -> builtins.str:
        """
        Slack channel to notify of an alert trigger.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        """
        Slack Webhook URL.
        """
        return pulumi.get(self, "url")


