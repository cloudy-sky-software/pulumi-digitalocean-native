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
    'ListRegistryRepositoriesItems',
    'AwaitableListRegistryRepositoriesItems',
    'list_registry_repositories',
    'list_registry_repositories_output',
]

@pulumi.output_type
class ListRegistryRepositoriesItems:
    def __init__(__self__, links=None, meta=None, repositories=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def repositories(self) -> Optional[Sequence['outputs.Repository']]:
        return pulumi.get(self, "repositories")


class AwaitableListRegistryRepositoriesItems(ListRegistryRepositoriesItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListRegistryRepositoriesItems(
            links=self.links,
            meta=self.meta,
            repositories=self.repositories)


def list_registry_repositories(registry_name: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListRegistryRepositoriesItems:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str registry_name: The name of a container registry.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:registry/v2:listRegistryRepositories', __args__, opts=opts, typ=ListRegistryRepositoriesItems).value

    return AwaitableListRegistryRepositoriesItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        repositories=pulumi.get(__ret__, 'repositories'))
def list_registry_repositories_output(registry_name: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListRegistryRepositoriesItems]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str registry_name: The name of a container registry.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:registry/v2:listRegistryRepositories', __args__, opts=opts, typ=ListRegistryRepositoriesItems)
    return __ret__.apply(lambda __response__: ListRegistryRepositoriesItems(
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta'),
        repositories=pulumi.get(__response__, 'repositories')))
