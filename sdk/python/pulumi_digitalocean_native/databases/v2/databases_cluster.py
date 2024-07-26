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
from ._inputs import *

__all__ = ['DatabasesClusterArgs', 'DatabasesCluster']

@pulumi.input_type
class DatabasesClusterArgs:
    def __init__(__self__, *,
                 engine: pulumi.Input['DatabaseClusterEngine'],
                 name: pulumi.Input[str],
                 num_nodes: pulumi.Input[int],
                 region: pulumi.Input[str],
                 size: pulumi.Input[str],
                 backup_restore: Optional[pulumi.Input['DatabaseBackupArgs']] = None,
                 connection: Optional[pulumi.Input['DatabaseClusterConnectionArgs']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 db_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 maintenance_window: Optional[pulumi.Input['DatabaseClusterMaintenanceWindowArgs']] = None,
                 private_connection: Optional[pulumi.Input['DatabaseClusterPrivateConnectionArgs']] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRuleArgs']]]] = None,
                 semantic_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseClusterStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseUserArgs']]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 version_end_of_availability: Optional[pulumi.Input[str]] = None,
                 version_end_of_life: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabasesCluster resource.
        :param pulumi.Input['DatabaseClusterEngine'] engine: A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        :param pulumi.Input[str] name: A unique, human-readable name referring to a database cluster.
        :param pulumi.Input[int] num_nodes: The number of nodes in the database cluster.
        :param pulumi.Input[str] region: The slug identifier for the region where the database cluster is located.
        :param pulumi.Input[str] size: The slug identifier representing the size of the nodes in the database cluster.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] db_names: An array of strings containing the names of databases created in the database cluster.
        :param pulumi.Input[str] private_network_uuid: A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        :param pulumi.Input[str] project_id: The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        :param pulumi.Input[str] semantic_version: A string representing the semantic version of the database engine in use for the cluster.
        :param pulumi.Input['DatabaseClusterStatus'] status: A string representing the current status of the database cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of tags that have been applied to the database cluster.
        :param pulumi.Input[str] version: A string representing the version of the database engine in use for the cluster.
        :param pulumi.Input[str] version_end_of_availability: A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        :param pulumi.Input[str] version_end_of_life: A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "num_nodes", num_nodes)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "size", size)
        if backup_restore is not None:
            pulumi.set(__self__, "backup_restore", backup_restore)
        if connection is not None:
            pulumi.set(__self__, "connection", connection)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if db_names is not None:
            pulumi.set(__self__, "db_names", db_names)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if private_connection is not None:
            pulumi.set(__self__, "private_connection", private_connection)
        if private_network_uuid is not None:
            pulumi.set(__self__, "private_network_uuid", private_network_uuid)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if semantic_version is not None:
            pulumi.set(__self__, "semantic_version", semantic_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if users is not None:
            pulumi.set(__self__, "users", users)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if version_end_of_availability is not None:
            pulumi.set(__self__, "version_end_of_availability", version_end_of_availability)
        if version_end_of_life is not None:
            pulumi.set(__self__, "version_end_of_life", version_end_of_life)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input['DatabaseClusterEngine']:
        """
        A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input['DatabaseClusterEngine']):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        A unique, human-readable name referring to a database cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="numNodes")
    def num_nodes(self) -> pulumi.Input[int]:
        """
        The number of nodes in the database cluster.
        """
        return pulumi.get(self, "num_nodes")

    @num_nodes.setter
    def num_nodes(self, value: pulumi.Input[int]):
        pulumi.set(self, "num_nodes", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The slug identifier for the region where the database cluster is located.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[str]:
        """
        The slug identifier representing the size of the nodes in the database cluster.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[str]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="backupRestore")
    def backup_restore(self) -> Optional[pulumi.Input['DatabaseBackupArgs']]:
        return pulumi.get(self, "backup_restore")

    @backup_restore.setter
    def backup_restore(self, value: Optional[pulumi.Input['DatabaseBackupArgs']]):
        pulumi.set(self, "backup_restore", value)

    @property
    @pulumi.getter
    def connection(self) -> Optional[pulumi.Input['DatabaseClusterConnectionArgs']]:
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: Optional[pulumi.Input['DatabaseClusterConnectionArgs']]):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dbNames")
    def db_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings containing the names of databases created in the database cluster.
        """
        return pulumi.get(self, "db_names")

    @db_names.setter
    def db_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "db_names", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['DatabaseClusterMaintenanceWindowArgs']]:
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['DatabaseClusterMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter(name="privateConnection")
    def private_connection(self) -> Optional[pulumi.Input['DatabaseClusterPrivateConnectionArgs']]:
        return pulumi.get(self, "private_connection")

    @private_connection.setter
    def private_connection(self, value: Optional[pulumi.Input['DatabaseClusterPrivateConnectionArgs']]):
        pulumi.set(self, "private_connection", value)

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        """
        return pulumi.get(self, "private_network_uuid")

    @private_network_uuid.setter
    def private_network_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_uuid", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="semanticVersion")
    def semantic_version(self) -> Optional[pulumi.Input[str]]:
        """
        A string representing the semantic version of the database engine in use for the cluster.
        """
        return pulumi.get(self, "semantic_version")

    @semantic_version.setter
    def semantic_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "semantic_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['DatabaseClusterStatus']]:
        """
        A string representing the current status of the database cluster.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['DatabaseClusterStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of tags that have been applied to the database cluster.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseUserArgs']]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatabaseUserArgs']]]]):
        pulumi.set(self, "users", value)

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

    @property
    @pulumi.getter(name="versionEndOfAvailability")
    def version_end_of_availability(self) -> Optional[pulumi.Input[str]]:
        """
        A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        """
        return pulumi.get(self, "version_end_of_availability")

    @version_end_of_availability.setter
    def version_end_of_availability(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_end_of_availability", value)

    @property
    @pulumi.getter(name="versionEndOfLife")
    def version_end_of_life(self) -> Optional[pulumi.Input[str]]:
        """
        A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        """
        return pulumi.get(self, "version_end_of_life")

    @version_end_of_life.setter
    def version_end_of_life(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_end_of_life", value)


class DatabasesCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_restore: Optional[pulumi.Input[Union['DatabaseBackupArgs', 'DatabaseBackupArgsDict']]] = None,
                 connection: Optional[pulumi.Input[Union['DatabaseClusterConnectionArgs', 'DatabaseClusterConnectionArgsDict']]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 db_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 engine: Optional[pulumi.Input['DatabaseClusterEngine']] = None,
                 maintenance_window: Optional[pulumi.Input[Union['DatabaseClusterMaintenanceWindowArgs', 'DatabaseClusterMaintenanceWindowArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 num_nodes: Optional[pulumi.Input[int]] = None,
                 private_connection: Optional[pulumi.Input[Union['DatabaseClusterPrivateConnectionArgs', 'DatabaseClusterPrivateConnectionArgsDict']]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRuleArgs', 'FirewallRuleArgsDict']]]]] = None,
                 semantic_version: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseClusterStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DatabaseUserArgs', 'DatabaseUserArgsDict']]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 version_end_of_availability: Optional[pulumi.Input[str]] = None,
                 version_end_of_life: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DatabasesCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] db_names: An array of strings containing the names of databases created in the database cluster.
        :param pulumi.Input['DatabaseClusterEngine'] engine: A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        :param pulumi.Input[str] name: A unique, human-readable name referring to a database cluster.
        :param pulumi.Input[int] num_nodes: The number of nodes in the database cluster.
        :param pulumi.Input[str] private_network_uuid: A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        :param pulumi.Input[str] project_id: The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        :param pulumi.Input[str] region: The slug identifier for the region where the database cluster is located.
        :param pulumi.Input[str] semantic_version: A string representing the semantic version of the database engine in use for the cluster.
        :param pulumi.Input[str] size: The slug identifier representing the size of the nodes in the database cluster.
        :param pulumi.Input['DatabaseClusterStatus'] status: A string representing the current status of the database cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of tags that have been applied to the database cluster.
        :param pulumi.Input[str] version: A string representing the version of the database engine in use for the cluster.
        :param pulumi.Input[str] version_end_of_availability: A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        :param pulumi.Input[str] version_end_of_life: A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasesClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DatabasesCluster resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatabasesClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasesClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_restore: Optional[pulumi.Input[Union['DatabaseBackupArgs', 'DatabaseBackupArgsDict']]] = None,
                 connection: Optional[pulumi.Input[Union['DatabaseClusterConnectionArgs', 'DatabaseClusterConnectionArgsDict']]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 db_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 engine: Optional[pulumi.Input['DatabaseClusterEngine']] = None,
                 maintenance_window: Optional[pulumi.Input[Union['DatabaseClusterMaintenanceWindowArgs', 'DatabaseClusterMaintenanceWindowArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 num_nodes: Optional[pulumi.Input[int]] = None,
                 private_connection: Optional[pulumi.Input[Union['DatabaseClusterPrivateConnectionArgs', 'DatabaseClusterPrivateConnectionArgsDict']]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FirewallRuleArgs', 'FirewallRuleArgsDict']]]]] = None,
                 semantic_version: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseClusterStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DatabaseUserArgs', 'DatabaseUserArgsDict']]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 version_end_of_availability: Optional[pulumi.Input[str]] = None,
                 version_end_of_life: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasesClusterArgs.__new__(DatabasesClusterArgs)

            __props__.__dict__["backup_restore"] = backup_restore
            __props__.__dict__["connection"] = connection
            __props__.__dict__["created_at"] = created_at
            __props__.__dict__["db_names"] = db_names
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["maintenance_window"] = maintenance_window
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if num_nodes is None and not opts.urn:
                raise TypeError("Missing required property 'num_nodes'")
            __props__.__dict__["num_nodes"] = num_nodes
            __props__.__dict__["private_connection"] = private_connection
            __props__.__dict__["private_network_uuid"] = private_network_uuid
            __props__.__dict__["project_id"] = project_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
            __props__.__dict__["semantic_version"] = semantic_version
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["users"] = users
            __props__.__dict__["version"] = version
            __props__.__dict__["version_end_of_availability"] = version_end_of_availability
            __props__.__dict__["version_end_of_life"] = version_end_of_life
            __props__.__dict__["database"] = None
        super(DatabasesCluster, __self__).__init__(
            'digitalocean-native:databases/v2:DatabasesCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabasesCluster':
        """
        Get an existing DatabasesCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabasesClusterArgs.__new__(DatabasesClusterArgs)

        __props__.__dict__["backup_restore"] = None
        __props__.__dict__["connection"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["database"] = None
        __props__.__dict__["db_names"] = None
        __props__.__dict__["engine"] = None
        __props__.__dict__["maintenance_window"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["num_nodes"] = None
        __props__.__dict__["private_connection"] = None
        __props__.__dict__["private_network_uuid"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["semantic_version"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["users"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["version_end_of_availability"] = None
        __props__.__dict__["version_end_of_life"] = None
        return DatabasesCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupRestore")
    def backup_restore(self) -> pulumi.Output[Optional['outputs.DatabaseBackup']]:
        return pulumi.get(self, "backup_restore")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output[Optional['outputs.DatabaseClusterConnection']]:
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        """
        A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output['outputs.DatabaseCluster']:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="dbNames")
    def db_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of strings containing the names of databases created in the database cluster.
        """
        return pulumi.get(self, "db_names")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional['DatabaseClusterEngine']]:
        """
        A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output[Optional['outputs.DatabaseClusterMaintenanceWindow']]:
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        A unique, human-readable name referring to a database cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numNodes")
    def num_nodes(self) -> pulumi.Output[Optional[int]]:
        """
        The number of nodes in the database cluster.
        """
        return pulumi.get(self, "num_nodes")

    @property
    @pulumi.getter(name="privateConnection")
    def private_connection(self) -> pulumi.Output[Optional['outputs.DatabaseClusterPrivateConnection']]:
        return pulumi.get(self, "private_connection")

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> pulumi.Output[Optional[str]]:
        """
        A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.
        """
        return pulumi.get(self, "private_network_uuid")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The slug identifier for the region where the database cluster is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallRule']]]:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="semanticVersion")
    def semantic_version(self) -> pulumi.Output[Optional[str]]:
        """
        A string representing the semantic version of the database engine in use for the cluster.
        """
        return pulumi.get(self, "semantic_version")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[str]]:
        """
        The slug identifier representing the size of the nodes in the database cluster.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['DatabaseClusterStatus']]:
        """
        A string representing the current status of the database cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of tags that have been applied to the database cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Optional[Sequence['outputs.DatabaseUser']]]:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        A string representing the version of the database engine in use for the cluster.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionEndOfAvailability")
    def version_end_of_availability(self) -> pulumi.Output[Optional[str]]:
        """
        A timestamp referring to the date when the particular version will no longer be available for creating new clusters. If null, the version does not have an end of availability timeline.
        """
        return pulumi.get(self, "version_end_of_availability")

    @property
    @pulumi.getter(name="versionEndOfLife")
    def version_end_of_life(self) -> pulumi.Output[Optional[str]]:
        """
        A timestamp referring to the date when the particular version will no longer be supported. If null, the version does not have an end of life timeline.
        """
        return pulumi.get(self, "version_end_of_life")

