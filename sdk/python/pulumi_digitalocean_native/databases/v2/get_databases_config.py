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
    'DatabaseConfig',
    'AwaitableDatabaseConfig',
    'get_databases_config',
    'get_databases_config_output',
]

@pulumi.output_type
class DatabaseConfig:
    def __init__(__self__, config=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter
    def config(self) -> Optional[Any]:
        return pulumi.get(self, "config")


class AwaitableDatabaseConfig(DatabaseConfig):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DatabaseConfig(
            config=self.config)


def get_databases_config(database_cluster_uuid: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableDatabaseConfig:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:getDatabasesConfig', __args__, opts=opts, typ=DatabaseConfig).value

    return AwaitableDatabaseConfig(
        config=pulumi.get(__ret__, 'config'))
def get_databases_config_output(database_cluster_uuid: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[DatabaseConfig]:
    """
    Use this data source to access information about an existing resource.

    :param str database_cluster_uuid: A unique identifier for a database cluster.
    """
    __args__ = dict()
    __args__['databaseClusterUuid'] = database_cluster_uuid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:getDatabasesConfig', __args__, opts=opts, typ=DatabaseConfig)
    return __ret__.apply(lambda __response__: DatabaseConfig(
        config=pulumi.get(__response__, 'config')))
