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
    'GetDatabasesUserProperties',
    'AwaitableGetDatabasesUserProperties',
    'get_databases_user',
    'get_databases_user_output',
]

@pulumi.output_type
class GetDatabasesUserProperties:
    def __init__(__self__, user=None):
        if user and not isinstance(user, dict):
            raise TypeError("Expected argument 'user' to be a dict")
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def user(self) -> 'outputs.DatabaseUser':
        return pulumi.get(self, "user")


class AwaitableGetDatabasesUserProperties(GetDatabasesUserProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesUserProperties(
            user=self.user)


def get_databases_user(database_cluster_uuid: Optional[str] = None,
                       username: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesUserProperties:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    :param str username: The name of the database user.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesUser', __args__, opts=opts, typ=GetDatabasesUserProperties).value

    return AwaitableGetDatabasesUserProperties(
        user=pulumi.get(__ret__, 'user'))
def get_databases_user_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                              username: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesUserProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    :param str username: The name of the database user.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:getDatabasesUser', __args__, opts=opts, typ=GetDatabasesUserProperties)
    return __ret__.apply(lambda __response__: GetDatabasesUserProperties(
        user=pulumi.get(__response__, 'user')))
