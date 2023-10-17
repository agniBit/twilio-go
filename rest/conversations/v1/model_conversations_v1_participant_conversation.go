/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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
// ConversationsV1ParticipantConversation struct for ConversationsV1ParticipantConversation
type ConversationsV1ParticipantConversation struct {
		// The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
		// The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
		// The unique ID of the [Participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource).
	ParticipantSid *string `json:"participant_sid,omitempty"`
		// The unique string that identifies the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource).
	ParticipantUserSid *string `json:"participant_user_sid,omitempty"`
		// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
	ParticipantIdentity *string `json:"participant_identity,omitempty"`
		// Information about how this participant exchanges messages with the conversation. A JSON parameter consisting of type and address fields of the participant.
	ParticipantMessagingBinding *interface{} `json:"participant_messaging_binding,omitempty"`
		// The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) this Participant belongs to.
	ConversationSid *string `json:"conversation_sid,omitempty"`
		// An application-defined string that uniquely identifies the Conversation resource.
	ConversationUniqueName *string `json:"conversation_unique_name,omitempty"`
		// The human-readable name of this conversation, limited to 256 characters. Optional.
	ConversationFriendlyName *string `json:"conversation_friendly_name,omitempty"`
		// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	ConversationAttributes *string `json:"conversation_attributes,omitempty"`
		// The date that this conversation was created, given in ISO 8601 format.
	ConversationDateCreated *time.Time `json:"conversation_date_created,omitempty"`
		// The date that this conversation was last updated, given in ISO 8601 format.
	ConversationDateUpdated *time.Time `json:"conversation_date_updated,omitempty"`
		// Identity of the creator of this Conversation.
	ConversationCreatedBy *string `json:"conversation_created_by,omitempty"`
	ConversationState *string `json:"conversation_state,omitempty"`
		// Timer date values representing state update for this conversation.
	ConversationTimers *interface{} `json:"conversation_timers,omitempty"`
		// Contains absolute URLs to access the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) and [conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) of this conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
}


