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
from ._inputs import *

__all__ = ['UptimeAlertArgs', 'UptimeAlert']

@pulumi.input_type
class UptimeAlertArgs:
    def __init__(__self__, *,
                 check_id: Optional[pulumi.Input[builtins.str]] = None,
                 comparison: Optional[pulumi.Input['AlertUpdatableComparison']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notifications: Optional[pulumi.Input['NotificationArgs']] = None,
                 period: Optional[pulumi.Input['AlertUpdatablePeriod']] = None,
                 threshold: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input['AlertUpdatableType']] = None):
        """
        The set of arguments for constructing a UptimeAlert resource.
        :param pulumi.Input[builtins.str] check_id: A unique identifier for a check.
        :param pulumi.Input['AlertUpdatableComparison'] comparison: The comparison operator used against the alert's threshold.
        :param pulumi.Input[builtins.str] name: A human-friendly display name.
        :param pulumi.Input['NotificationArgs'] notifications: The notification settings for a trigger alert.
        :param pulumi.Input['AlertUpdatablePeriod'] period: Period of time the threshold must be exceeded to trigger the alert.
        :param pulumi.Input[builtins.int] threshold: The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        :param pulumi.Input['AlertUpdatableType'] type: The type of alert.
        """
        if check_id is not None:
            pulumi.set(__self__, "check_id", check_id)
        if comparison is not None:
            pulumi.set(__self__, "comparison", comparison)
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
    @pulumi.getter(name="checkId")
    def check_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique identifier for a check.
        """
        return pulumi.get(self, "check_id")

    @check_id.setter
    def check_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "check_id", value)

    @property
    @pulumi.getter
    def comparison(self) -> Optional[pulumi.Input['AlertUpdatableComparison']]:
        """
        The comparison operator used against the alert's threshold.
        """
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: Optional[pulumi.Input['AlertUpdatableComparison']]):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A human-friendly display name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input['NotificationArgs']]:
        """
        The notification settings for a trigger alert.
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input['NotificationArgs']]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input['AlertUpdatablePeriod']]:
        """
        Period of time the threshold must be exceeded to trigger the alert.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input['AlertUpdatablePeriod']]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter
    def threshold(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        """
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['AlertUpdatableType']]:
        """
        The type of alert.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['AlertUpdatableType']]):
        pulumi.set(self, "type", value)


class UptimeAlert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_id: Optional[pulumi.Input[builtins.str]] = None,
                 comparison: Optional[pulumi.Input['AlertUpdatableComparison']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notifications: Optional[pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']]] = None,
                 period: Optional[pulumi.Input['AlertUpdatablePeriod']] = None,
                 threshold: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input['AlertUpdatableType']] = None,
                 __props__=None):
        """
        Create a UptimeAlert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] check_id: A unique identifier for a check.
        :param pulumi.Input['AlertUpdatableComparison'] comparison: The comparison operator used against the alert's threshold.
        :param pulumi.Input[builtins.str] name: A human-friendly display name.
        :param pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']] notifications: The notification settings for a trigger alert.
        :param pulumi.Input['AlertUpdatablePeriod'] period: Period of time the threshold must be exceeded to trigger the alert.
        :param pulumi.Input[builtins.int] threshold: The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        :param pulumi.Input['AlertUpdatableType'] type: The type of alert.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UptimeAlertArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UptimeAlert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UptimeAlertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UptimeAlertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_id: Optional[pulumi.Input[builtins.str]] = None,
                 comparison: Optional[pulumi.Input['AlertUpdatableComparison']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notifications: Optional[pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']]] = None,
                 period: Optional[pulumi.Input['AlertUpdatablePeriod']] = None,
                 threshold: Optional[pulumi.Input[builtins.int]] = None,
                 type: Optional[pulumi.Input['AlertUpdatableType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UptimeAlertArgs.__new__(UptimeAlertArgs)

            __props__.__dict__["check_id"] = check_id
            __props__.__dict__["comparison"] = comparison
            __props__.__dict__["name"] = name
            __props__.__dict__["notifications"] = notifications
            __props__.__dict__["period"] = period
            __props__.__dict__["threshold"] = threshold
            __props__.__dict__["type"] = type
            __props__.__dict__["alert"] = None
        super(UptimeAlert, __self__).__init__(
            'digitalocean-native:uptime/v2:UptimeAlert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UptimeAlert':
        """
        Get an existing UptimeAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UptimeAlertArgs.__new__(UptimeAlertArgs)

        __props__.__dict__["alert"] = None
        __props__.__dict__["comparison"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notifications"] = None
        __props__.__dict__["period"] = None
        __props__.__dict__["threshold"] = None
        __props__.__dict__["type"] = None
        return UptimeAlert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alert(self) -> pulumi.Output[Optional['outputs.Alert']]:
        return pulumi.get(self, "alert")

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Output[Optional['AlertUpdatableComparison']]:
        """
        The comparison operator used against the alert's threshold.
        """
        return pulumi.get(self, "comparison")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A human-friendly display name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output['outputs.Notification']:
        """
        The notification settings for a trigger alert.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional['AlertUpdatablePeriod']]:
        """
        Period of time the threshold must be exceeded to trigger the alert.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
        """
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['AlertUpdatableType']:
        """
        The type of alert.
        """
        return pulumi.get(self, "type")

