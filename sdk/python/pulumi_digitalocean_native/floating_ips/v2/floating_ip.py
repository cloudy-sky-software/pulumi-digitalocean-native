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

__all__ = ['FloatingIPArgs', 'FloatingIP']

@pulumi.input_type
class FloatingIPArgs:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a FloatingIP resource.
        :param pulumi.Input[builtins.str] project_id: The UUID of the project to which the floating IP will be assigned.
        :param pulumi.Input[builtins.str] region: The slug identifier for the region the floating IP will be reserved to.
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UUID of the project to which the floating IP will be assigned.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The slug identifier for the region the floating IP will be reserved to.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


class FloatingIP(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a FloatingIP resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] project_id: The UUID of the project to which the floating IP will be assigned.
        :param pulumi.Input[builtins.str] region: The slug identifier for the region the floating IP will be reserved to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FloatingIPArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FloatingIP resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FloatingIPArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FloatingIPArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FloatingIPArgs.__new__(FloatingIPArgs)

            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["floating_ip"] = None
            __props__.__dict__["links"] = None
        super(FloatingIP, __self__).__init__(
            'digitalocean-native:floating_ips/v2:FloatingIP',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FloatingIP':
        """
        Get an existing FloatingIP resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FloatingIPArgs.__new__(FloatingIPArgs)

        __props__.__dict__["floating_ip"] = None
        __props__.__dict__["links"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["region"] = None
        return FloatingIP(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Output[Optional['outputs.FloatingIp']]:
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Optional['outputs.LinksProperties']]:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The UUID of the project to which the floating IP will be assigned.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The slug identifier for the region the floating IP will be reserved to.
        """
        return pulumi.get(self, "region")

