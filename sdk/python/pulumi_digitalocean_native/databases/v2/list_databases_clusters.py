# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'ListDatabasesClustersProperties',
    'AwaitableListDatabasesClustersProperties',
    'list_databases_clusters',
    'list_databases_clusters_output',
]

@pulumi.output_type
class ListDatabasesClustersProperties:
    def __init__(__self__, databases=None):
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)

    @property
    @pulumi.getter
    def databases(self) -> Optional[Sequence['outputs.DatabaseCluster']]:
        return pulumi.get(self, "databases")


class AwaitableListDatabasesClustersProperties(ListDatabasesClustersProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDatabasesClustersProperties(
            databases=self.databases)


def list_databases_clusters(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDatabasesClustersProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:listDatabasesClusters', __args__, opts=opts, typ=ListDatabasesClustersProperties).value

    return AwaitableListDatabasesClustersProperties(
        databases=pulumi.get(__ret__, 'databases'))
def list_databases_clusters_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListDatabasesClustersProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:listDatabasesClusters', __args__, opts=opts, typ=ListDatabasesClustersProperties)
    return __ret__.apply(lambda __response__: ListDatabasesClustersProperties(
        databases=pulumi.get(__response__, 'databases')))
