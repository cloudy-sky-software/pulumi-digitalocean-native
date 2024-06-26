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
    'GetDatabasesSqlModeResult',
    'AwaitableGetDatabasesSqlModeResult',
    'get_databases_sql_mode',
    'get_databases_sql_mode_output',
]

@pulumi.output_type
class GetDatabasesSqlModeResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.SqlMode':
        return pulumi.get(self, "items")


class AwaitableGetDatabasesSqlModeResult(GetDatabasesSqlModeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesSqlModeResult(
            items=self.items)


def get_databases_sql_mode(database_cluster_uuid: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesSqlModeResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesSqlMode', __args__, opts=opts, typ=GetDatabasesSqlModeResult).value

    return AwaitableGetDatabasesSqlModeResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_databases_sql_mode)
def get_databases_sql_mode_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesSqlModeResult]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    ...
