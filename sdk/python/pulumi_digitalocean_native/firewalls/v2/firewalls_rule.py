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

__all__ = ['FirewallsRuleArgs', 'FirewallsRule']

@pulumi.input_type
class FirewallsRuleArgs:
    def __init__(__self__, *,
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]] = None):
        """
        The set of arguments for constructing a FirewallsRule resource.
        :param pulumi.Input[builtins.str] firewall_id: A unique ID that can be used to identify and reference a firewall.
        """
        if firewall_id is not None:
            pulumi.set(__self__, "firewall_id", firewall_id)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)

    @property
    @pulumi.getter(name="firewallId")
    def firewall_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique ID that can be used to identify and reference a firewall.
        """
        return pulumi.get(self, "firewall_id")

    @firewall_id.setter
    def firewall_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "firewall_id", value)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]]:
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesInboundRulesItemArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]]:
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRulesOutboundRulesItemArgs']]]]):
        pulumi.set(self, "outbound_rules", value)


class FirewallsRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRulesInboundRulesItemArgs', 'FirewallRulesInboundRulesItemArgsDict']]]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRulesOutboundRulesItemArgs', 'FirewallRulesOutboundRulesItemArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a FirewallsRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] firewall_id: A unique ID that can be used to identify and reference a firewall.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallsRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallsRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallsRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallsRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRulesInboundRulesItemArgs', 'FirewallRulesInboundRulesItemArgsDict']]]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRulesOutboundRulesItemArgs', 'FirewallRulesOutboundRulesItemArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallsRuleArgs.__new__(FirewallsRuleArgs)

            __props__.__dict__["firewall_id"] = firewall_id
            __props__.__dict__["inbound_rules"] = inbound_rules
            __props__.__dict__["outbound_rules"] = outbound_rules
        super(FirewallsRule, __self__).__init__(
            'digitalocean-native:firewalls/v2:FirewallsRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FirewallsRule':
        """
        Get an existing FirewallsRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FirewallsRuleArgs.__new__(FirewallsRuleArgs)

        __props__.__dict__["inbound_rules"] = None
        __props__.__dict__["outbound_rules"] = None
        return FirewallsRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallRulesInboundRulesItem']]]:
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallRulesOutboundRulesItem']]]:
        return pulumi.get(self, "outbound_rules")

