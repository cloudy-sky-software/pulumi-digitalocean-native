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
    'ListImageActionsResult',
    'AwaitableListImageActionsResult',
    'list_image_actions',
    'list_image_actions_output',
]

@pulumi.output_type
class ListImageActionsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListImageActionsItems':
        return pulumi.get(self, "items")


class AwaitableListImageActionsResult(ListImageActionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListImageActionsResult(
            items=self.items)


def list_image_actions(image_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListImageActionsResult:
    """
    Use this data source to access information about an existing resource.

    :param str image_id: A unique number that can be used to identify and reference a specific image.
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:images/v2:listImageActions', __args__, opts=opts, typ=ListImageActionsResult).value

    return AwaitableListImageActionsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_image_actions)
def list_image_actions_output(image_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListImageActionsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str image_id: A unique number that can be used to identify and reference a specific image.
    """
    ...
