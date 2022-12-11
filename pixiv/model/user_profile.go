package model

import "time"

type UserProfile struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts     map[string]struct{} `json:"illusts"`
		Manga       map[string]struct{} `json:"manga"`
		Novels      map[string]struct{} `json:"novels"`
		MangaSeries []struct {
			ID             string      `json:"id"`
			UserID         string      `json:"userId"`
			Title          string      `json:"title"`
			Description    string      `json:"description"`
			Caption        string      `json:"caption"`
			Total          int         `json:"total"`
			ContentOrder   interface{} `json:"content_order"`
			URL            string      `json:"url"`
			CoverImageSl   int         `json:"coverImageSl"`
			FirstIllustID  string      `json:"firstIllustId"`
			LatestIllustID string      `json:"latestIllustId"`
			CreateDate     time.Time   `json:"createDate"`
			UpdateDate     time.Time   `json:"updateDate"`
			WatchCount     interface{} `json:"watchCount"`
			IsWatched      bool        `json:"isWatched"`
			IsNotifying    bool        `json:"isNotifying"`
		} `json:"mangaSeries"`
		NovelSeries []struct {
			ID                            string    `json:"id"`
			UserID                        string    `json:"userId"`
			UserName                      string    `json:"userName"`
			ProfileImageURL               string    `json:"profileImageUrl"`
			XRestrict                     int       `json:"xRestrict"`
			IsOriginal                    bool      `json:"isOriginal"`
			IsConcluded                   bool      `json:"isConcluded"`
			GenreID                       string    `json:"genreId"`
			Title                         string    `json:"title"`
			Caption                       string    `json:"caption"`
			Language                      string    `json:"language"`
			Tags                          []string  `json:"tags"`
			PublishedContentCount         int       `json:"publishedContentCount"`
			PublishedTotalCharacterCount  int       `json:"publishedTotalCharacterCount"`
			PublishedTotalWordCount       int       `json:"publishedTotalWordCount"`
			PublishedReadingTime          int       `json:"publishedReadingTime"`
			UseWordCount                  bool      `json:"useWordCount"`
			LastPublishedContentTimestamp int       `json:"lastPublishedContentTimestamp"`
			CreatedTimestamp              int       `json:"createdTimestamp"`
			UpdatedTimestamp              int       `json:"updatedTimestamp"`
			CreateDate                    time.Time `json:"createDate"`
			UpdateDate                    time.Time `json:"updateDate"`
			FirstNovelID                  string    `json:"firstNovelId"`
			LatestNovelID                 string    `json:"latestNovelId"`
			DisplaySeriesContentCount     int       `json:"displaySeriesContentCount"`
			ShareText                     string    `json:"shareText"`
			Total                         int       `json:"total"`
			FirstEpisode                  struct {
				URL string `json:"url"`
			} `json:"firstEpisode"`
			WatchCount   interface{} `json:"watchCount"`
			MaxXRestrict interface{} `json:"maxXRestrict"`
			Cover        struct {
				Urls struct {
					Two40Mw     string `json:"240mw"`
					Four80Mw    string `json:"480mw"`
					One200X1200 string `json:"1200x1200"`
					One28X128   string `json:"128x128"`
					Original    string `json:"original"`
				} `json:"urls"`
			} `json:"cover"`
			IsWatched   bool `json:"isWatched"`
			IsNotifying bool `json:"isNotifying"`
			AiType      int  `json:"aiType"`
		} `json:"novelSeries"`
		Pickup []struct {
			Type            string      `json:"type"`
			Deletable       bool        `json:"deletable"`
			Draggable       bool        `json:"draggable"`
			UserName        string      `json:"userName,omitempty"`
			UserImageURL    string      `json:"userImageUrl,omitempty"`
			ContentURL      string      `json:"contentUrl"`
			Description     string      `json:"description"`
			ImageURL        string      `json:"imageUrl,omitempty"`
			ImageURLMobile  string      `json:"imageUrlMobile,omitempty"`
			HasAdultContent bool        `json:"hasAdultContent,omitempty"`
			ID              string      `json:"id,omitempty"`
			UserID          string      `json:"userId,omitempty"`
			Title           string      `json:"title,omitempty"`
			Caption         string      `json:"caption,omitempty"`
			Total           int         `json:"total,omitempty"`
			ContentOrder    interface{} `json:"content_order,omitempty"`
			URL             string      `json:"url,omitempty"`
			CoverImageSl    int         `json:"coverImageSl,omitempty"`
			FirstIllustID   string      `json:"firstIllustId,omitempty"`
			LatestIllustID  string      `json:"latestIllustId,omitempty"`
			CreateDate      time.Time   `json:"createDate,omitempty"`
			UpdateDate      time.Time   `json:"updateDate,omitempty"`
			WatchCount      interface{} `json:"watchCount,omitempty"`
			IsWatched       bool        `json:"isWatched,omitempty"`
			IsNotifying     bool        `json:"isNotifying,omitempty"`
		} `json:"pickup"`
		BookmarkCount struct {
			Public struct {
				Illust int `json:"illust"`
				Novel  int `json:"novel"`
			} `json:"public"`
			Private struct {
				Illust int `json:"illust"`
				Novel  int `json:"novel"`
			} `json:"private"`
		} `json:"bookmarkCount"`
		ExternalSiteWorksStatus struct {
			Booth    bool `json:"booth"`
			Sketch   bool `json:"sketch"`
			VroidHub bool `json:"vroidHub"`
		} `json:"externalSiteWorksStatus"`
		Request struct {
			ShowRequestTab     bool `json:"showRequestTab"`
			ShowRequestSentTab bool `json:"showRequestSentTab"`
			PostWorks          struct {
				Artworks []interface{} `json:"artworks"`
				Novels   []interface{} `json:"novels"`
			} `json:"postWorks"`
		} `json:"request"`
	} `json:"body"`
}
