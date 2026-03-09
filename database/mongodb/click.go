package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schema struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	ApplicationID string             `bson:"aid" json:"aid"`

	AdvertiserID string `bson:"oid" json:"oid"`
	PartnerID    string `bson:"pid,omitempty" json:"pid,omitempty"`
	AgencyID     string `bson:"agid,omitempty" json:"agid,omitempty"`

	WebSessionID   string `bson:"wsid,omitempty" json:"wsid,omitempty"` // Web Session ID for click (if click is from web)
	CustomerUserID string `bson:"cuid,omitempty" json:"cuid,omitempty"` // Advertiser's Customer User ID (if provided by advertiser)

	Type string `bson:"type,omitempty" json:"type,omitempty"` // can contains 's2s', 'unilink', 'default'
	URL  string `bson:"url,omitempty" json:"url,omitempty"`

	LookbackWindow        string `bson:"lbws" json:"lbws"`
	LookbackWindowSeconds int64  `bson:"lbw" json:"lbw"`

	ClickToEventWindow        string `bson:"ev_ctlws,omitempty" json:"ev_ctlws"`
	ClickToEventWindowSeconds int64  `bson:"ev_ctlw,omitempty" json:"ev_ctlw"`

	BlackOutEnabled bool `bson:"blk_enabled,omitempty" json:"blk_enabled"`

	InstallToEventWindow        string `bson:"ite_lwds,omitempty" json:"ite_lwds"`
	InstallToEventWindowSeconds int64  `bson:"ite_lwd,omitempty" json:"ite_lwd"`

	// Location related properties
	IP             string `bson:"ip,omitempty" json:"ip,omitempty"`
	ISP            string `bson:"isp,omitempty" json:"isp,omitempty"`
	ConnectionType string `bson:"ct,omitempty" json:"ct,omitempty"`

	HeaderXForwardedFor string `bson:"h_xff,omitempty" json:"h_xff,omitempty"`
	HeaderUserAgent     string `bson:"h_ua,omitempty" json:"h_ua,omitempty"`

	Country   string  `bson:"geo,omitempty" json:"geo,omitempty"`
	City      string  `bson:"city,omitempty" json:"city,omitempty"`
	Region    string  `bson:"rgn,omitempty" json:"rgn,omitempty"`
	Latitude  float64 `bson:"lat,omitempty" json:"lat,omitempty"`
	Longitude float64 `bson:"lng,omitempty" json:"lng,omitempty"`

	ProxyCountry string `bson:"pxc,omitempty" json:"pxc,omitempty"`

	UserAgent   string `bson:"ua,omitempty" json:"ua,omitempty"`
	ReferrerURL string `bson:"cref,omitempty" json:"cref,omitempty"` // Click http referrer header

	OS          string `bson:"os,omitempty" json:"os,omitempty"`
	OSVersion   string `bson:"osv,omitempty" json:"osv,omitempty"`
	DeviceModel string `bson:"dm,omitempty" json:"dm,omitempty"`
	Device      string `bson:"dvc,omitempty" json:"dvc,omitempty"`
	Browser     string `bson:"br,omitempty" json:"br,omitempty"`
	IsUnique    int    `bson:"iu,omitempty" json:"iu,omitempty"`
	DeviceLang  string `bson:"dl,omitempty" json:"dl,omitempty"`

	// Advertising identifiers
	GAID         string `bson:"gaid,omitempty" json:"gaid,omitempty"`
	IDFA         string `bson:"idfa,omitempty" json:"idfa,omitempty"`
	IDFV         string `bson:"idfv,omitempty" json:"idfv,omitempty"`
	AndroidID    string `bson:"anid,omitempty" json:"anid,omitempty"`
	AmazonFireID string `bson:"fid,omitempty" json:"fid,omitempty"`

	// Application detail
	BundleID string `bson:"bid,omitempty" json:"bid,omitempty"`
	AppName  string `bson:"an,omitempty" json:"an,omitempty"`

	// Device Hardware
	IMEI       string `bson:"imei,omitempty" json:"imei,omitempty"`
	IMEI2      string `bson:"imei2,omitempty" json:"imei2,omitempty"`
	MacAddress string `bson:"mac,omitempty" json:"mac,omitempty"`

	// Optional (Click) custom parameter properties
	Param1 string `bson:"p1,omitempty" json:"p1,omitempty"`
	Param2 string `bson:"p2,omitempty" json:"p2,omitempty"`
	Param3 string `bson:"p3,omitempty" json:"p3,omitempty"`
	Param4 string `bson:"p4,omitempty" json:"p4,omitempty"`
	Param5 string `bson:"p5,omitempty" json:"p5,omitempty"`

	// Ad, AdGroup, Site (Publisher) information
	RefClickID string `bson:"cid,omitempty" json:"cid,omitempty"` // Ad network unique Click ID
	AdGroup    string `bson:"adg,omitempty" json:"adg,omitempty"`
	AdGroupID  string `bson:"adgid,omitempty" json:"adgid,omitempty"`
	Ad         string `bson:"ad,omitempty" json:"ad,omitempty"`
	AdID       string `bson:"adid,omitempty" json:"adid,omitempty"`
	Campaign   string `bson:"camp,omitempty" json:"camp,omitempty"`     // Campaign (Provided by advertiser or publisher)
	CampaignID string `bson:"campid,omitempty" json:"campid,omitempty"` // Campaign ID (Provided by advertiser or publisher)
	AdType     string `bson:"adt,omitempty" json:"adt,omitempty"`       // e.g. text, banner, interstitial, video, sponsored_content, etc.
	SiteID     string `bson:"sid,omitempty" json:"sid,omitempty"`       // Ad Network publisher ID
	SubSiteID  string `bson:"ssid,omitempty" json:"ssid,omitempty"`     // Ad Network Sub Publisher ID
	Channel    string `bson:"ch,omitempty" json:"ch,omitempty"`         // Media source channel through which the ads are distributed

	IsRetargeting int    `bson:"rt,omitempty" json:"rt,omitempty"`
	Keywords      string `bson:"kwd,omitempty" json:"kwd,omitempty"`

	// cost
	CostCurrency string  `bson:"curr" json:"curr"` // 3 letter currency code
	CostValue    float64 `bson:"cost" json:"cost"` // Cost value in using cost currency
	CostValueUSD float64 `bson:"cusd" json:"cusd"` // Cost value in USD
	CurrencyRate float64 `bson:"crt" json:"crt"`

	// redirect's
	GenericRedirectURL string `bson:"r,omitempty" json:"r,omitempty"`
	WebRedirectURL     string `bson:"webr,omitempty" json:"webr,omitempty"`
	IOSRedirectURL     string `bson:"iosr,omitempty" json:"iosr,omitempty"`
	AndroidRedirectURL string `bson:"andr,omitempty" json:"andr,omitempty"`

	// Unilink related field
	CustomLinkID   string `bson:"clid,omitempty" json:"clid,omitempty"`
	LinkTemplateID string `bson:"ltid,omitempty" json:"ltid,omitempty"`

	// fraudelent
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	Reason string `bson:"reason,omitempty" json:"reason,omitempty"`

	DynamicLinkID primitive.ObjectID `bson:"dlid,omitempty" json:"dlid,omitempty"`

	IsInstallPostbackDisabled uint8 `bson:"dipb" json:"dipb"`

	CreatedTS time.Time `bson:"crtd" json:"crtd"`
}

func (s *Schema) IsEmpty() bool {
	return s.ID.IsZero()
}
