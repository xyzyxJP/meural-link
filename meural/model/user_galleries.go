package model

type UserGalleries struct {
	Data []struct {
		CategoryIds     []interface{} `json:"categoryIds"`
		CategoryNames   []interface{} `json:"categoryNames"`
		Channel         interface{}   `json:"channel"`
		Cover           string        `json:"cover"`
		Cover2X         string        `json:"cover2x"`
		CoverID         int           `json:"coverId"`
		Currency        interface{}   `json:"currency"`
		Description     interface{}   `json:"description"`
		FavoriteCount   int           `json:"favoriteCount"`
		FeaturedAt      interface{}   `json:"featuredAt"`
		Hero            string        `json:"hero"`
		Hero2X          string        `json:"hero2x"`
		Icon            interface{}   `json:"icon"`
		Icon2X          interface{}   `json:"icon2x"`
		ID              int           `json:"id"`
		IsPublic        bool          `json:"isPublic"`
		IsPurchased     bool          `json:"isPurchased"`
		IsSampler       bool          `json:"isSampler"`
		IsPremium       bool          `json:"isPremium"`
		ItemCount       int           `json:"itemCount"`
		ItemIds         []int         `json:"itemIds"`
		ItemExternalIds []interface{} `json:"itemExternalIds"`
		Name            string        `json:"name"`
		Orientation     string        `json:"orientation"`
		PreviewItems    []struct {
			ArtistID       interface{}   `json:"artistId"`
			ArtistName     interface{}   `json:"artistName"`
			Author         string        `json:"author"`
			CategoryIds    []interface{} `json:"categoryIds"`
			CategoryNames  []interface{} `json:"categoryNames"`
			Channel        interface{}   `json:"channel"`
			Copyright      interface{}   `json:"copyright"`
			CreatedAt      string        `json:"createdAt"`
			CroppedHeight  int           `json:"croppedHeight"`
			CroppedWidth   int           `json:"croppedWidth"`
			CroppedX1      int           `json:"croppedX1"`
			CroppedY1      int           `json:"croppedY1"`
			Currency       interface{}   `json:"currency"`
			Description    string        `json:"description"`
			Dimensions     interface{}   `json:"dimensions"`
			FavoriteCount  int           `json:"favoriteCount"`
			Hero           string        `json:"hero"`
			Hero2X         string        `json:"hero2x"`
			Icon           interface{}   `json:"icon"`
			Icon2X         interface{}   `json:"icon2x"`
			ID             int           `json:"id"`
			Image          string        `json:"image"`
			IsDetailed     bool          `json:"isDetailed"`
			IsPublic       bool          `json:"isPublic"`
			IsSampler      bool          `json:"isSampler"`
			IsPremium      bool          `json:"isPremium"`
			IsOnUserDevice bool          `json:"isOnUserDevice"`
			Medium         string        `json:"medium"`
			Name           string        `json:"name"`
			OriginalImage  string        `json:"originalImage"`
			OriginalHeight int           `json:"originalHeight"`
			OriginalWidth  int           `json:"originalWidth"`
			PremiumDate    interface{}   `json:"premiumDate"`
			SortedOrder    interface{}   `json:"sortedOrder"`
			Src            string        `json:"src"`
			SrcAttrs       struct {
			} `json:"srcAttrs"`
			NonsubscriberPrice int    `json:"nonsubscriberPrice"`
			SubscriberPrice    int    `json:"subscriberPrice"`
			Slug               string `json:"slug"`
			Type               string `json:"type"`
			UpdatedAt          string `json:"updatedAt"`
			UserID             int    `json:"userId"`
			Year               string `json:"year"`
		} `json:"previewItems"`
		Slug               interface{} `json:"slug"`
		NonsubscriberPrice int         `json:"nonsubscriberPrice"`
		SubscriberPrice    int         `json:"subscriberPrice"`
		PendingInviteCount int         `json:"pendingInviteCount"`
		PremiumDate        interface{} `json:"premiumDate"`
		ShareCount         int         `json:"shareCount"`
		Type               string      `json:"type"`
		UpdatedAt          string      `json:"updatedAt"`
		Username           string      `json:"username"`
		UserID             int         `json:"userId"`
		UserFirstName      string      `json:"userFirstName"`
		UserLastName       string      `json:"userLastName"`
	} `json:"data"`
	IsPaginated bool   `json:"isPaginated"`
	IsLast      bool   `json:"isLast"`
	Count       int    `json:"count"`
	Message     string `json:"message"`
}