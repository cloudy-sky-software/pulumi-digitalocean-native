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
    'ListProjectsResourcesResult',
    'AwaitableListProjectsResourcesResult',
    'list_projects_resources',
    'list_projects_resources_output',
]

@pulumi.output_type
class ListProjectsResourcesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListProjectsResourcesItems':
        return pulumi.get(self, "items")


class AwaitableListProjectsResourcesResult(ListProjectsResourcesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectsResourcesResult(
            items=self.items)


def list_projects_resources(project_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectsResourcesResult:
    """
    Use this data source to access information about an existing resource.

    :param str project_id: A unique identifier for a project.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:projects/v2:listProjectsResources', __args__, opts=opts, typ=ListProjectsResourcesResult).value

    return AwaitableListProjectsResourcesResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_projects_resources)
def list_projects_resources_output(project_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListProjectsResourcesResult]:
    """
    Use this data source to access information about an existing resource.

    :param str project_id: A unique identifier for a project.
    """
    ...
