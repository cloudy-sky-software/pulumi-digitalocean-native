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
    'Alert',
    'Check',
    'GetUptimeAlertProperties',
    'GetUptimeCheckProperties',
    'GetUptimeCheckStateProperties',
    'ListUptimeAlertsItems',
    'ListUptimeChecksItems',
    'MetaMeta',
    'Notification',
    'NotificationSlackItemProperties',
    'PageLinks',
    'PageLinksPagesProperties',
    'PreviousOutage',
    'RegionState',
    'RegionalState',
    'State',
]

@pulumi.output_type
class Alert(dict):
    def __init__(__self__, *,
                 comparison: Optional['AlertUpdatableComparison'] = None,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 notifications: Optional['outputs.Notification'] = None,
                 period: Optional['AlertUpdatablePeriod'] = None,
                 threshold: Optional[int] = None,
                 type: Optional['AlertUpdatableType'] = None):
        """
        :param 'AlertUpdatableComparison' comparison: The comparison operator used against the alert's threshold.
        :param str id: A unique ID that can be used to identify and reference the alert.
        :param str name: A human-friendly display name.
        :param 'Notification' notifications: The notification settings for a trigger alert.
        :param 'AlertUpdatablePeriod' period: Period of time the threshold must be exceeded to trigger the alert.
        :param int threshold: The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        :param 'AlertUpdatableType' type: The type of alert.
        """
        if comparison is not None:
            pulumi.set(__self__, "comparison", comparison)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if threshold is not None:
            pulumi.set(__self__, "threshold", threshold)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def comparison(self) -> Optional['AlertUpdatableComparison']:
        """
        The comparison operator used against the alert's threshold.
        """
        return pulumi.get(self, "comparison")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        A unique ID that can be used to identify and reference the alert.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A human-friendly display name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> Optional['outputs.Notification']:
        """
        The notification settings for a trigger alert.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter
    def period(self) -> Optional['AlertUpdatablePeriod']:
        """
        Period of time the threshold must be exceeded to trigger the alert.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def threshold(self) -> Optional[int]:
        """
        The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        """
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter
    def type(self) -> Optional['AlertUpdatableType']:
        """
        The type of alert.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class Check(dict):
    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 regions: Optional[Sequence['CheckUpdatableRegionsItem']] = None,
                 target: Optional[str] = None,
                 type: Optional['CheckUpdatableType'] = None):
        """
        :param bool enabled: A boolean value indicating whether the check is enabled/disabled.
        :param str id: A unique ID that can be used to identify and reference the check.
        :param str name: A human-friendly display name.
        :param Sequence['CheckUpdatableRegionsItem'] regions: An array containing the selected regions to perform healthchecks from.
        :param str target: The endpoint to perform healthchecks on.
        :param 'CheckUpdatableType' type: The type of health check to perform.
        """
        if enabled is None:
            enabled = True
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        A boolean value indicating whether the check is enabled/disabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        A unique ID that can be used to identify and reference the check.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A human-friendly display name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> Optional[Sequence['CheckUpdatableRegionsItem']]:
        """
        An array containing the selected regions to perform healthchecks from.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        The endpoint to perform healthchecks on.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def type(self) -> Optional['CheckUpdatableType']:
        """
        The type of health check to perform.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetUptimeAlertProperties(dict):
    def __init__(__self__, *,
                 alert: Optional['outputs.Alert'] = None):
        if alert is not None:
            pulumi.set(__self__, "alert", alert)

    @property
    @pulumi.getter
    def alert(self) -> Optional['outputs.Alert']:
        return pulumi.get(self, "alert")


@pulumi.output_type
class GetUptimeCheckProperties(dict):
    def __init__(__self__, *,
                 check: Optional['outputs.Check'] = None):
        if check is not None:
            pulumi.set(__self__, "check", check)

    @property
    @pulumi.getter
    def check(self) -> Optional['outputs.Check']:
        return pulumi.get(self, "check")


@pulumi.output_type
class GetUptimeCheckStateProperties(dict):
    def __init__(__self__, *,
                 state: Optional['outputs.State'] = None):
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def state(self) -> Optional['outputs.State']:
        return pulumi.get(self, "state")


@pulumi.output_type
class ListUptimeAlertsItems(dict):
    def __init__(__self__, *,
                 meta: 'outputs.MetaMeta',
                 alerts: Optional[Sequence['outputs.Alert']] = None,
                 links: Optional['outputs.PageLinks'] = None):
        pulumi.set(__self__, "meta", meta)
        if alerts is not None:
            pulumi.set(__self__, "alerts", alerts)
        if links is not None:
            pulumi.set(__self__, "links", links)

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def alerts(self) -> Optional[Sequence['outputs.Alert']]:
        return pulumi.get(self, "alerts")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")


@pulumi.output_type
class ListUptimeChecksItems(dict):
    def __init__(__self__, *,
                 meta: 'outputs.MetaMeta',
                 checks: Optional[Sequence['outputs.Check']] = None,
                 links: Optional['outputs.PageLinks'] = None):
        pulumi.set(__self__, "meta", meta)
        if checks is not None:
            pulumi.set(__self__, "checks", checks)
        if links is not None:
            pulumi.set(__self__, "links", links)

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def checks(self) -> Optional[Sequence['outputs.Check']]:
        return pulumi.get(self, "checks")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")


@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[int] = None):
        """
        :param int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class Notification(dict):
    """
    The notification settings for a trigger alert.
    """
    def __init__(__self__, *,
                 email: Sequence[str],
                 slack: Sequence['outputs.NotificationSlackItemProperties']):
        """
        The notification settings for a trigger alert.
        :param Sequence[str] email: An email to notify on an alert trigger.
        :param Sequence['NotificationSlackItemProperties'] slack: Slack integration details.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "slack", slack)

    @property
    @pulumi.getter
    def email(self) -> Sequence[str]:
        """
        An email to notify on an alert trigger.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def slack(self) -> Sequence['outputs.NotificationSlackItemProperties']:
        """
        Slack integration details.
        """
        return pulumi.get(self, "slack")


