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

__all__ = [
    'GetDatabasesConnectionPoolResult',
    'AwaitableGetDatabasesConnectionPoolResult',
    'get_databases_connection_pool',
    'get_databases_connection_pool_output',
]

@pulumi.output_type
class GetDatabasesConnectionPoolResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetDatabasesConnectionPoolProperties':
        return pulumi.get(self, "items")


class AwaitableGetDatabasesConnectionPoolResult(GetDatabasesConnectionPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesConnectionPoolResult(
            items=self.items)


def get_databases_connection_pool(database_cluster_uuid: Optional[str] = None,
                                  pool_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesConnectionPoolResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    :param str pool_name: The name used to identify the connection pool.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    __args__['poolName'] = pool_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesConnectionPool', __args__, opts=opts, typ=GetDatabasesConnectionPoolResult).value

    return AwaitableGetDatabasesConnectionPoolResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_databases_connection_pool)
def get_databases_connection_pool_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                         pool_name: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesConnectionPoolResult]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    :param str pool_name: The name used to identify the connection pool.
    """
    ...