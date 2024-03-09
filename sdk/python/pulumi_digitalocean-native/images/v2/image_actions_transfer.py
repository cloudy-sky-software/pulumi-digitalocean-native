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

__all__ = ['ImageActionsTransferArgs', 'ImageActionsTransfer']

@pulumi.input_type
class ImageActionsTransferArgs:
    def __init__(__self__, *,
                 region: pulumi.Input['ImageActionsTransferPropertiesRegion'],
                 type: pulumi.Input['ImageActionBaseType'],
                 image_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ImageActionsTransfer resource.
        :param pulumi.Input['ImageActionsTransferPropertiesRegion'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input['ImageActionBaseType'] type: The action to be taken on the image. Can be either `convert` or `transfer`.
        :param pulumi.Input[str] image_id: A unique number that can be used to identify and reference a specific image.
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "type", type)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input['ImageActionsTransferPropertiesRegion']:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input['ImageActionsTransferPropertiesRegion']):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ImageActionBaseType']:
        """
        The action to be taken on the image. Can be either `convert` or `transfer`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ImageActionBaseType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique number that can be used to identify and reference a specific image.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)


class ImageActionsTransfer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['ImageActionsTransferPropertiesRegion']] = None,
                 type: Optional[pulumi.Input['ImageActionBaseType']] = None,
                 __props__=None):
        """
        Create a ImageActionsTransfer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: A unique number that can be used to identify and reference a specific image.
        :param pulumi.Input['ImageActionsTransferPropertiesRegion'] region: The slug identifier for the region where the resource will initially be  available.
        :param pulumi.Input['ImageActionBaseType'] type: The action to be taken on the image. Can be either `convert` or `transfer`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageActionsTransferArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ImageActionsTransfer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ImageActionsTransferArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageActionsTransferArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['ImageActionsTransferPropertiesRegion']] = None,
                 type: Optional[pulumi.Input['ImageActionBaseType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageActionsTransferArgs.__new__(ImageActionsTransferArgs)

            __props__.__dict__["image_id"] = image_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["completed_at"] = None
            __props__.__dict__["region_slug"] = None
            __props__.__dict__["resource_id"] = None
            __props__.__dict__["resource_type"] = None
            __props__.__dict__["started_at"] = None
            __props__.__dict__["status"] = None
        super(ImageActionsTransfer, __self__).__init__(
            'digitalocean-native:images/v2:ImageActionsTransfer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ImageActionsTransfer':
        """
        Get an existing ImageActionsTransfer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ImageActionsTransferArgs.__new__(ImageActionsTransferArgs)

        __props__.__dict__["completed_at"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["region_slug"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["resource_type"] = None
        __props__.__dict__["started_at"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["type"] = None
        return ImageActionsTransfer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="completedAt")
    def completed_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the action was completed.
        """
        return pulumi.get(self, "completed_at")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional['ImageActionsTransferPropertiesRegion']]:
        """
        The slug identifier for the region where the resource will initially be  available.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionSlug")
    def region_slug(self) -> pulumi.Output[Optional['outputs.RegionSlug']]:
        return pulumi.get(self, "region_slug")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[Optional[int]]:
        """
        A unique identifier for the resource that the action is associated with.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of resource that the action is associated with.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="startedAt")
    def started_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the action was initiated.
        """
        return pulumi.get(self, "started_at")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['Status']]:
        """
        The current status of the action. This can be "in-progress", "completed", or "errored".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['ImageActionBaseType']]:
        """
        The action to be taken on the image. Can be either `convert` or `transfer`.
        """
        return pulumi.get(self, "type")

