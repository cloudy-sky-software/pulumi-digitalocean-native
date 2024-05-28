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

__all__ = ['UnassignArgs', 'Unassign']

@pulumi.input_type
class UnassignArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['ReservedIpActionTypeType'],
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 reserved_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Unassign resource.
        :param pulumi.Input['ReservedIpActionTypeType'] type: The type of action to initiate for the reserved IP.
        :param pulumi.Input[int] droplet_id: The ID of the Droplet that the reserved IP will be unassigned from.
        :param pulumi.Input[str] reserved_ip: A reserved IP address.
        """
        pulumi.set(__self__, "type", type)
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)
        if reserved_ip is not None:
            pulumi.set(__self__, "reserved_ip", reserved_ip)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ReservedIpActionTypeType']:
        """
        The type of action to initiate for the reserved IP.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ReservedIpActionTypeType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the Droplet that the reserved IP will be unassigned from.
        """
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter(name="reservedIp")
    def reserved_ip(self) -> Optional[pulumi.Input[str]]:
        """
        A reserved IP address.
        """
        return pulumi.get(self, "reserved_ip")

    @reserved_ip.setter
    def reserved_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_ip", value)


class Unassign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 reserved_ip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['ReservedIpActionTypeType']] = None,
                 __props__=None):
        """
        Create a Unassign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] droplet_id: The ID of the Droplet that the reserved IP will be unassigned from.
        :param pulumi.Input[str] reserved_ip: A reserved IP address.
        :param pulumi.Input['ReservedIpActionTypeType'] type: The type of action to initiate for the reserved IP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UnassignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Unassign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UnassignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UnassignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 reserved_ip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['ReservedIpActionTypeType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UnassignArgs.__new__(UnassignArgs)

            __props__.__dict__["droplet_id"] = droplet_id
            __props__.__dict__["reserved_ip"] = reserved_ip
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["action"] = None
        super(Unassign, __self__).__init__(
            'digitalocean-native:reserved_ips/v2:Unassign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Unassign':
        """
        Get an existing Unassign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UnassignArgs.__new__(UnassignArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["droplet_id"] = None
        __props__.__dict__["type"] = None
        return Unassign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['outputs.Action']]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Output[int]:
        """
        The ID of the Droplet that the reserved IP will be unassigned from.
        """
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['ReservedIpActionTypeType']]:
        """
        The type of action to initiate for the reserved IP.
        """
        return pulumi.get(self, "type")
