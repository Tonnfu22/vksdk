package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMessagesAddChatUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesAddChatUserBuilder()

	b.ChatID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesAllowMessagesFromGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesAllowMessagesFromGroupBuilder()

	b.GroupID(1)
	b.Key("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["key"], "text")
}

func TestMessagesCreateChatBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesCreateChatBuilder()

	b.UserIDs([]int{1})
	b.Title("text")

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["title"], "text")
}

func TestMessagesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteBuilder()

	b.MessageIDs([]int{1})
	b.PeerID(1)
	b.ConversationMessageIDs([]int{1})
	b.Spam(true)
	b.GroupID(1)
	b.DeleteForAll(true)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["conversation_message_ids"], []int{1})
	assert.Equal(t, b.Params["spam"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["delete_for_all"], true)
}

func TestMessagesDeleteChatPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteChatPhotoBuilder()

	b.ChatID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesDeleteConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteConversationBuilder()

	b.UserID(1)
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesDenyMessagesFromGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDenyMessagesFromGroupBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesEditBuilder()

	b.PeerID(1)
	b.Message("text")
	b.MessageID(1)
	b.Lat(1.1)
	b.Long(1.1)
	b.Attachment("text")
	b.KeepForwardMessages(true)
	b.KeepSnippets(true)
	b.GroupID(1)
	b.Keyboard("text")
	b.Template("text")
	b.ConversationMessageID(1)
	b.DontParseLinks(true)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["message_id"], 1)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["attachment"], "text")
	assert.Equal(t, b.Params["keep_forward_messages"], true)
	assert.Equal(t, b.Params["keep_snippets"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["keyboard"], "text")
	assert.Equal(t, b.Params["template"], "text")
	assert.Equal(t, b.Params["conversation_message_id"], 1)
	assert.Equal(t, b.Params["dont_parse_links"], true)
}

func TestMessagesEditChatBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesEditChatBuilder()

	b.ChatID(1)
	b.Title("text")

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
}

func TestMessagesForceCallFinishBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesForceCallFinishBuilder()

	b.CallID("text")

	assert.Equal(t, b.Params["call_id"], "text")
}

func TestMessagesGetByConversationMessageIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetByConversationMessageIDBuilder()

	b.PeerID(1)
	b.ConversationMessageIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["conversation_message_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetByIDBuilder()

	b.MessageIDs([]int{1})
	b.PreviewLength(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetChatPreviewBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetChatPreviewBuilder()

	b.PeerID(1)
	b.Link("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestMessagesGetConversationMembersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationMembersBuilder()

	b.PeerID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetConversationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationsBuilder()

	b.Offset(1)
	b.Count(1)
	b.Filter("text")
	b.Extended(true)
	b.StartMessageID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["filter"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetConversationsByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationsByIDBuilder()

	b.PeerIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetHistoryBuilder()

	b.Offset(1)
	b.Count(1)
	b.UserID(1)
	b.PeerID(1)
	b.StartMessageID(1)
	b.Rev(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["rev"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetHistoryAttachmentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetHistoryAttachmentsBuilder()

	b.PeerID(1)
	b.MediaType("text")
	b.StartFrom("text")
	b.Count(1)
	b.PhotoSizes(true)
	b.Fields([]string{"test"})
	b.GroupID(1)
	b.PreserveOrder(true)
	b.MaxForwardsLevel(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["media_type"], "text")
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["photo_sizes"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["preserve_order"], true)
	assert.Equal(t, b.Params["max_forwards_level"], 1)
}

func TestMessagesGetIntentUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetIntentUsersBuilder()

	b.Intent("test")
	b.SubscribeID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.NameCase([]string{"test"})
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["intent"], "test")
	assert.Equal(t, b.Params["subscribe_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["name_case"], []string{"test"})
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestMessagesGetInviteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetInviteLinkBuilder()

	b.PeerID(1)
	b.Reset(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["reset"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetLastActivityBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLastActivityBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesGetLongPollHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLongPollHistoryBuilder()

	b.Ts(1)
	b.Pts(1)
	b.PreviewLength(1)
	b.Onlines(true)
	b.Fields([]string{"test"})
	b.EventsLimit(1)
	b.MsgsLimit(1)
	b.MaxMsgID(1)
	b.GroupID(1)
	b.LpVersion(1)
	b.LastN(1)
	b.Credentials(true)

	assert.Equal(t, b.Params["ts"], 1)
	assert.Equal(t, b.Params["pts"], 1)
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["onlines"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["events_limit"], 1)
	assert.Equal(t, b.Params["msgs_limit"], 1)
	assert.Equal(t, b.Params["max_msg_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["lp_version"], 1)
	assert.Equal(t, b.Params["last_n"], 1)
	assert.Equal(t, b.Params["credentials"], true)
}

func TestMessagesGetLongPollServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLongPollServerBuilder()

	b.NeedPts(true)
	b.GroupID(1)
	b.LpVersion(1)

	assert.Equal(t, b.Params["need_pts"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["lp_version"], 1)
}

func TestMessagesIsMessagesFromGroupAllowedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesIsMessagesFromGroupAllowedBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesJoinChatByInviteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesJoinChatByInviteLinkBuilder()

	b.Link("text")

	assert.Equal(t, b.Params["link"], "text")
}

func TestMessagesMarkAsAnsweredConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsAnsweredConversationBuilder()

	b.PeerID(1)
	b.Answered(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["answered"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesMarkAsImportantBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsImportantBuilder()

	b.MessageIDs([]int{1})
	b.Important(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["important"], 1)
}

func TestMessagesMarkAsImportantConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsImportantConversationBuilder()

	b.PeerID(1)
	b.Important(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["important"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesMarkAsReadBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsReadBuilder()

	b.MessageIDs([]int{1})
	b.PeerID(1)
	b.StartMessageID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesPinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesPinBuilder()

	b.PeerID(1)
	b.MessageID(1)
	b.ConversationMessageID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["message_id"], 1)
	assert.Equal(t, b.Params["conversation_message_id"], 1)
}

func TestMessagesRemoveChatUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesRemoveChatUserBuilder()

	b.ChatID(1)
	b.UserID(1)
	b.MemberID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["member_id"], 1)
}

func TestMessagesRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesRestoreBuilder()

	b.MessageID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["message_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSearchBuilder()

	b.Q("text")
	b.PeerID(1)
	b.Date(1)
	b.PreviewLength(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"text"})
	b.GroupID(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["date"], 1)
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"text"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSearchConversationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSearchConversationsBuilder()

	b.Q("text")
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSendBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSendBuilder()

	b.UserID(1)
	b.RandomID(1)
	b.PeerID(1)
	b.Domain("text")
	b.ChatID(1)
	b.UserIDs([]int{1})
	b.PeerIDs([]int{1})
	b.Message("text")
	b.Lat(1.1)
	b.Long(1.1)
	b.Attachment("text")
	b.ReplyTo(1)
	b.ForwardMessages([]int{1})
	b.Forward("text")
	b.StickerID(1)
	b.GroupID(1)
	b.Keyboard("text")
	b.Template("text")
	b.Payload("text")
	b.ContentSource("text")
	b.DontParseLinks(true)
	b.DisableMentions(true)
	b.Intent("text")
	b.SubscribeID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["random_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["domain"], "text")
	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["peer_ids"], []int{1})
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["attachment"], "text")
	assert.Equal(t, b.Params["reply_to"], 1)
	assert.Equal(t, b.Params["forward_messages"], []int{1})
	assert.Equal(t, b.Params["forward"], "text")
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["keyboard"], "text")
	assert.Equal(t, b.Params["template"], "text")
	assert.Equal(t, b.Params["payload"], "text")
	assert.Equal(t, b.Params["content_source"], "text")
	assert.Equal(t, b.Params["dont_parse_links"], true)
	assert.Equal(t, b.Params["disable_mentions"], true)
	assert.Equal(t, b.Params["intent"], "text")
	assert.Equal(t, b.Params["subscribe_id"], 1)
}

func TestMessagesSendMessageEventAnswerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSendMessageEventAnswerBuilder()

	b.EventID("text")
	b.UserID(1)
	b.PeerID(1)
	b.EventData("text")

	assert.Equal(t, b.Params["event_id"], "text")
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["event_data"], "text")
}

func TestMessagesSetActivityBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSetActivityBuilder()

	b.UserID(1)
	b.Type("text")
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSetChatPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSetChatPhotoBuilder()

	b.File("text")

	assert.Equal(t, b.Params["file"], "text")
}

func TestMessagesStartCallBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesStartCallBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesUnpinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesUnpinBuilder()

	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}
