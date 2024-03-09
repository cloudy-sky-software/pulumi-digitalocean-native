# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ActionStatus',
    'SnapshotsPropertiesResourceType',
    'VolumeActionPostBaseRegion',
    'VolumeActionPostBaseType',
]


class ActionStatus(str, Enum):
    """
    The current status of the action. This can be "in-progress", "completed", or "errored".
    """
    IN_PROGRESS = "in-progress"
    COMPLETED = "completed"
    ERRORED = "errored"


class SnapshotsPropertiesResourceType(str, Enum):
    """
    The type of resource that the snapshot originated from.
    """
    DROPLET = "droplet"
    VOLUME = "volume"


class VolumeActionPostBaseRegion(str, Enum):
    """
    The slug identifier for the region where the resource will initially be  available.
    """
    AMS1 = "ams1"
    AMS2 = "ams2"
    AMS3 = "ams3"
    BLR1 = "blr1"
    FRA1 = "fra1"
    LON1 = "lon1"
    NYC1 = "nyc1"
    NYC2 = "nyc2"
    NYC3 = "nyc3"
    SFO1 = "sfo1"
    SFO2 = "sfo2"
    SFO3 = "sfo3"
    SGP1 = "sgp1"
    TOR1 = "tor1"


class VolumeActionPostBaseType(str, Enum):
    """
    The volume action to initiate.
    """
    ATTACH = "attach"
    DETACH = "detach"
    RESIZE = "resize"