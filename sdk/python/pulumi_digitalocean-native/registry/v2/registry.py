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

__all__ = ['RegistryArgs', 'Registry']

@pulumi.input_type
class RegistryArgs:
    def __init__(__self__, *,
                 subscription_tier_slug: pulumi.Input['SubscriptionTierSlug'],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['Region']] = None):
        """
        The set of arguments for constructing a Registry resource.
        :param pulumi.Input['SubscriptionTierSlug'] subscription_tier_slug: The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        :param pulumi.Input[str] name: A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        :param pulumi.Input['Region'] region: Slug of the region where registry data is stored. When not provided, a region will be selected.
        """
        pulumi.set(__self__, "subscription_tier_slug", subscription_tier_slug)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="subscriptionTierSlug")
    def subscription_tier_slug(self) -> pulumi.Input['SubscriptionTierSlug']:
        """
        The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        """
        return pulumi.get(self, "subscription_tier_slug")

    @subscription_tier_slug.setter
    def subscription_tier_slug(self, value: pulumi.Input['SubscriptionTierSlug']):
        pulumi.set(self, "subscription_tier_slug", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input['Region']]:
        """
        Slug of the region where registry data is stored. When not provided, a region will be selected.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input['Region']]):
        pulumi.set(self, "region", value)


class Registry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['Region']] = None,
                 subscription_tier_slug: Optional[pulumi.Input['SubscriptionTierSlug']] = None,
                 __props__=None):
        """
        Create a Registry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        :param pulumi.Input['Region'] region: Slug of the region where registry data is stored. When not provided, a region will be selected.
        :param pulumi.Input['SubscriptionTierSlug'] subscription_tier_slug: The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Registry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input['Region']] = None,
                 subscription_tier_slug: Optional[pulumi.Input['SubscriptionTierSlug']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryArgs.__new__(RegistryArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if subscription_tier_slug is None and not opts.urn:
                raise TypeError("Missing required property 'subscription_tier_slug'")
            __props__.__dict__["subscription_tier_slug"] = subscription_tier_slug
            __props__.__dict__["registry"] = None
        super(Registry, __self__).__init__(
            'digitalocean-native:registry/v2:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Registry':
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegistryArgs.__new__(RegistryArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["registry"] = None
        __props__.__dict__["subscription_tier_slug"] = None
        return Registry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional['Region']]:
        """
        Slug of the region where registry data is stored. When not provided, a region will be selected.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Output[Optional['outputs.Registry']]:
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter(name="subscriptionTierSlug")
    def subscription_tier_slug(self) -> pulumi.Output['SubscriptionTierSlug']:
        """
        The slug of the subscription tier to sign up for. Valid values can be retrieved using the options endpoint.
        """
        return pulumi.get(self, "subscription_tier_slug")

