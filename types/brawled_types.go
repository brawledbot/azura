package types

type Profile struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	NameColor string `json:"nameColor"`
	Icon      struct {
		ID int `json:"id"`
	} `json:"icon"`
	Trophies              int `json:"trophies"`
	HighestTrophies       int `json:"highestTrophies"`
	ExpLevel              int `json:"expLevel"`
	ExpPoints             int `json:"expPoints"`
	ThreeVsThreeVictories int `json:"3vs3Victories"`
	SoloVictories         int `json:"soloVictories"`
	DuoVictories          int `json:"duoVictories"`
	BrawledProfile        struct {
		ID        string `json:"id"`
		PlayerTag string `json:"player_tag"`
		Verified  bool   `json:"verified"`
		Premium   bool   `json:"premium"`
		CreatedAt int64  `json:"created_at"`
	} `json:"brawledProfile"`
	Club struct {
		Tag              string `json:"tag"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		Type             string `json:"type"`
		BadgeID          int    `json:"badgeId"`
		RequiredTrophies int    `json:"requiredTrophies"`
		Trophies         int    `json:"trophies"`
		Members          []struct {
			Tag       string `json:"tag"`
			Name      string `json:"name"`
			NameColor string `json:"nameColor"`
			Role      string `json:"role"`
			Trophies  int    `json:"trophies"`
			Icon      struct {
				ID int `json:"id"`
			} `json:"icon"`
		} `json:"members"`
	} `json:"club"`
	Calculated struct {
		FavouriteBrawler    int    `json:"favouriteBrawler"`
		WinRate             string `json:"winRate"`
		WinStreak           int    `json:"winStreak"`
		LeagueRank          string `json:"leagueRank"`
		LeagueRankSub       string `json:"leagueRankSub"`
		LeagueRankInt       int    `json:"leagueRankInt"`
		HoursPlayed         string `json:"hoursPlayed"`
		LeaderboardPosition int    `json:"leaderboardPosition"`
		BattleCard          struct {
			FirstIcon  int `json:"FirstIcon"`
			SecondIcon int `json:"SecondIcon"`
			Emoji      int `json:"Emoji"`
			Title      int `json:"Title"`
		} `json:"battleCard"`
	} `json:"calculated"`
	Brawlers []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Power           int    `json:"power"`
		Rank            int    `json:"rank"`
		Trophies        int    `json:"trophies"`
		HighestTrophies int    `json:"highestTrophies"`
		Rarity          struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"rarity"`
		Brawled struct {
			DiscordEmoji string `json:"discordEmoji"`
			BrawlerImage string `json:"brawlerImage"`
		} `json:"brawled"`
	} `json:"brawlers"`
	BattleLogs []struct {
		BattleTime string `json:"battleTime"`
		Event      struct {
			ID   int    `json:"id"`
			Mode string `json:"mode"`
			Map  string `json:"map"`
		} `json:"event"`
		Battle struct {
			Mode         string `json:"mode"`
			Type         string `json:"type"`
			Rank         int    `json:"rank"`
			TrophyChange int    `json:"trophyChange"`
			Teams        [][]struct {
				Tag     string `json:"tag"`
				Name    string `json:"name"`
				Brawler struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Power    int    `json:"power"`
					Trophies int    `json:"trophies"`
				} `json:"brawler"`
			} `json:"teams"`
		} `json:"battle"`
	} `json:"battleLogs"`
}

type Club struct {
	Tag              string `json:"tag"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Type             string `json:"type"`
	BadgeID          int    `json:"badgeId"`
	RequiredTrophies int    `json:"requiredTrophies"`
	Trophies         int    `json:"trophies"`
	Members          []struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		NameColor string `json:"nameColor"`
		Role      string `json:"role"`
		Trophies  int    `json:"trophies"`
		Icon      struct {
			ID int `json:"id"`
		} `json:"icon"`
	} `json:"members"`
}

type Leaderboard struct {
	Position int    `json:"position"`
	Tag      string `json:"tag"`
	Name     string `json:"name"`
	Trophies int    `json:"trophies"`
}

type Brawler struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	ImageURL2   string `json:"imageUrl2"`
	ImageURL3   string `json:"imageUrl3"`
	Description string `json:"description"`
	StarPowers  []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		Version     int    `json:"version"`
		Description string `json:"description"`
		ImageURL    string `json:"imageUrl"`
		Released    bool   `json:"released"`
	} `json:"starPowers"`
	Gadgets []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		Version     int    `json:"version"`
		Description string `json:"description"`
		ImageURL    string `json:"imageUrl"`
		Released    bool   `json:"released"`
	} `json:"gadgets"`
	Rarity struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"rarity"`
	Brawled struct {
		DiscordEmoji string `json:"discordEmoji"`
		BrawlerImage string `json:"brawlerImage"`
	} `json:"brawled"`
}

type GeneratedImage struct {
	URL           string `json:"url"`
	Tag           string `json:"tag"`
	Configuration string `json:"configuration"`
	Bytes         []byte `json:"bytes"`
	Timestamp     int64  `json:"timestamp"`
}
