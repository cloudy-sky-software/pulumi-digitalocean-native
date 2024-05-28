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

__all__ = ['DatabasesReplicaArgs', 'DatabasesReplica']

@pulumi.input_type
class DatabasesReplicaArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 connection: Optional[pulumi.Input['DatabaseReplicaConnectionArgs']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 private_connection: Optional[pulumi.Input['DatabaseReplicaPrivateConnectionArgs']] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseReplicaStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DatabasesReplica resource.
        :param pulumi.Input[str] name: The name to give the read-only replicating
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[str] private_network_uuid: A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.
        :param pulumi.Input[str] region: A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.
        :param pulumi.Input[str] size: A slug identifier representing the size of the node for the read-only replica. The size of the replica must be at least as large as the node size for the database cluster from which it is replicating.
        :param pulumi.Input['DatabaseReplicaStatus'] status: A string representing the current status of the database cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.
        """
        pulumi.set(__self__, "name", name)
        if connection is not None:
            pulumi.set(__self__, "connection", connection)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if database_cluster_uuid is not None:
            pulumi.set(__self__, "database_cluster_uuid", database_cluster_uuid)
        if private_connection is not None:
            pulumi.set(__self__, "private_connection", private_connection)
        if private_network_uuid is not None:
            pulumi.set(__self__, "private_network_uuid", private_network_uuid)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name to give the read-only replicating
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def connection(self) -> Optional[pulumi.Input['DatabaseReplicaConnectionArgs']]:
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: Optional[pulumi.Input['DatabaseReplicaConnectionArgs']]):
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
    @pulumi.getter(name="privateConnection")
    def private_connection(self) -> Optional[pulumi.Input['DatabaseReplicaPrivateConnectionArgs']]:
        return pulumi.get(self, "private_connection")

    @private_connection.setter
    def private_connection(self, value: Optional[pulumi.Input['DatabaseReplicaPrivateConnectionArgs']]):
        pulumi.set(self, "private_connection", value)

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.
        """
        return pulumi.get(self, "private_network_uuid")

    @private_network_uuid.setter
    def private_network_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_uuid", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        A slug identifier representing the size of the node for the read-only replica. The size of the replica must be at least as large as the node size for the database cluster from which it is replicating.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['DatabaseReplicaStatus']]:
        """
        A string representing the current status of the database cluster.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['DatabaseReplicaStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class DatabasesReplica(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['DatabaseReplicaConnectionArgs']]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_connection: Optional[pulumi.Input[pulumi.InputType['DatabaseReplicaPrivateConnectionArgs']]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseReplicaStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a DatabasesReplica resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format that represents when the database cluster was created.
        :param pulumi.Input[str] database_cluster_uuid: A unique identifier for a database cluster.
        :param pulumi.Input[str] name: The name to give the read-only replicating
        :param pulumi.Input[str] private_network_uuid: A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.
        :param pulumi.Input[str] region: A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.
        :param pulumi.Input[str] size: A slug identifier representing the size of the node for the read-only replica. The size of the replica must be at least as large as the node size for the database cluster from which it is replicating.
        :param pulumi.Input['DatabaseReplicaStatus'] status: A string representing the current status of the database cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasesReplicaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DatabasesReplica resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatabasesReplicaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasesReplicaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['DatabaseReplicaConnectionArgs']]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_connection: Optional[pulumi.Input[pulumi.InputType['DatabaseReplicaPrivateConnectionArgs']]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['DatabaseReplicaStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasesReplicaArgs.__new__(DatabasesReplicaArgs)

            __props__.__dict__["connection"] = connection
            __props__.__dict__["created_at"] = created_at
            __props__.__dict__["database_cluster_uuid"] = database_cluster_uuid
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["private_connection"] = private_connection
            __props__.__dict__["private_network_uuid"] = private_network_uuid
            __props__.__dict__["region"] = region
            __props__.__dict__["size"] = size
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["replica"] = None
        super(DatabasesReplica, __self__).__init__(
            'digitalocean-native:databases/v2:DatabasesReplica',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabasesReplica':
        """
        Get an existing DatabasesReplica resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabasesReplicaArgs.__new__(DatabasesReplicaArgs)

        __props__.__dict__["connection"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_connection"] = None
        __props__.__dict__["private_network_uuid"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["replica"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        return DatabasesReplica(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output[Optional['outputs.DatabaseReplicaConnection']]:
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
    def name(self) -> pulumi.Output[str]:
        """
        The name to give the read-only replicating
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateConnection")
    def private_connection(self) -> pulumi.Output[Optional['outputs.DatabaseReplicaPrivateConnection']]:
        return pulumi.get(self, "private_connection")

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> pulumi.Output[Optional[str]]:
        """
        A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.
        """
        return pulumi.get(self, "private_network_uuid")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def replica(self) -> pulumi.Output[Optional['outputs.DatabaseReplica']]:
        return pulumi.get(self, "replica")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        """
        A slug identifier representing the size of the node for the read-only replica. The size of the replica must be at least as large as the node size for the database cluster from which it is replicating.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['DatabaseReplicaStatus']]:
        """
        A string representing the current status of the database cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.
        """
        return pulumi.get(self, "tags")