@pulumi.output_type
class NotificationSlackItemProperties(dict):
    def __init__(__self__, *,
                 channel: str,
                 url: str):
        """
        :param str channel: Slack channel to notify of an alert trigger.
        :param str url: Slack Webhook URL.
        """
        pulumi.set(__self__, "channel", channel)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def channel(self) -> str:
        """
        Slack channel to notify of an alert trigger.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        Slack Webhook URL.
        """
        return pulumi.get(self, "url")


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
                 first: Optional[str] = None,
                 last: Optional[str] = None,
                 next: Optional[str] = None,
                 prev: Optional[str] = None):
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
    def first(self) -> Optional[str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[str]:
        return pulumi.get(self, "prev")


@pulumi.output_type
class PreviousOutage(dict):
    def __init__(__self__, *,
                 duration_seconds: Optional[int] = None,
                 ended_at: Optional[str] = None,
                 region: Optional[str] = None,
                 started_at: Optional[str] = None):
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if ended_at is not None:
            pulumi.set(__self__, "ended_at", ended_at)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if started_at is not None:
            pulumi.set(__self__, "started_at", started_at)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[int]:
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter(name="endedAt")
    def ended_at(self) -> Optional[str]:
        return pulumi.get(self, "ended_at")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="startedAt")
    def started_at(self) -> Optional[str]:
        return pulumi.get(self, "started_at")


@pulumi.output_type
class RegionState(dict):
    def __init__(__self__, *,
                 status: Optional['RegionStateStatus'] = None,
                 status_changed_at: Optional[str] = None,
                 thirty_day_uptime_percentage: Optional[float] = None):
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_changed_at is not None:
            pulumi.set(__self__, "status_changed_at", status_changed_at)
        if thirty_day_uptime_percentage is not None:
            pulumi.set(__self__, "thirty_day_uptime_percentage", thirty_day_uptime_percentage)

    @property
    @pulumi.getter
    def status(self) -> Optional['RegionStateStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusChangedAt")
    def status_changed_at(self) -> Optional[str]:
        return pulumi.get(self, "status_changed_at")

    @property
    @pulumi.getter(name="thirtyDayUptimePercentage")
    def thirty_day_uptime_percentage(self) -> Optional[float]:
        return pulumi.get(self, "thirty_day_uptime_percentage")


@pulumi.output_type
class RegionalState(dict):
    """
    A map of region to regional state
    """
    def __init__(__self__, *,
                 eu_west: Optional['outputs.RegionState'] = None,
                 us_east: Optional['outputs.RegionState'] = None):
        """
        A map of region to regional state
        """
        if eu_west is not None:
            pulumi.set(__self__, "eu_west", eu_west)
        if us_east is not None:
            pulumi.set(__self__, "us_east", us_east)

    @property
    @pulumi.getter(name="euWest")
    def eu_west(self) -> Optional['outputs.RegionState']:
        return pulumi.get(self, "eu_west")

    @property
    @pulumi.getter(name="usEast")
    def us_east(self) -> Optional['outputs.RegionState']:
        return pulumi.get(self, "us_east")


@pulumi.output_type
class State(dict):
    def __init__(__self__, *,
                 previous_outage: Optional['outputs.PreviousOutage'] = None,
                 regions: Optional['outputs.RegionalState'] = None):
        """
        :param 'RegionalState' regions: A map of region to regional state
        """
        if previous_outage is not None:
            pulumi.set(__self__, "previous_outage", previous_outage)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)

    @property
    @pulumi.getter(name="previousOutage")
    def previous_outage(self) -> Optional['outputs.PreviousOutage']:
        return pulumi.get(self, "previous_outage")

    @property
    @pulumi.getter
    def regions(self) -> Optional['outputs.RegionalState']:
        """
        A map of region to regional state
        """
        return pulumi.get(self, "regions")


