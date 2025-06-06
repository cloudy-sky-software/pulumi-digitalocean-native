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
    'ListSshKeysItems',
    'AwaitableListSshKeysItems',
    'list_ssh_keys',
    'list_ssh_keys_output',
]

@pulumi.output_type
class ListSshKeysItems:
    def __init__(__self__, links=None, meta=None, ssh_keys=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if ssh_keys and not isinstance(ssh_keys, list):
            raise TypeError("Expected argument 'ssh_keys' to be a list")
        pulumi.set(__self__, "ssh_keys", ssh_keys)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> Optional[Sequence['outputs.SshKeys']]:
        return pulumi.get(self, "ssh_keys")


class AwaitableListSshKeysItems(ListSshKeysItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSshKeysItems(
            links=self.links,
            meta=self.meta,
            ssh_keys=self.ssh_keys)


def list_ssh_keys(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSshKeysItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:account/v2:listSshKeys', __args__, opts=opts, typ=ListSshKeysItems).value

    return AwaitableListSshKeysItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        ssh_keys=pulumi.get(__ret__, 'ssh_keys'))
def list_ssh_keys_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListSshKeysItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:account/v2:listSshKeys', __args__, opts=opts, typ=ListSshKeysItems)
    return __ret__.apply(lambda __response__: ListSshKeysItems(
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta'),
        ssh_keys=pulumi.get(__response__, 'ssh_keys')))
