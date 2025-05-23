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

__all__ = ['UptimeCheckArgs', 'UptimeCheck']

@pulumi.input_type
class UptimeCheckArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]]] = None,
                 target: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['CheckUpdatableType']] = None):
        """
        The set of arguments for constructing a UptimeCheck resource.
        :param pulumi.Input[builtins.bool] enabled: A boolean value indicating whether the check is enabled/disabled.
        :param pulumi.Input[builtins.str] name: A human-friendly display name.
        :param pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]] regions: An array containing the selected regions to perform healthchecks from.
        :param pulumi.Input[builtins.str] target: The endpoint to perform healthchecks on.
        :param pulumi.Input['CheckUpdatableType'] type: The type of health check to perform.
        """
        if enabled is None:
            enabled = True
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
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
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean value indicating whether the check is enabled/disabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

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
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]]]:
        """
        An array containing the selected regions to perform healthchecks from.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The endpoint to perform healthchecks on.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['CheckUpdatableType']]:
        """
        The type of health check to perform.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['CheckUpdatableType']]):
        pulumi.set(self, "type", value)


class UptimeCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]]] = None,
                 target: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['CheckUpdatableType']] = None,
                 __props__=None):
        """
        Create a UptimeCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] enabled: A boolean value indicating whether the check is enabled/disabled.
        :param pulumi.Input[builtins.str] name: A human-friendly display name.
        :param pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]] regions: An array containing the selected regions to perform healthchecks from.
        :param pulumi.Input[builtins.str] target: The endpoint to perform healthchecks on.
        :param pulumi.Input['CheckUpdatableType'] type: The type of health check to perform.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UptimeCheckArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UptimeCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UptimeCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UptimeCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['CheckUpdatableRegionsItem']]]] = None,
                 target: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['CheckUpdatableType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UptimeCheckArgs.__new__(UptimeCheckArgs)

            if enabled is None:
                enabled = True
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["regions"] = regions
            __props__.__dict__["target"] = target
            __props__.__dict__["type"] = type
            __props__.__dict__["check"] = None
        super(UptimeCheck, __self__).__init__(
            'digitalocean-native:uptime/v2:UptimeCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UptimeCheck':
        """
        Get an existing UptimeCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UptimeCheckArgs.__new__(UptimeCheckArgs)

        __props__.__dict__["check"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["regions"] = None
        __props__.__dict__["target"] = None
        __props__.__dict__["type"] = None
        return UptimeCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def check(self) -> pulumi.Output[Optional['outputs.Check']]:
        return pulumi.get(self, "check")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[builtins.bool]:
        """
        A boolean value indicating whether the check is enabled/disabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A human-friendly display name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence['CheckUpdatableRegionsItem']]:
        """
        An array containing the selected regions to perform healthchecks from.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[builtins.str]:
        """
        The endpoint to perform healthchecks on.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['CheckUpdatableType']:
        """
        The type of health check to perform.
        """
        return pulumi.get(self, "type")

