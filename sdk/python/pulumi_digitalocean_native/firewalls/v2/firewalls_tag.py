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
from ._inputs import *

__all__ = ['FirewallsTagArgs', 'FirewallsTag']

@pulumi.input_type
class FirewallsTagArgs:
    def __init__(__self__, *,
                 tags: pulumi.Input['TagsArgs'],
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a FirewallsTag resource.
        :param pulumi.Input[builtins.str] firewall_id: A unique ID that can be used to identify and reference a firewall.
        """
        pulumi.set(__self__, "tags", tags)
        if firewall_id is not None:
            pulumi.set(__self__, "firewall_id", firewall_id)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Input['TagsArgs']:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: pulumi.Input['TagsArgs']):
        pulumi.set(self, "tags", value)

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


class FirewallsTag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Union['TagsArgs', 'TagsArgsDict']]] = None,
                 __props__=None):
        """
        Create a FirewallsTag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] firewall_id: A unique ID that can be used to identify and reference a firewall.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallsTagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallsTag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallsTagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallsTagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Union['TagsArgs', 'TagsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallsTagArgs.__new__(FirewallsTagArgs)

            __props__.__dict__["firewall_id"] = firewall_id
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
        super(FirewallsTag, __self__).__init__(
            'digitalocean-native:firewalls/v2:FirewallsTag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FirewallsTag':
        """
        Get an existing FirewallsTag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FirewallsTagArgs.__new__(FirewallsTagArgs)

        __props__.__dict__["tags"] = None
        return FirewallsTag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output['outputs.Tags']:
        return pulumi.get(self, "tags")

