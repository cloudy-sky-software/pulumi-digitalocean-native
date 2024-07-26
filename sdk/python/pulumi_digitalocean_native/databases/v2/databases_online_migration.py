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
from ._inputs import *

__all__ = ['DatabasesOnlineMigrationArgs', 'DatabasesOnlineMigration']

@pulumi.input_type
class DatabasesOnlineMigrationArgs:
    def __init__(__self__, *,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input['SourcePropertiesArgs']] = None):
        """
        The set of arguments for constructing a DatabasesOnlineMigration resource.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[bool] disable_ssl: Enables SSL encryption when connecting to the source database.
        """
        if database_cluster_uuid is not None:
            pulumi.set(__self__, "database_cluster_uuid", database_cluster_uuid)
        if disable_ssl is not None:
            pulumi.set(__self__, "disable_ssl", disable_ssl)
        if source is not None:
            pulumi.set(__self__, "source", source)

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
    @pulumi.getter(name="disableSsl")
    def disable_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables SSL encryption when connecting to the source database.
        """
        return pulumi.get(self, "disable_ssl")

    @disable_ssl.setter
    def disable_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_ssl", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input['SourcePropertiesArgs']]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input['SourcePropertiesArgs']]):
        pulumi.set(self, "source", value)


class DatabasesOnlineMigration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input[Union['SourcePropertiesArgs', 'SourcePropertiesArgsDict']]] = None,
                 __props__=None):
        """
        Create a DatabasesOnlineMigration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[bool] disable_ssl: Enables SSL encryption when connecting to the source database.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatabasesOnlineMigrationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DatabasesOnlineMigration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatabasesOnlineMigrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasesOnlineMigrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input[Union['SourcePropertiesArgs', 'SourcePropertiesArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasesOnlineMigrationArgs.__new__(DatabasesOnlineMigrationArgs)

            __props__.__dict__["database_cluster_uuid"] = database_cluster_uuid
            __props__.__dict__["disable_ssl"] = disable_ssl
            __props__.__dict__["source"] = source
        super(DatabasesOnlineMigration, __self__).__init__(
            'digitalocean-native:databases/v2:DatabasesOnlineMigration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabasesOnlineMigration':
        """
        Get an existing DatabasesOnlineMigration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabasesOnlineMigrationArgs.__new__(DatabasesOnlineMigrationArgs)

        __props__.__dict__["disable_ssl"] = None
        __props__.__dict__["source"] = None
        return DatabasesOnlineMigration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableSsl")
    def disable_ssl(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables SSL encryption when connecting to the source database.
        """
        return pulumi.get(self, "disable_ssl")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional['outputs.SourceProperties']]:
        return pulumi.get(self, "source")

