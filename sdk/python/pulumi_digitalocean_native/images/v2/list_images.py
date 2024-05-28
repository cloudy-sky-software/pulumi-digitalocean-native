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
    'ListImagesResult',
    'AwaitableListImagesResult',
    'list_images',
    'list_images_output',
]

@pulumi.output_type
class ListImagesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListImagesItems':
        return pulumi.get(self, "items")


class AwaitableListImagesResult(ListImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListImagesResult(
            items=self.items)


def list_images(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListImagesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:images/v2:listImages', __args__, opts=opts, typ=ListImagesResult).value

    return AwaitableListImagesResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_images)
def list_images_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListImagesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...