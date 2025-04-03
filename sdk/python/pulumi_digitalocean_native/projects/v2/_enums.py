# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'ProjectBaseEnvironment',
    'ResourceStatus',
]


class ProjectBaseEnvironment(builtins.str, Enum):
    """
    The environment of the project's resources.
    """
    DEVELOPMENT = "Development"
    STAGING = "Staging"
    PRODUCTION = "Production"


class ResourceStatus(builtins.str, Enum):
    """
    The status of assigning and fetching the resources.
    """
    OK = "ok"
    NOT_FOUND = "not_found"
    ASSIGNED = "assigned"
    ALREADY_ASSIGNED = "already_assigned"
    SERVICE_DOWN = "service_down"
