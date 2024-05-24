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

__all__ = ['PowerOffArgs', 'PowerOff']

@pulumi.input_type
class PowerOffArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['Type'],
                 droplet_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PowerOff resource.
        :param pulumi.Input['Type'] type: The type of action to initiate for the Droplet.
        :param pulumi.Input[str] droplet_id: A unique identifier for a Droplet instance.
        """
        pulumi.set(__self__, "type", type)
        if droplet_id is not None:
            pulumi.set(__self__, "droplet_id", droplet_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['Type']:
        """
        The type of action to initiate for the Droplet.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['Type']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for a Droplet instance.
        """
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "droplet_id", value)


class PowerOff(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        """
        Specifies the action that will be taken on the Droplet.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] droplet_id: A unique identifier for a Droplet instance.
        :param pulumi.Input['Type'] type: The type of action to initiate for the Droplet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PowerOffArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies the action that will be taken on the Droplet.

        :param str resource_name: The name of the resource.
        :param PowerOffArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PowerOffArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PowerOffArgs.__new__(PowerOffArgs)

            __props__.__dict__["droplet_id"] = droplet_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["action"] = None
        super(PowerOff, __self__).__init__(
            'digitalocean-native:droplets/v2:PowerOff',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PowerOff':
        """
        Get an existing PowerOff resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PowerOffArgs.__new__(PowerOffArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["type"] = None
        return PowerOff(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['outputs.Action']]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['Type']:
        """
        The type of action to initiate for the Droplet.
        """
        return pulumi.get(self, "type")

