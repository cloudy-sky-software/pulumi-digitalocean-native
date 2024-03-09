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

__all__ = ['VolumeActionsByIdDetachArgs', 'VolumeActionsByIdDetach']

@pulumi.input_type
class VolumeActionsByIdDetachArgs:
    def __init__(__self__, *,
                 droplet_id: pulumi.Input[int],
                 type: pulumi.Input['VolumeActionPostBaseType'],
                 region: Optional[pulumi.Input['VolumeActionPostBaseRegion']] = None,
                 volume_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VolumeActionsByIdDetach resource.
        :param pulumi.Input[int] droplet_id: The unique identifier for the Droplet the volume will be attached or detached from.
        :param pulumi.Input['VolumeActionPostBaseType'] type: The volume action to initiate.
        :param pulumi.Input['VolumeActionPostBaseRegion'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input[str] volume_id: The ID of the block storage volume.
        """
        pulumi.set(__self__, "droplet_id", droplet_id)
        pulumi.set(__self__, "type", type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Input[int]:
        """
        The unique identifier for the Droplet the volume will be attached or detached from.
        """
        return pulumi.get(self, "droplet_id")

    @droplet_id.setter
    def droplet_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "droplet_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['VolumeActionPostBaseType']:
        """
        The volume action to initiate.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['VolumeActionPostBaseType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input['VolumeActionPostBaseRegion']]:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input['VolumeActionPostBaseRegion']]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the block storage volume.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_id", value)


class VolumeActionsByIdDetach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input['VolumeActionPostBaseRegion']] = None,
                 type: Optional[pulumi.Input['VolumeActionPostBaseType']] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VolumeActionsByIdDetach resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] droplet_id: The unique identifier for the Droplet the volume will be attached or detached from.
        :param pulumi.Input['VolumeActionPostBaseRegion'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input['VolumeActionPostBaseType'] type: The volume action to initiate.
        :param pulumi.Input[str] volume_id: The ID of the block storage volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeActionsByIdDetachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VolumeActionsByIdDetach resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VolumeActionsByIdDetachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeActionsByIdDetachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input['VolumeActionPostBaseRegion']] = None,
                 type: Optional[pulumi.Input['VolumeActionPostBaseType']] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeActionsByIdDetachArgs.__new__(VolumeActionsByIdDetachArgs)

            if droplet_id is None and not opts.urn:
                raise TypeError("Missing required property 'droplet_id'")
            __props__.__dict__["droplet_id"] = droplet_id
            __props__.__dict__["region"] = region
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["volume_id"] = volume_id
            __props__.__dict__["action"] = None
        super(VolumeActionsByIdDetach, __self__).__init__(
            'digitalocean-native:volumes/v2:VolumeActionsByIdDetach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VolumeActionsByIdDetach':
        """
        Get an existing VolumeActionsByIdDetach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VolumeActionsByIdDetachArgs.__new__(VolumeActionsByIdDetachArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["droplet_id"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["type"] = None
        return VolumeActionsByIdDetach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['outputs.VolumeAction']]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Output[Optional[int]]:
        """
        The unique identifier for the Droplet the volume will be attached or detached from.
        """
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional['VolumeActionPostBaseRegion']]:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['VolumeActionPostBaseType']]:
        """
        The volume action to initiate.
        """
        return pulumi.get(self, "type")
