/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// UsageRecordThisMonthCategory the model 'UsageRecordThisMonthCategory'
type UsageRecordThisMonthCategory string

// List of usage_record_this_month_category
const (
	USAGERECORDTHISMONTHCATEGORY_AGENT_CONFERENCE                                       UsageRecordThisMonthCategory = "agent-conference"
	USAGERECORDTHISMONTHCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordThisMonthCategory = "answering-machine-detection"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordThisMonthCategory = "authy-authentications"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordThisMonthCategory = "authy-calls-outbound"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordThisMonthCategory = "authy-monthly-fees"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordThisMonthCategory = "authy-phone-intelligence"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordThisMonthCategory = "authy-phone-verifications"
	USAGERECORDTHISMONTHCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordThisMonthCategory = "authy-sms-outbound"
	USAGERECORDTHISMONTHCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordThisMonthCategory = "call-progess-events"
	USAGERECORDTHISMONTHCATEGORY_CALLERIDLOOKUPS                                        UsageRecordThisMonthCategory = "calleridlookups"
	USAGERECORDTHISMONTHCATEGORY_CALLS                                                  UsageRecordThisMonthCategory = "calls"
	USAGERECORDTHISMONTHCATEGORY_CALLS_CLIENT                                           UsageRecordThisMonthCategory = "calls-client"
	USAGERECORDTHISMONTHCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordThisMonthCategory = "calls-globalconference"
	USAGERECORDTHISMONTHCATEGORY_CALLS_INBOUND                                          UsageRecordThisMonthCategory = "calls-inbound"
	USAGERECORDTHISMONTHCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordThisMonthCategory = "calls-inbound-local"
	USAGERECORDTHISMONTHCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordThisMonthCategory = "calls-inbound-mobile"
	USAGERECORDTHISMONTHCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordThisMonthCategory = "calls-inbound-tollfree"
	USAGERECORDTHISMONTHCATEGORY_CALLS_OUTBOUND                                         UsageRecordThisMonthCategory = "calls-outbound"
	USAGERECORDTHISMONTHCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordThisMonthCategory = "calls-pay-verb-transactions"
	USAGERECORDTHISMONTHCATEGORY_CALLS_RECORDINGS                                       UsageRecordThisMonthCategory = "calls-recordings"
	USAGERECORDTHISMONTHCATEGORY_CALLS_SIP                                              UsageRecordThisMonthCategory = "calls-sip"
	USAGERECORDTHISMONTHCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordThisMonthCategory = "calls-sip-inbound"
	USAGERECORDTHISMONTHCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordThisMonthCategory = "calls-sip-outbound"
	USAGERECORDTHISMONTHCATEGORY_CARRIER_LOOKUPS                                        UsageRecordThisMonthCategory = "carrier-lookups"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS                                          UsageRecordThisMonthCategory = "conversations"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordThisMonthCategory = "conversations-api-requests"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordThisMonthCategory = "conversations-conversation-events"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordThisMonthCategory = "conversations-endpoint-connectivity"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordThisMonthCategory = "conversations-events"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordThisMonthCategory = "conversations-participant-events"
	USAGERECORDTHISMONTHCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordThisMonthCategory = "conversations-participants"
	USAGERECORDTHISMONTHCATEGORY_CPS                                                    UsageRecordThisMonthCategory = "cps"
	USAGERECORDTHISMONTHCATEGORY_FRAUD_LOOKUPS                                          UsageRecordThisMonthCategory = "fraud-lookups"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS                                            UsageRecordThisMonthCategory = "group-rooms"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordThisMonthCategory = "group-rooms-data-track"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordThisMonthCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordThisMonthCategory = "group-rooms-media-downloaded"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordThisMonthCategory = "group-rooms-media-recorded"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordThisMonthCategory = "group-rooms-media-routed"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordThisMonthCategory = "group-rooms-media-stored"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordThisMonthCategory = "group-rooms-participant-minutes"
	USAGERECORDTHISMONTHCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordThisMonthCategory = "group-rooms-recorded-minutes"
	USAGERECORDTHISMONTHCATEGORY_IMP_V1_USAGE                                           UsageRecordThisMonthCategory = "imp-v1-usage"
	USAGERECORDTHISMONTHCATEGORY_LOOKUPS                                                UsageRecordThisMonthCategory = "lookups"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE                                            UsageRecordThisMonthCategory = "marketplace"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordThisMonthCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordThisMonthCategory = "marketplace-cadence-transcription"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordThisMonthCategory = "marketplace-cadence-translation"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordThisMonthCategory = "marketplace-capio-speech-to-text"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordThisMonthCategory = "marketplace-convriza-ababa"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordThisMonthCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordThisMonthCategory = "marketplace-digital-segment-business-info"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordThisMonthCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordThisMonthCategory = "marketplace-google-speech-to-text"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordThisMonthCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordThisMonthCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordThisMonthCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordThisMonthCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordThisMonthCategory = "marketplace-icehook-systems-scout"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordThisMonthCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordThisMonthCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordThisMonthCategory = "marketplace-marchex-cleancall"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordThisMonthCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordThisMonthCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordThisMonthCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordThisMonthCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordThisMonthCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordThisMonthCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordThisMonthCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordThisMonthCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordThisMonthCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordThisMonthCategory = "marketplace-telo-opencnam"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordThisMonthCategory = "marketplace-truecnam-true-spam"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordThisMonthCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordThisMonthCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordThisMonthCategory = "marketplace-voicebase-pci"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordThisMonthCategory = "marketplace-voicebase-transcription"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordThisMonthCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordThisMonthCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordThisMonthCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordThisMonthCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordThisMonthCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordThisMonthCategory = "marketplace-wolfram-short-answer"
	USAGERECORDTHISMONTHCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordThisMonthCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDTHISMONTHCATEGORY_MEDIASTORAGE                                           UsageRecordThisMonthCategory = "mediastorage"
	USAGERECORDTHISMONTHCATEGORY_MMS                                                    UsageRecordThisMonthCategory = "mms"
	USAGERECORDTHISMONTHCATEGORY_MMS_INBOUND                                            UsageRecordThisMonthCategory = "mms-inbound"
	USAGERECORDTHISMONTHCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordThisMonthCategory = "mms-inbound-longcode"
	USAGERECORDTHISMONTHCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordThisMonthCategory = "mms-inbound-shortcode"
	USAGERECORDTHISMONTHCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordThisMonthCategory = "mms-messages-carrierfees"
	USAGERECORDTHISMONTHCATEGORY_MMS_OUTBOUND                                           UsageRecordThisMonthCategory = "mms-outbound"
	USAGERECORDTHISMONTHCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordThisMonthCategory = "mms-outbound-longcode"
	USAGERECORDTHISMONTHCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordThisMonthCategory = "mms-outbound-shortcode"
	USAGERECORDTHISMONTHCATEGORY_MONITOR_READS                                          UsageRecordThisMonthCategory = "monitor-reads"
	USAGERECORDTHISMONTHCATEGORY_MONITOR_STORAGE                                        UsageRecordThisMonthCategory = "monitor-storage"
	USAGERECORDTHISMONTHCATEGORY_MONITOR_WRITES                                         UsageRecordThisMonthCategory = "monitor-writes"
	USAGERECORDTHISMONTHCATEGORY_NOTIFY                                                 UsageRecordThisMonthCategory = "notify"
	USAGERECORDTHISMONTHCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordThisMonthCategory = "notify-actions-attempts"
	USAGERECORDTHISMONTHCATEGORY_NOTIFY_CHANNELS                                        UsageRecordThisMonthCategory = "notify-channels"
	USAGERECORDTHISMONTHCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordThisMonthCategory = "number-format-lookups"
	USAGERECORDTHISMONTHCATEGORY_PCHAT                                                  UsageRecordThisMonthCategory = "pchat"
	USAGERECORDTHISMONTHCATEGORY_PCHAT_USERS                                            UsageRecordThisMonthCategory = "pchat-users"
	USAGERECORDTHISMONTHCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordThisMonthCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDTHISMONTHCATEGORY_PFAX                                                   UsageRecordThisMonthCategory = "pfax"
	USAGERECORDTHISMONTHCATEGORY_PFAX_MINUTES                                           UsageRecordThisMonthCategory = "pfax-minutes"
	USAGERECORDTHISMONTHCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordThisMonthCategory = "pfax-minutes-inbound"
	USAGERECORDTHISMONTHCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordThisMonthCategory = "pfax-minutes-outbound"
	USAGERECORDTHISMONTHCATEGORY_PFAX_PAGES                                             UsageRecordThisMonthCategory = "pfax-pages"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS                                           UsageRecordThisMonthCategory = "phonenumbers"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_CPS                                       UsageRecordThisMonthCategory = "phonenumbers-cps"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordThisMonthCategory = "phonenumbers-emergency"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordThisMonthCategory = "phonenumbers-local"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordThisMonthCategory = "phonenumbers-mobile"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordThisMonthCategory = "phonenumbers-setups"
	USAGERECORDTHISMONTHCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordThisMonthCategory = "phonenumbers-tollfree"
	USAGERECORDTHISMONTHCATEGORY_PREMIUMSUPPORT                                         UsageRecordThisMonthCategory = "premiumsupport"
	USAGERECORDTHISMONTHCATEGORY_PROXY                                                  UsageRecordThisMonthCategory = "proxy"
	USAGERECORDTHISMONTHCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordThisMonthCategory = "proxy-active-sessions"
	USAGERECORDTHISMONTHCATEGORY_PSTNCONNECTIVITY                                       UsageRecordThisMonthCategory = "pstnconnectivity"
	USAGERECORDTHISMONTHCATEGORY_PV                                                     UsageRecordThisMonthCategory = "pv"
	USAGERECORDTHISMONTHCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordThisMonthCategory = "pv-composition-media-downloaded"
	USAGERECORDTHISMONTHCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordThisMonthCategory = "pv-composition-media-encrypted"
	USAGERECORDTHISMONTHCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordThisMonthCategory = "pv-composition-media-stored"
	USAGERECORDTHISMONTHCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordThisMonthCategory = "pv-composition-minutes"
	USAGERECORDTHISMONTHCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordThisMonthCategory = "pv-recording-compositions"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordThisMonthCategory = "pv-room-participants"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordThisMonthCategory = "pv-room-participants-au1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordThisMonthCategory = "pv-room-participants-br1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordThisMonthCategory = "pv-room-participants-ie1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordThisMonthCategory = "pv-room-participants-jp1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordThisMonthCategory = "pv-room-participants-sg1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordThisMonthCategory = "pv-room-participants-us1"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordThisMonthCategory = "pv-room-participants-us2"
	USAGERECORDTHISMONTHCATEGORY_PV_ROOMS                                               UsageRecordThisMonthCategory = "pv-rooms"
	USAGERECORDTHISMONTHCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordThisMonthCategory = "pv-sip-endpoint-registrations"
	USAGERECORDTHISMONTHCATEGORY_RECORDINGS                                             UsageRecordThisMonthCategory = "recordings"
	USAGERECORDTHISMONTHCATEGORY_RECORDINGSTORAGE                                       UsageRecordThisMonthCategory = "recordingstorage"
	USAGERECORDTHISMONTHCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordThisMonthCategory = "rooms-group-bandwidth"
	USAGERECORDTHISMONTHCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordThisMonthCategory = "rooms-group-minutes"
	USAGERECORDTHISMONTHCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordThisMonthCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES                                             UsageRecordThisMonthCategory = "shortcodes"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordThisMonthCategory = "shortcodes-customerowned"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordThisMonthCategory = "shortcodes-mms-enablement"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_MPS                                         UsageRecordThisMonthCategory = "shortcodes-mps"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_RANDOM                                      UsageRecordThisMonthCategory = "shortcodes-random"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_UK                                          UsageRecordThisMonthCategory = "shortcodes-uk"
	USAGERECORDTHISMONTHCATEGORY_SHORTCODES_VANITY                                      UsageRecordThisMonthCategory = "shortcodes-vanity"
	USAGERECORDTHISMONTHCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordThisMonthCategory = "small-group-rooms"
	USAGERECORDTHISMONTHCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordThisMonthCategory = "small-group-rooms-data-track"
	USAGERECORDTHISMONTHCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordThisMonthCategory = "small-group-rooms-participant-minutes"
	USAGERECORDTHISMONTHCATEGORY_SMS                                                    UsageRecordThisMonthCategory = "sms"
	USAGERECORDTHISMONTHCATEGORY_SMS_INBOUND                                            UsageRecordThisMonthCategory = "sms-inbound"
	USAGERECORDTHISMONTHCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordThisMonthCategory = "sms-inbound-longcode"
	USAGERECORDTHISMONTHCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordThisMonthCategory = "sms-inbound-shortcode"
	USAGERECORDTHISMONTHCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordThisMonthCategory = "sms-messages-carrierfees"
	USAGERECORDTHISMONTHCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordThisMonthCategory = "sms-messages-features"
	USAGERECORDTHISMONTHCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordThisMonthCategory = "sms-messages-features-senderid"
	USAGERECORDTHISMONTHCATEGORY_SMS_OUTBOUND                                           UsageRecordThisMonthCategory = "sms-outbound"
	USAGERECORDTHISMONTHCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordThisMonthCategory = "sms-outbound-content-inspection"
	USAGERECORDTHISMONTHCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordThisMonthCategory = "sms-outbound-longcode"
	USAGERECORDTHISMONTHCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordThisMonthCategory = "sms-outbound-shortcode"
	USAGERECORDTHISMONTHCATEGORY_SPEECH_RECOGNITION                                     UsageRecordThisMonthCategory = "speech-recognition"
	USAGERECORDTHISMONTHCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordThisMonthCategory = "studio-engagements"
	USAGERECORDTHISMONTHCATEGORY_SYNC                                                   UsageRecordThisMonthCategory = "sync"
	USAGERECORDTHISMONTHCATEGORY_SYNC_ACTIONS                                           UsageRecordThisMonthCategory = "sync-actions"
	USAGERECORDTHISMONTHCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordThisMonthCategory = "sync-endpoint-hours"
	USAGERECORDTHISMONTHCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordThisMonthCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDTHISMONTHCATEGORY_TASKROUTER_TASKS                                       UsageRecordThisMonthCategory = "taskrouter-tasks"
	USAGERECORDTHISMONTHCATEGORY_TOTALPRICE                                             UsageRecordThisMonthCategory = "totalprice"
	USAGERECORDTHISMONTHCATEGORY_TRANSCRIPTIONS                                         UsageRecordThisMonthCategory = "transcriptions"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_CPS                                           UsageRecordThisMonthCategory = "trunking-cps"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordThisMonthCategory = "trunking-emergency-calls"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordThisMonthCategory = "trunking-origination"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordThisMonthCategory = "trunking-origination-local"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordThisMonthCategory = "trunking-origination-mobile"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordThisMonthCategory = "trunking-origination-tollfree"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordThisMonthCategory = "trunking-recordings"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_SECURE                                        UsageRecordThisMonthCategory = "trunking-secure"
	USAGERECORDTHISMONTHCATEGORY_TRUNKING_TERMINATION                                   UsageRecordThisMonthCategory = "trunking-termination"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES                                          UsageRecordThisMonthCategory = "turnmegabytes"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordThisMonthCategory = "turnmegabytes-australia"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordThisMonthCategory = "turnmegabytes-brasil"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordThisMonthCategory = "turnmegabytes-germany"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordThisMonthCategory = "turnmegabytes-india"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordThisMonthCategory = "turnmegabytes-ireland"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordThisMonthCategory = "turnmegabytes-japan"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordThisMonthCategory = "turnmegabytes-singapore"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordThisMonthCategory = "turnmegabytes-useast"
	USAGERECORDTHISMONTHCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordThisMonthCategory = "turnmegabytes-uswest"
	USAGERECORDTHISMONTHCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordThisMonthCategory = "twilio-interconnect"
	USAGERECORDTHISMONTHCATEGORY_VERIFY_PUSH                                            UsageRecordThisMonthCategory = "verify-push"
	USAGERECORDTHISMONTHCATEGORY_VIDEO_RECORDINGS                                       UsageRecordThisMonthCategory = "video-recordings"
	USAGERECORDTHISMONTHCATEGORY_VOICE_INSIGHTS                                         UsageRecordThisMonthCategory = "voice-insights"
	USAGERECORDTHISMONTHCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordThisMonthCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDTHISMONTHCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordThisMonthCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDTHISMONTHCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordThisMonthCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDTHISMONTHCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordThisMonthCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS                                               UsageRecordThisMonthCategory = "wireless"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_ORDERS                                        UsageRecordThisMonthCategory = "wireless-orders"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordThisMonthCategory = "wireless-orders-artwork"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordThisMonthCategory = "wireless-orders-bulk"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordThisMonthCategory = "wireless-orders-esim"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordThisMonthCategory = "wireless-orders-starter"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE                                         UsageRecordThisMonthCategory = "wireless-usage"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordThisMonthCategory = "wireless-usage-commands"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordThisMonthCategory = "wireless-usage-commands-africa"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordThisMonthCategory = "wireless-usage-commands-asia"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordThisMonthCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordThisMonthCategory = "wireless-usage-commands-europe"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordThisMonthCategory = "wireless-usage-commands-home"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordThisMonthCategory = "wireless-usage-commands-northamerica"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordThisMonthCategory = "wireless-usage-commands-oceania"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordThisMonthCategory = "wireless-usage-commands-roaming"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordThisMonthCategory = "wireless-usage-data"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordThisMonthCategory = "wireless-usage-data-africa"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordThisMonthCategory = "wireless-usage-data-asia"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordThisMonthCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordThisMonthCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordThisMonthCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordThisMonthCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordThisMonthCategory = "wireless-usage-data-europe"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordThisMonthCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordThisMonthCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordThisMonthCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordThisMonthCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordThisMonthCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordThisMonthCategory = "wireless-usage-data-northamerica"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordThisMonthCategory = "wireless-usage-data-oceania"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordThisMonthCategory = "wireless-usage-data-pooled"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordThisMonthCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordThisMonthCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordThisMonthCategory = "wireless-usage-mrc"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordThisMonthCategory = "wireless-usage-mrc-custom"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordThisMonthCategory = "wireless-usage-mrc-individual"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordThisMonthCategory = "wireless-usage-mrc-pooled"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordThisMonthCategory = "wireless-usage-mrc-suspended"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordThisMonthCategory = "wireless-usage-sms"
	USAGERECORDTHISMONTHCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordThisMonthCategory = "wireless-usage-voice"
)