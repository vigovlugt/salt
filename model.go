package main

type ThreadMainListItemNormalizer struct {
	Name  string `json:"name"`
	Props struct {
		Thread Thread `json:"thread"`
	} `json:"props"`
}

type Thread struct {
	ThreadID                       string  `json:"threadId"`
	ThreadTypeID                   int     `json:"threadTypeId"`
	TitleSlug                      string  `json:"titleSlug"`
	Title                          string  `json:"title"`
	CurrentUserVoteDirection       string  `json:"currentUserVoteDirection"`
	CommentCount                   int     `json:"commentCount"`
	Status                         string  `json:"status"`
	IsExpired                      bool    `json:"isExpired"`
	IsNew                          bool    `json:"isNew"`
	IsPinned                       bool    `json:"isPinned"`
	IsTrending                     *bool   `json:"isTrending"`
	IsBookmarked                   bool    `json:"isBookmarked"`
	IsLocal                        bool    `json:"isLocal"`
	Temperature                    float64 `json:"temperature"`
	TemperatureLevel               string  `json:"temperatureLevel"`
	Type                           string  `json:"type"`
	NSFW                           bool    `json:"nsfw"`
	DeletedAt                      *string `json:"deletedAt"`
	IsAffiliateTrackingDisabled    bool    `json:"isAffiliateTrackingDisabled"`
	IsAffiliateDescriptionDisabled bool    `json:"isAffiliateDescriptionDisabled"`
	IsEditLocked                   bool    `json:"isEditLocked"`
	IsExpireLocked                 bool    `json:"isExpireLocked"`
	ContentLockedBy                *string `json:"contentLockedBy"`
	IsSpamLocked                   bool    `json:"isSpamLocked"`
	IsLocked                       bool    `json:"isLocked"`
	IsNewsletterPicked             *bool   `json:"isNewsletterPicked"`
	IsCommentsModerationOn         *bool   `json:"isCommentsModerationOn"`
	IsPushed                       *bool   `json:"isPushed"`
	IsCommunityFavorite            bool    `json:"isCommunityFavorite"`
	IsCategoryCommunityFavorite    bool    `json:"isCategoryCommunityFavorite"`
	IsTopDeal                      bool    `json:"isTopDeal"`
	IsCategoryTopDeal              bool    `json:"isCategoryTopDeal"`
	PinID                          *string `json:"pinId"`
	BumpedAt                       int     `json:"bumpedAt"`
	BumpedAndReset                 bool    `json:"bumpedAndReset"`
	PublishedAt                    int     `json:"publishedAt"`
	VoucherCode                    string  `json:"voucherCode"`
	Link                           string  `json:"link"`
	ShareableLink                  string  `json:"shareableLink"`
	Merchant                       struct {
		MerchantID            int    `json:"merchantId"`
		MerchantName          string `json:"merchantName"`
		MerchantUrlName       string `json:"merchantUrlName"`
		IsMerchantPageEnabled bool   `json:"isMerchantPageEnabled"`
		Avatar                struct {
			Path       string `json:"path"`
			Name       string `json:"name"`
			SlotID     string `json:"slotId"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			Version    int    `json:"version"`
			Unattached bool   `json:"unattached"`
			UID        string `json:"uid"`
			Ext        string `json:"ext"`
		} `json:"avatar"`
	} `json:"merchant"`
	MainGroup struct {
		ThreadGroupID      int    `json:"threadGroupId"`
		ThreadGroupName    string `json:"threadGroupName"`
		ThreadGroupUrlName string `json:"threadGroupUrlName"`
	} `json:"mainGroup"`
	MainImage struct {
		Path       string `json:"path"`
		Name       string `json:"name"`
		SlotID     string `json:"slotId"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Version    int    `json:"version"`
		Unattached bool   `json:"unattached"`
		UID        string `json:"uid"`
		Ext        string `json:"ext"`
	} `json:"mainImage"`
	Price         float64 `json:"price"`
	NextBestPrice float64 `json:"nextBestPrice"`
	Percentage    int     `json:"percentage"`
	DiscountType  *string `json:"discountType"`
	Shipping      struct {
		IsFree *interface{} `json:"isFree"`
		Price  float64      `json:"price"`
	} `json:"shipping"`
	User struct {
		UserID   int    `json:"userId"`
		Username string `json:"username"`
		Title    string `json:"title"`
		Avatar   struct {
			Path       string `json:"path"`
			Name       string `json:"name"`
			SlotID     string `json:"slotId"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			Version    int    `json:"version"`
			Unattached bool   `json:"unattached"`
			UID        string `json:"uid"`
			Ext        string `json:"ext"`
		} `json:"avatar"`
		Persona struct {
			Text *string `json:"text"`
			Type *string `json:"type"`
		} `json:"persona"`
		IsBanned                   bool `json:"isBanned"`
		IsDeletedOrPendingDeletion bool `json:"isDeletedOrPendingDeletion"`
		IsUserProfileHidden        bool `json:"isUserProfileHidden"`
	} `json:"user"`
	StartDate struct {
		Timestamp int `json:"timestamp"`
	} `json:"startDate"`
	SelectedLocations struct {
		IsNational bool `json:"isNational"`
	} `json:"selectedLocations"`
	EndDate struct {
		Timestamp int `json:"timestamp"`
	} `json:"endDate"`
	IsExclusive         bool    `json:"isExclusive"`
	IsVoucherCodeHidden bool    `json:"isVoucherCodeHidden"`
	ClaimCodeCampaignID *string `json:"claimCodeCampaignId"`
	ClaimCodeCampaign   *string `json:"claimCodeCampaign"`
	CanVote             bool    `json:"canVote"`
}
