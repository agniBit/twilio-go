/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi
import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)
// VideoV1Composition struct for VideoV1Composition
type VideoV1Composition struct {
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Composition resource.
	AccountSid *string `json:"account_sid,omitempty"`
	Status *string `json:"status,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the composition's media processing task finished, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCompleted *time.Time `json:"date_completed,omitempty"`
		// The date and time in GMT when the composition generated media was deleted, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateDeleted *time.Time `json:"date_deleted,omitempty"`
		// The unique string that we created to identify the Composition resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the Group Room that generated the audio and video tracks used in the composition. All media sources included in a composition must belong to the same Group Room.
	RoomSid *string `json:"room_sid,omitempty"`
		// The array of track names to include in the composition. The composition includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this property can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`.
	AudioSources *[]string `json:"audio_sources,omitempty"`
		// The array of track names to exclude from the composition. The composition includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this property can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty.
	AudioSourcesExcluded *[]string `json:"audio_sources_excluded,omitempty"`
		// An object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	VideoLayout *interface{} `json:"video_layout,omitempty"`
		// The dimensions of the video image in pixels expressed as columns (width) and rows (height). The string's format is `{width}x{height}`, such as `640x480`.
	Resolution *string `json:"resolution,omitempty"`
		// Whether to remove intervals with no media, as specified in the POST request that created the composition. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Trim *bool `json:"trim,omitempty"`
	Format *string `json:"format,omitempty"`
		// The average bit rate of the composition's media.
	Bitrate *int `json:"bitrate,omitempty"`
		// The size of the composed media file in bytes.
	Size *int64 `json:"size,omitempty"`
		// The duration of the composition's media file in seconds.
	Duration *int `json:"duration,omitempty"`
		// The URL of the media file associated with the composition when stored externally. See [External S3 Compositions](/docs/video/api/external-s3-compositions) for more details.
	MediaExternalLocation *string `json:"media_external_location,omitempty"`
		// The URL called using the `status_callback_method` to send status information on every composition event.
	StatusCallback *string `json:"status_callback,omitempty"`
		// The HTTP method used to call `status_callback`. Can be: `POST` or `GET`, defaults to `POST`.
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
		// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
		// The URL of the media file associated with the composition.
	Links *map[string]interface{} `json:"links,omitempty"`
}


