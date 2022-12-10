package model

type UserDetails struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		UserDetails struct {
			UserID      string `json:"user_id"`
			UserStatus  string `json:"user_status"`
			UserAccount string `json:"user_account"`
			UserName    string `json:"user_name"`
			UserPremium string `json:"user_premium"`
			UserWebpage string `json:"user_webpage"`
			UserBirthRe string `json:"user_birth_re"`
			UserComment string `json:"user_comment"`
			ProfileImg  struct {
				Main  string `json:"main"`
				MainS string `json:"main_s"`
			} `json:"profile_img"`
			UserBirthTxt  string `json:"user_birth_txt"`
			IsFollowed    bool   `json:"is_followed"`
			IsFollowing   bool   `json:"is_following"`
			IsMypixiv     bool   `json:"is_mypixiv"`
			IsBlocked     bool   `json:"is_blocked"`
			IsBlocking    bool   `json:"is_blocking"`
			AcceptRequest bool   `json:"accept_request"`
			IsOfficial    bool   `json:"is_official"`
			Follows       int    `json:"follows"`
			Social        struct {
				Twitter struct {
					URL string `json:"url"`
				} `json:"twitter"`
			} `json:"social"`
			UserCommentHTML         string `json:"user_comment_html"`
			CoverImage              bool   `json:"cover_image"`
			HasIllusts              bool   `json:"has_illusts"`
			HasMangas               bool   `json:"has_mangas"`
			HasNovels               bool   `json:"has_novels"`
			HasBookmarks            bool   `json:"has_bookmarks"`
			HasIllustBookmarks      bool   `json:"has_illust_bookmarks"`
			HasNovelBookmarks       bool   `json:"has_novel_bookmarks"`
			ExternalSiteWorksStatus struct {
				Booth    bool `json:"booth"`
				Sketch   bool `json:"sketch"`
				VroidHub bool `json:"vroidHub"`
			} `json:"external_site_works_status"`
			HasFollows    bool `json:"has_follows"`
			HasMypixiv    bool `json:"has_mypixiv"`
			FanboxDetails struct {
				UserID               string `json:"user_id"`
				CreatorID            string `json:"creator_id"`
				Description          string `json:"description"`
				HasAdultContent      string `json:"has_adult_content"`
				RegistrationDatetime string `json:"registration_datetime"`
				UpdatedDatetime      string `json:"updated_datetime"`
				CoverImageURL        string `json:"cover_image_url"`
				URL                  string `json:"url"`
			} `json:"fanbox_details"`
			ShowRequestTab     bool `json:"show_request_tab"`
			ShowRequestSentTab bool `json:"show_request_sent_tab"`
			CanSendMessage     bool `json:"can_send_message"`
			Meta               struct {
				TwitterCard struct {
					Card        string `json:"card"`
					Site        string `json:"site"`
					Title       string `json:"title"`
					Image       string `json:"image"`
					Description string `json:"description"`
				} `json:"twitter_card"`
				Ogp struct {
					Title       string `json:"title"`
					Type        string `json:"type"`
					Image       string `json:"image"`
					Description string `json:"description"`
				} `json:"ogp"`
				Title              string `json:"title"`
				Description        string `json:"description"`
				DescriptionHeader  string `json:"description_header"`
				Canonical          string `json:"canonical"`
				AlternateLanguages struct {
					Ja string `json:"ja"`
					En string `json:"en"`
				} `json:"alternate_languages"`
			} `json:"meta"`
		} `json:"user_details"`
	} `json:"body"`
}
