# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ActionStatus',
    'ConvertStatus',
    'ConvertType',
    'ImageDistribution',
    'ImageRegionsItem',
    'ImageStatus',
    'ImageType',
    'ImagesCustomImageUpdateDistribution',
    'ImagesCustomPropertiesRegion',
    'TransferImageActionBaseType',
    'TransferPropertiesRegion',
    'TransferStatus',
]


class ActionStatus(str, Enum):
    """
    The current status of the action. This can be "in-progress", "completed", or "errored".
    """
    IN_PROGRESS = "in-progress"
    COMPLETED = "completed"
    ERRORED = "errored"


class ConvertStatus(str, Enum):
    """
    The current status of the action. This can be "in-progress", "completed", or "errored".
    """
    IN_PROGRESS = "in-progress"
    COMPLETED = "completed"
    ERRORED = "errored"


class ConvertType(str, Enum):
    """
    The action to be taken on the image. Can be either `convert` or `transfer`.
    """
    CONVERT = "convert"
    TRANSFER = "transfer"


class ImageDistribution(str, Enum):
    """
    The name of a custom image's distribution. Currently, the valid values are  `Arch Linux`, `CentOS`, `CoreOS`, `Debian`, `Fedora`, `Fedora Atomic`,  `FreeBSD`, `Gentoo`, `openSUSE`, `RancherOS`, `Rocky Linux`, `Ubuntu`, and `Unknown`.  Any other value will be accepted but ignored, and `Unknown` will be used in its place.
    """
    ARCH_LINUX = "Arch Linux"
    CENT_OS = "CentOS"
    CORE_OS = "CoreOS"
    DEBIAN = "Debian"
    FEDORA = "Fedora"
    FEDORA_ATOMIC = "Fedora Atomic"
    FREE_BSD = "FreeBSD"
    GENTOO = "Gentoo"
    OPEN_SUSE = "openSUSE"
    RANCHER_OS = "RancherOS"
    ROCKY_LINUX = "Rocky Linux"
    UBUNTU = "Ubuntu"
    UNKNOWN = "Unknown"


class ImageRegionsItem(str, Enum):
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


class ImageStatus(str, Enum):
    """
    A status string indicating the state of a custom image. This may be `NEW`,
     `available`, `pending`, `deleted`, or `retired`.
    """
    NEW = "NEW"
    AVAILABLE = "available"
    PENDING = "pending"
    DELETED = "deleted"
    RETIRED = "retired"


class ImageType(str, Enum):
    """
    Describes the kind of image. It may be one of `base`, `snapshot`, `backup`, `custom`, or `admin`. Respectively, this specifies whether an image is a DigitalOcean base OS image, user-generated Droplet snapshot, automatically created Droplet backup, user-provided virtual machine image, or an image used for DigitalOcean managed resources (e.g. DOKS worker nodes).
    """
    BASE = "base"
    SNAPSHOT = "snapshot"
    BACKUP = "backup"
    CUSTOM = "custom"
    ADMIN = "admin"


class ImagesCustomImageUpdateDistribution(str, Enum):
    """
    The name of a custom image's distribution. Currently, the valid values are  `Arch Linux`, `CentOS`, `CoreOS`, `Debian`, `Fedora`, `Fedora Atomic`,  `FreeBSD`, `Gentoo`, `openSUSE`, `RancherOS`, `Rocky Linux`, `Ubuntu`, and `Unknown`.  Any other value will be accepted but ignored, and `Unknown` will be used in its place.
    """
    ARCH_LINUX = "Arch Linux"
    CENT_OS = "CentOS"
    CORE_OS = "CoreOS"
    DEBIAN = "Debian"
    FEDORA = "Fedora"
    FEDORA_ATOMIC = "Fedora Atomic"
    FREE_BSD = "FreeBSD"
    GENTOO = "Gentoo"
    OPEN_SUSE = "openSUSE"
    RANCHER_OS = "RancherOS"
    ROCKY_LINUX = "Rocky Linux"
    UBUNTU = "Ubuntu"
    UNKNOWN = "Unknown"


class ImagesCustomPropertiesRegion(str, Enum):
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


class TransferImageActionBaseType(str, Enum):
    """
    The action to be taken on the image. Can be either `convert` or `transfer`.
    """
    CONVERT = "convert"
    TRANSFER = "transfer"


class TransferPropertiesRegion(str, Enum):
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


class TransferStatus(str, Enum):
    """
    The current status of the action. This can be "in-progress", "completed", or "errored".
    """
    IN_PROGRESS = "in-progress"
    COMPLETED = "completed"
    ERRORED = "errored"
