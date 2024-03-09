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

__all__ = ['DropletActionsByTagDisableBackupsArgs', 'DropletActionsByTagDisableBackups']

@pulumi.input_type
class DropletActionsByTagDisableBackupsArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['Type']):
        """
        The set of arguments for constructing a DropletActionsByTagDisableBackups resource.
        :param pulumi.Input['Type'] type: The type of action to initiate for the Droplet.
        """
        pulumi.set(__self__, "type", type)

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


class DropletActionsByTagDisableBackups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        """
        Specifies the action that will be taken on the Droplet.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Type'] type: The type of action to initiate for the Droplet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DropletActionsByTagDisableBackupsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies the action that will be taken on the Droplet.

        :param str resource_name: The name of the resource.
        :param DropletActionsByTagDisableBackupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DropletActionsByTagDisableBackupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DropletActionsByTagDisableBackupsArgs.__new__(DropletActionsByTagDisableBackupsArgs)

            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["actions"] = None
        super(DropletActionsByTagDisableBackups, __self__).__init__(
            'digitalocean-native:droplets/v2:DropletActionsByTagDisableBackups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DropletActionsByTagDisableBackups':
        """
        Get an existing DropletActionsByTagDisableBackups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DropletActionsByTagDisableBackupsArgs.__new__(DropletActionsByTagDisableBackupsArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["type"] = None
        return DropletActionsByTagDisableBackups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[Sequence['outputs.Action']]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['Type']:
        """
        The type of action to initiate for the Droplet.
        """
        return pulumi.get(self, "type")

