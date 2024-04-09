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
    'ListProjectsResult',
    'AwaitableListProjectsResult',
    'list_projects',
    'list_projects_output',
]

@pulumi.output_type
class ListProjectsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListProjectsItems':
        return pulumi.get(self, "items")


class AwaitableListProjectsResult(ListProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectsResult(
            items=self.items)


def list_projects(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:projects/v2:listProjects', __args__, opts=opts, typ=ListProjectsResult).value

    return AwaitableListProjectsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_projects)
def list_projects_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListProjectsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
