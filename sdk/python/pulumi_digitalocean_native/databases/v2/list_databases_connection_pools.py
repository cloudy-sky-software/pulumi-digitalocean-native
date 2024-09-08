# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from ... import _utilities
from . import outputs

__all__ = [
    'ConnectionPools',
    'AwaitableConnectionPools',
    'list_databases_connection_pools',
    'list_databases_connection_pools_output',
]

@pulumi.output_type
class ConnectionPools:
    def __init__(__self__, pools=None):
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)

    @property
    @pulumi.getter
    def pools(self) -> Optional[Sequence['outputs.ConnectionPool']]:
        """
        An array of connection pool objects.
        """
        return pulumi.get(self, "pools")


class AwaitableConnectionPools(ConnectionPools):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ConnectionPools(
            pools=self.pools)


def list_databases_connection_pools(database_cluster_uuid: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableConnectionPools:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:listDatabasesConnectionPools', __args__, opts=opts, typ=ConnectionPools).value

    return AwaitableConnectionPools(
        pools=pulumi.get(__ret__, 'pools'))


@_utilities.lift_output_func(list_databases_connection_pools)
def list_databases_connection_pools_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ConnectionPools]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    ...
