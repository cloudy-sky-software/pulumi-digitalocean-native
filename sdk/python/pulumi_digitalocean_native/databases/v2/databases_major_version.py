# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['DatabasesMajorVersionArgs', 'DatabasesMajorVersion']

@pulumi.input_type
class DatabasesMajorVersionArgs:
    def __init__(__self__, *,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabasesMajorVersion resource.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[str] version: A string representing the version of the database engine in use for the cluster.
        """
        if database_cluster_uuid is not None:
            pulumi.set(__self__, "database_cluster_uuid", database_cluster_uuid)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="databaseClusterUuid")
    def database_cluster_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for a database cluster.
        """
        return pulumi.get(self, "database_cluster_uuid")

    @database_cluster_uuid.setter
    def database_cluster_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_cluster_uuid", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        A string representing the version of the database engine in use for the cluster.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class DatabasesMajorVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DatabasesMajorVersion resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[str] version: A string representing the version of the database engine in use for the cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatabasesMajorVersionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DatabasesMajorVersion resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatabasesMajorVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasesMajorVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasesMajorVersionArgs.__new__(DatabasesMajorVersionArgs)

            __props__.__dict__["database_cluster_uuid"] = database_cluster_uuid
            __props__.__dict__["version"] = version
        super(DatabasesMajorVersion, __self__).__init__(
            'digitalocean-native:databases/v2:DatabasesMajorVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabasesMajorVersion':
        """
        Get an existing DatabasesMajorVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabasesMajorVersionArgs.__new__(DatabasesMajorVersionArgs)

        __props__.__dict__["version"] = None
        return DatabasesMajorVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        A string representing the version of the database engine in use for the cluster.
        """
        return pulumi.get(self, "version")

