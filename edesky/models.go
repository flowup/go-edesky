package edesky

type Meta struct {

}

type APIOptions struct {
	APIKey string `json:"api_key,omitempty"`
}

type DocumentOptions struct {
	APIOptions
	Keywords string
}

type DocumentRoot struct {
	Meta Meta `xml:"meta"`
	Documents []Document `xml:"documents>document"`
}

type Document struct {
	DashboardCategory string `xml:"dashboard_category,attr"`
	DashboardName string `xml:"dashboard_name,attr"`
	EDeskyURL string `xml:"edesky_url,attr"`
	Name string `xml:"name,attr"`
}

type DashboardOptions struct {
	APIOptions
	ID int `json:"id,omitempty"`
	IncludeSubordinated int `json:"include_subordinated,omitempty"`
}

type DashboardRoot struct {
	Meta Meta `xml:"meta"`
	Dashboards []Dashboard `xml:"dashboards>dashboard"`
}

type Dashboard struct {
	Category string `xml:"category,attr"`
	EDeskyID int `xml:"edesky_id,attr"`
	EDeskyURL string `xml:"edesky_url,attr"`
	Name string `xml:"name,attr"`
	Region string `xml:"nuts3_name,attr"`
	District string `xml:"nuts4_name,attr"`
}