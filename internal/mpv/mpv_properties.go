package mpv

/*
#include <mpv/client.h>
*/
import "C"

type Format = C.mpv_format

const (
	FormatNone      Format = C.MPV_FORMAT_NONE
	FormatString    Format = C.MPV_FORMAT_STRING
	FormatOSDString Format = C.MPV_FORMAT_OSD_STRING
	FormatFlag      Format = C.MPV_FORMAT_FLAG
	FormatInt64     Format = C.MPV_FORMAT_INT64
	FormatDouble    Format = C.MPV_FORMAT_DOUBLE
	FormatNode      Format = C.MPV_FORMAT_NODE
)

type (
	Property   string
	PropertyID uint64
)

const (
	// Property Names
	PropertyTimePos       Property = "time-pos"
	PropertyTimeRemaining Property = "time-remaining"
	PropertyDuration      Property = "duration"
	PropertyPercentPos    Property = "percent-pos"

	PropertyPause      Property = "pause"
	PropertyVolume     Property = "volume"
	PropertyMute       Property = "mute"
	PropertyMediaTitle Property = "media-title"
	PropertyPath       Property = "path"
	PropertyFilename   Property = "filename"
	PropertyIdleActive Property = "idle-active"
	PropertyCoreIdle   Property = "core-idle"

	PropertyAudioDevice Property = "audio-device"
	PropertyVO          Property = "vo"
	PropertyVID         Property = "vid"
	PropertyAID         Property = "aid"

	PropertyMetadata         Property = "metadata"
	PropertyMetadataArtist   Property = "metadata/by-key/artist"
	PropertyMetadataTitle    Property = "metadata/by-key/title"
	PropertyMetadataAlbum    Property = "metadata/by-key/album"
	PropertyMetadataGenre    Property = "metadata/by-key/genre"
	PropertyMetadataDate     Property = "metadata/by-key/date"
	PropertyFilteredMetadata Property = "filtered-metadata"
)

const (
	// Unique numeric IDs for mpv_observe_property (reply_userdata)
	PropertyIDTimePos       PropertyID = 1
	PropertyIDTimeRemaining PropertyID = 2
	PropertyIDDuration      PropertyID = 3
	PropertyIDPercentPos    PropertyID = 4

	PropertyIDPause      PropertyID = 5
	PropertyIDVolume     PropertyID = 6
	PropertyIDMute       PropertyID = 7
	PropertyIDMediaTitle PropertyID = 8
	PropertyIDPath       PropertyID = 9
	PropertyIDFilename   PropertyID = 10
	PropertyIDIdleActive PropertyID = 11
	PropertyIDCoreIdle   PropertyID = 12

	PropertyIDAudioDevice PropertyID = 13
	PropertyIDVO          PropertyID = 14
	PropertyIDVID         PropertyID = 15
	PropertyIDAID         PropertyID = 16

	PropertyIDMetadata         PropertyID = 17
	PropertyIDMetadataArtist   PropertyID = 18
	PropertyIDMetadataTitle    PropertyID = 19
	PropertyIDMetadataAlbum    PropertyID = 20
	PropertyIDMetadataGenre    PropertyID = 21
	PropertyIDMetadataDate     PropertyID = 22
	PropertyIDFilteredMetadata PropertyID = 23
)

type PropertyInfo struct {
	ID     PropertyID
	Name   Property
	Format Format
}

var PropertyCatalog = map[PropertyID]PropertyInfo{
	PropertyIDTimePos:       {ID: PropertyIDTimePos, Name: PropertyTimePos, Format: FormatDouble},
	PropertyIDTimeRemaining: {ID: PropertyIDTimeRemaining, Name: PropertyTimeRemaining, Format: FormatDouble},
	PropertyIDDuration:      {ID: PropertyIDDuration, Name: PropertyDuration, Format: FormatDouble},
	PropertyIDPercentPos:    {ID: PropertyIDPercentPos, Name: PropertyPercentPos, Format: FormatDouble},

	PropertyIDPause:      {ID: PropertyIDPause, Name: PropertyPause, Format: FormatFlag},
	PropertyIDVolume:     {ID: PropertyIDVolume, Name: PropertyVolume, Format: FormatDouble},
	PropertyIDMute:       {ID: PropertyIDMute, Name: PropertyMute, Format: FormatFlag},
	PropertyIDMediaTitle: {ID: PropertyIDMediaTitle, Name: PropertyMediaTitle, Format: FormatString},
	PropertyIDPath:       {ID: PropertyIDPath, Name: PropertyPath, Format: FormatString},
	PropertyIDFilename:   {ID: PropertyIDFilename, Name: PropertyFilename, Format: FormatString},
	PropertyIDIdleActive: {ID: PropertyIDIdleActive, Name: PropertyIdleActive, Format: FormatFlag},
	PropertyIDCoreIdle:   {ID: PropertyIDCoreIdle, Name: PropertyCoreIdle, Format: FormatFlag},

	PropertyIDAudioDevice: {ID: PropertyIDAudioDevice, Name: PropertyAudioDevice, Format: FormatString},
	PropertyIDVO:          {ID: PropertyIDVO, Name: PropertyVO, Format: FormatString},
	PropertyIDVID:         {ID: PropertyIDVID, Name: PropertyVID, Format: FormatString},
	PropertyIDAID:         {ID: PropertyIDAID, Name: PropertyAID, Format: FormatString},

	PropertyIDMetadata:         {ID: PropertyIDMetadata, Name: PropertyMetadata, Format: FormatString},
	PropertyIDMetadataArtist:   {ID: PropertyIDMetadataArtist, Name: PropertyMetadataArtist, Format: FormatString},
	PropertyIDMetadataTitle:    {ID: PropertyIDMetadataTitle, Name: PropertyMetadataTitle, Format: FormatString},
	PropertyIDMetadataAlbum:    {ID: PropertyIDMetadataAlbum, Name: PropertyMetadataAlbum, Format: FormatString},
	PropertyIDMetadataGenre:    {ID: PropertyIDMetadataGenre, Name: PropertyMetadataGenre, Format: FormatString},
	PropertyIDMetadataDate:     {ID: PropertyIDMetadataDate, Name: PropertyMetadataDate, Format: FormatString},
	PropertyIDFilteredMetadata: {ID: PropertyIDFilteredMetadata, Name: PropertyFilteredMetadata, Format: FormatString},
}

func (p Property) String() string {
	return string(p)
}

func GetPropertyByName(name Property) (PropertyInfo, bool) {
	for _, info := range PropertyCatalog {
		if info.Name == name {
			return info, true
		}
	}
	return PropertyInfo{}, false
}
