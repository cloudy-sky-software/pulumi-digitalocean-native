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

__all__ = [
    'GetDatabasesCaProperties',
    'AwaitableGetDatabasesCaProperties',
    'get_databases_ca',
    'get_databases_ca_output',
]

@pulumi.output_type
class GetDatabasesCaProperties:
    def __init__(__self__, ca=None):
        if ca and not isinstance(ca, dict):
            raise TypeError("Expected argument 'ca' to be a dict")
        pulumi.set(__self__, "ca", ca)

    @property
    @pulumi.getter
    def ca(self) -> 'outputs.Ca':
        return pulumi.get(self, "ca")


class AwaitableGetDatabasesCaProperties(GetDatabasesCaProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesCaProperties(
            ca=self.ca)


def get_databases_ca(database_cluster_uuid: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesCaProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesCa', __args__, opts=opts, typ=GetDatabasesCaProperties).value

    return AwaitableGetDatabasesCaProperties(
        ca=pulumi.get(__ret__, 'ca'))
def get_databases_ca_output(database_cluster_uuid: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabasesCaProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:getDatabasesCa', __args__, opts=opts, typ=GetDatabasesCaProperties)
    return __ret__.apply(lambda __response__: GetDatabasesCaProperties(
        ca=pulumi.get(__response__, 'ca')))
