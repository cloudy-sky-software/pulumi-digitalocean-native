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
    'ListDropletsBackupsResult',
    'AwaitableListDropletsBackupsResult',
    'list_droplets_backups',
    'list_droplets_backups_output',
]

@pulumi.output_type
class ListDropletsBackupsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDropletsBackupsItems':
        return pulumi.get(self, "items")


class AwaitableListDropletsBackupsResult(ListDropletsBackupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDropletsBackupsResult(
            items=self.items)


def list_droplets_backups(droplet_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDropletsBackupsResult:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:listDropletsBackups', __args__, opts=opts, typ=ListDropletsBackupsResult).value

    return AwaitableListDropletsBackupsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_droplets_backups)
def list_droplets_backups_output(droplet_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDropletsBackupsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    ...
