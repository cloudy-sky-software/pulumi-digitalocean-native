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
    'GetDatabasesCaResult',
    'AwaitableGetDatabasesCaResult',
    'get_databases_ca',
    'get_databases_ca_output',
]

@pulumi.output_type
class GetDatabasesCaResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetDatabasesCaProperties':
        return pulumi.get(self, "items")


class AwaitableGetDatabasesCaResult(GetDatabasesCaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesCaResult(
            items=self.items)


def get_databases_ca(database_cluster_uuid: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesCaResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesCa', __args__, opts=opts, typ=GetDatabasesCaResult).value

    return AwaitableGetDatabasesCaResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_databases_ca)
def get_databases_ca_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesCaResult]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    ...
