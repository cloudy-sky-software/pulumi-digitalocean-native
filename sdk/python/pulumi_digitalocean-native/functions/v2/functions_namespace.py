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

__all__ = ['FunctionsNamespaceArgs', 'FunctionsNamespace']

@pulumi.input_type
class FunctionsNamespaceArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[str],
                 region: pulumi.Input[str]):
        """
        The set of arguments for constructing a FunctionsNamespace resource.
        :param pulumi.Input[str] label: The namespace's unique name.
        :param pulumi.Input[str] region: The [datacenter region](https://docs.digitalocean.com/products/platform/availability-matrix/#available-datacenters) in which to create the namespace.
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[str]:
        """
        The namespace's unique name.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The [datacenter region](https://docs.digitalocean.com/products/platform/availability-matrix/#available-datacenters) in which to create the namespace.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


class FunctionsNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FunctionsNamespace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The namespace's unique name.
        :param pulumi.Input[str] region: The [datacenter region](https://docs.digitalocean.com/products/platform/availability-matrix/#available-datacenters) in which to create the namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionsNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FunctionsNamespace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FunctionsNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionsNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionsNamespaceArgs.__new__(FunctionsNamespaceArgs)

            if label is None and not opts.urn:
                raise TypeError("Missing required property 'label'")
            __props__.__dict__["label"] = label
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["namespace"] = None
        super(FunctionsNamespace, __self__).__init__(
            'digitalocean-native:functions/v2:FunctionsNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FunctionsNamespace':
        """
        Get an existing FunctionsNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FunctionsNamespaceArgs.__new__(FunctionsNamespaceArgs)

        __props__.__dict__["label"] = None
        __props__.__dict__["namespace"] = None
        __props__.__dict__["region"] = None
        return FunctionsNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[str]:
        """
        The namespace's unique name.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional['outputs.NamespaceInfo']]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The [datacenter region](https://docs.digitalocean.com/products/platform/availability-matrix/#available-datacenters) in which to create the namespace.
        """
        return pulumi.get(self, "region")

