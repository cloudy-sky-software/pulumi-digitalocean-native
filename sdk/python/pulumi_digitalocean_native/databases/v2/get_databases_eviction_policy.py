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

__all__ = [
    'GetDatabasesEvictionPolicyResult',
    'AwaitableGetDatabasesEvictionPolicyResult',
    'get_databases_eviction_policy',
    'get_databases_eviction_policy_output',
]

@pulumi.output_type
class GetDatabasesEvictionPolicyResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetDatabasesEvictionPolicyProperties':
        return pulumi.get(self, "items")


class AwaitableGetDatabasesEvictionPolicyResult(GetDatabasesEvictionPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesEvictionPolicyResult(
            items=self.items)


def get_databases_eviction_policy(database_cluster_uuid: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesEvictionPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesEvictionPolicy', __args__, opts=opts, typ=GetDatabasesEvictionPolicyResult).value

    return AwaitableGetDatabasesEvictionPolicyResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_databases_eviction_policy)
def get_databases_eviction_policy_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesEvictionPolicyResult]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    ...
