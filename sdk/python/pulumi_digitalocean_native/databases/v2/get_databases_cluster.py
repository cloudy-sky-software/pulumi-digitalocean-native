# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetDatabasesClusterProperties',
    'AwaitableGetDatabasesClusterProperties',
    'get_databases_cluster',
    'get_databases_cluster_output',
]

@pulumi.output_type
class GetDatabasesClusterProperties:
    def __init__(__self__, database=None):
        if database and not isinstance(database, dict):
            raise TypeError("Expected argument 'database' to be a dict")
        pulumi.set(__self__, "database", database)

    @property
    @pulumi.getter
    def database(self) -> 'outputs.DatabaseCluster':
        return pulumi.get(self, "database")


class AwaitableGetDatabasesClusterProperties(GetDatabasesClusterProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesClusterProperties(
            database=self.database)


def get_databases_cluster(database_cluster_uuid: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesClusterProperties:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesCluster', __args__, opts=opts, typ=GetDatabasesClusterProperties).value

    return AwaitableGetDatabasesClusterProperties(
        database=pulumi.get(__ret__, 'database'))
def get_databases_cluster_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabasesClusterProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:getDatabasesCluster', __args__, opts=opts, typ=GetDatabasesClusterProperties)
    return __ret__.apply(lambda __response__: GetDatabasesClusterProperties(
        database=pulumi.get(__response__, 'database')))
