package types

type User struct {
	ID               string `json:"id"`
	GuildID          string `json:"guild_id"`
	PlayerTag        string `json:"player_tag"`
	Verified         bool   `json:"verified"`
	Premium          bool   `json:"premium"`
	PremiumExpiresAt int    `json:"premium_expires_at"`
	FreeUses         int    `json:"free_uses"`
	CreatedAt        int    `json:"created_at"`
}

type Verification struct {
	ID        string `json:"id"`
	PlayerTag string `json:"player_tag"`
	NameColor string `json:"name_color"`
	CreatedAt int    `json:"created_at"`
}

type Suspension struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Reason    string `json:"reason"`
	CreatedAt int    `json:"created_at"`
}

type Guild struct {
	ID                        string `json:"id"`
	OwnerID                   string `json:"owner_id"`
	NotificationChannelId     string `json:"notification_channel_id"`
	ClanTag                   string `json:"clan_tag"`
	PresidentRoleId           string `json:"president_role_id"`
	VicePresidentRoleId       string `json:"vice_president_role_id"`
	SeniorRoleId              string `json:"senior_role_id"`
	MemberRoleId              string `json:"member_role_id"`
	LastRoleSyncAt            int    `json:"last_role_sync_at"`
	LastLeaderboardRoleSyncAt int    `json:"last_leaderboard_role_sync_at"`
	ReplaceInviteLinks        bool   `json:"replace_invite_links"`
	CreatedAt                 int    `json:"created_at"`
}
