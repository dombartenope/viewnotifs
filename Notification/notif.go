package Notification

type Notification struct {
	DelayedOption          *string  `json:"delayed_option"`
	DeliveryTimeOfDay      *string  `json:"delivery_time_of_day"`
	Errored                int      `json:"errored"`
	Failed                 int      `json:"failed"`
	ID                     string   `json:"id"`
	IncludePlayerIds       []string `json:"include_player_ids"`
	IncludeExternalUserIds []string `json:"include_external_user_ids"`
	IncludeAliases         []string `json:"include_aliases"`
	IncludedSegments       []string `json:"included_segments"`
	SendAfter              int64    `json:"send_after"`
	CompletedAt            int64    `json:"completed_at"`
	Successful             int      `json:"successful"`
	Received               *int     `json:"received"`
	TemplateId             *string  `json:"template_id"`
}

type ApiResponse struct {
	TotalCount    int            `json:"total_count"`
	Offset        int            `json:"offset"`
	Limit         interface{}    `json:"limit"`
	Notifications []Notification `json:"notifications"`
}
