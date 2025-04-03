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
    'ListDatabasesBackupsProperties',
    'AwaitableListDatabasesBackupsProperties',
    'list_databases_backups',
    'list_databases_backups_output',
]

@pulumi.output_type
class ListDatabasesBackupsProperties:
    def __init__(__self__, backups=None):
        if backups and not isinstance(backups, list):
            raise TypeError("Expected argument 'backups' to be a list")
        pulumi.set(__self__, "backups", backups)

    @property
    @pulumi.getter
    def backups(self) -> Sequence['outputs.Backup']:
        return pulumi.get(self, "backups")


class AwaitableListDatabasesBackupsProperties(ListDatabasesBackupsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDatabasesBackupsProperties(
            backups=self.backups)


def list_databases_backups(database_cluster_uuid: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDatabasesBackupsProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:listDatabasesBackups', __args__, opts=opts, typ=ListDatabasesBackupsProperties).value

    return AwaitableListDatabasesBackupsProperties(
        backups=pulumi.get(__ret__, 'backups'))
def list_databases_backups_output(database_cluster_uuid: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListDatabasesBackupsProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:listDatabasesBackups', __args__, opts=opts, typ=ListDatabasesBackupsProperties)
    return __ret__.apply(lambda __response__: ListDatabasesBackupsProperties(
        backups=pulumi.get(__response__, 'backups')))
