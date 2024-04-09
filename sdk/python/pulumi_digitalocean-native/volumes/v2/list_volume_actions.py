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
    'ListVolumeActionsResult',
    'AwaitableListVolumeActionsResult',
    'list_volume_actions',
    'list_volume_actions_output',
]

@pulumi.output_type
class ListVolumeActionsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListVolumeActionsItems':
        return pulumi.get(self, "items")


class AwaitableListVolumeActionsResult(ListVolumeActionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListVolumeActionsResult(
            items=self.items)


def list_volume_actions(volume_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListVolumeActionsResult:
    """
    Use this data source to access information about an existing resource.

    :param str volume_id: The ID of the block storage volume.
    """
    __args__ = dict()
    __args__['volumeId'] = volume_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:volumes/v2:listVolumeActions', __args__, opts=opts, typ=ListVolumeActionsResult).value

    return AwaitableListVolumeActionsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_volume_actions)
def list_volume_actions_output(volume_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListVolumeActionsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str volume_id: The ID of the block storage volume.
    """
    ...
