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

__all__ = ['AssignArgs', 'Assign']

@pulumi.input_type
class AssignArgs:
    def __init__(__self__, *,
                 droplet_id: pulumi.Input[builtins.int],
                 type: pulumi.Input['FloatingIPsActionType'],
                 floating_ip: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Assign resource.
        :param pulumi.Input[builtins.int] droplet_id: The ID of the Droplet that the floating IP will be assigned to.
        :param pulumi.Input['FloatingIPsActionType'] type: The type of action to initiate for the floating IP.
        :param pulumi.Input[builtins.str] floating_ip: A floating IP address.
        """
        pulumi.set(__self__, "droplet_id", droplet_id)
        pulumi.set(__self__, "type", type)
        if floating_ip is not None:
            pulumi.set(__self__, "floating_ip", floating_ip)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Input[builtins.int]:
        """
        The ID of the Droplet that the floating IP will be assigned to.
        """
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['FloatingIPsActionType']:
        """
        The type of action to initiate for the floating IP.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['FloatingIPsActionType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A floating IP address.
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "floating_ip", value)


class Assign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[builtins.int]] = None,
                 floating_ip: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['FloatingIPsActionType']] = None,
                 __props__=None):
        """
        Create a Assign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] droplet_id: The ID of the Droplet that the floating IP will be assigned to.
        :param pulumi.Input[builtins.str] floating_ip: A floating IP address.
        :param pulumi.Input['FloatingIPsActionType'] type: The type of action to initiate for the floating IP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Assign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AssignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[builtins.int]] = None,
                 floating_ip: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['FloatingIPsActionType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssignArgs.__new__(AssignArgs)

            if droplet_id is None and not opts.urn:
                raise TypeError("Missing required property 'droplet_id'")
            __props__.__dict__["droplet_id"] = droplet_id
            __props__.__dict__["floating_ip"] = floating_ip
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["action"] = None
        super(Assign, __self__).__init__(
            'digitalocean-native:floating_ips/v2:Assign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Assign':
        """
        Get an existing Assign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AssignArgs.__new__(AssignArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["droplet_id"] = None
        __props__.__dict__["type"] = None
        return Assign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['outputs.Action']]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The ID of the Droplet that the floating IP will be assigned to.
        """
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['FloatingIPsActionType']]:
        """
        The type of action to initiate for the floating IP.
        """
        return pulumi.get(self, "type")

